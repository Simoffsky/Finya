package telebot

import (
	"context"
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/models"
	"fmt"
	"strings"
	"sync"
)

type CommandHandler func(models.Message) error

type TelebotAPI interface {
	GetMe() (*models.User, error)
	SendMessageText(chatID int, text string) (models.Message, error)
	GetUpdates(offset int) ([]models.Update, error)
}

type FailedUpdate struct {
	Error  error
	Update models.Update
}

type Telebot struct {
	apiClient       TelebotAPI
	logger          log.Logger
	mx              sync.RWMutex
	offset          int
	commandHandlers map[string]CommandHandler
	errch           chan FailedUpdate // channel for errors from updates
}

func NewTelebot(token string, logger log.Logger, api TelebotAPI) *Telebot {
	return &Telebot{
		apiClient: api,
		logger:    logger,
		errch:     make(chan FailedUpdate),
	}
}

func (t *Telebot) Errors() <-chan FailedUpdate {
	return t.errch
}

// Closing the channels, stops the bot
// TODO: Tests?
func (t *Telebot) Close() error {
	close(t.errch)
	return nil
}

// RegisterCommand registers a command with a handler
// will overwrite the handler if the command already exists
func (t *Telebot) RegisterCommand(command string, handler CommandHandler) {
	if t.commandHandlers == nil {
		t.commandHandlers = make(map[string]CommandHandler)
	}
	t.commandHandlers[command] = handler
}

func (t *Telebot) GetCommandHandlers() map[string]CommandHandler {
	return t.commandHandlers
}

func (t *Telebot) GetOffset() int {
	return t.offset
}

// handleUpdates handles updates from the Telegram API
// and calls the appropriate handler for each update concurrently
// errors are sent to the error channel - t.Errors()
func (t *Telebot) handleUpdates(updates []models.Update) {
	if len(updates) == 0 {
		return
	}
	t.offset = updates[len(updates)-1].UpdateID + 1
	for _, update := range updates {
		update := update
		go func() {
			err := t.handleUpdate(update)
			if err != nil {
				t.errch <- FailedUpdate{err, update}
			}
		}()
	}
}

func (t *Telebot) handleUpdate(update models.Update) error {
	t.logger.Debug(fmt.Sprintf("handling update:\n\tID: %d", update.UpdateID))
	if update.Message == nil {
		return models.ErrUpdateEmptyMessage
	}
	msg := update.Message

	if err := t.handleMessage(*msg); err != nil {
		return err
	}
	return nil
}

func (t *Telebot) handleMessage(message models.Message) error {
	t.logger.Debug(
		fmt.Sprintf("handling message:\n\tID: %d\n\tDate: %d\n\tChat: %s\n\tText: %s", message.MessageID, message.Date, message.Chat, message.Text),
	)

	for _, entity := range message.Entities {
		if entity.Type != models.EntityTypeBotCommand {
			continue
		}
		if err := t.handleCommand(message); err != nil {
			return err
		}
	}
	return nil
}

// can handle message that starts with '/'
func (t *Telebot) handleCommand(message models.Message) error {
	splitted := strings.Split(message.Text, " ")
	if len(splitted) == 0 {
		return models.ErrCommandNotFound
	}
	command := strings.TrimSpace(splitted[0])

	t.mx.RLock()
	handler, ok := t.commandHandlers[command]
	t.mx.RUnlock()

	if !ok {
		return models.ErrCommandNotFound
	}
	return handler(message)
}

// returns *models.User that provides information about the bot
func (t *Telebot) GetMe() (*models.User, error) {
	t.logger.Debug("getting information about the bot")
	return t.apiClient.GetMe()
}

func (t *Telebot) SendMessage(chatID int, text string) (models.Message, error) {
	t.logger.Debug("sending message to chat with ID: " + fmt.Sprint(chatID))
	return t.apiClient.SendMessageText(chatID, text)
}

func (t *Telebot) LongPooling(ctx context.Context) error {
	t.logger.Info("long pooling started")
	for {
		select {
		case <-ctx.Done():
			t.logger.Info("long pooling stopped")
			return nil
		default:
			updates, err := t.apiClient.GetUpdates(t.offset)
			if err != nil {
				return err
			}

			t.handleUpdates(updates)
		}
	}
}
