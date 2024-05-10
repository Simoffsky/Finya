package telebot

import (
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/api"
	"finance-bot/pkg/telebot/models"
	"fmt"
	"strings"
	"sync"
)

type CommandHandler func(models.Message)

type Telebot struct {
	apiClient *api.TelegramBotAPI
	logger    log.Logger

	mx              sync.Mutex
	offset          int
	commandHandlers map[string]CommandHandler
}

func (t *Telebot) RegisterCommand(command string, handler CommandHandler) {
	if t.commandHandlers == nil {
		t.commandHandlers = make(map[string]CommandHandler)
	}
	t.commandHandlers[command] = handler
}

func (t *Telebot) HandleUpdates(updates []models.Update) {

	fmt.Println(len(updates))
	if len(updates) == 0 {
		return
	}

	t.offset = updates[len(updates)-1].UpdateID + 1
	for _, update := range updates {
		go t.handleUpdate(update)
	}
}

func (t *Telebot) handleUpdate(update models.Update) {
	t.logger.Debug(fmt.Sprintf("handling update:\n\tID: %d\n\tMessage ID: %d\n\tText: %s", update.UpdateID, update.Message.MessageID, update.Message.Text))

	if update.Message == nil {
		return
	}
	msg := update.Message

	for _, entity := range msg.Entities {
		if entity.Type == models.MessageEntityTypeBotCommand {
			t.handleCommand(*msg)
		}
	}

}

// can handle message that starts with '/'
func (t *Telebot) handleCommand(message models.Message) {
	command := strings.Split(message.Text, " ")[0]
	t.mx.Lock()
	handler, ok := t.commandHandlers[command]
	t.mx.Unlock()
	if !ok {
		t.logger.Debug("command not found: " + command)
		return
	}
	handler(message)
}

func NewTelebot(token string, logger log.Logger) *Telebot {
	return &Telebot{
		apiClient: api.NewTelegramBotAPI(token, logger),
		logger:    logger,
	}
}

func (t *Telebot) GetMe() (*models.User, error) {
	t.logger.Debug("getting information about the bot")
	return t.apiClient.GetMe()
}

func (t *Telebot) SendMessage(chatID int, text string) error {
	t.logger.Debug("sending message to chat with ID: " + fmt.Sprint(chatID))
	return t.apiClient.SendMessageText(chatID, text)
}

func (t *Telebot) LongPooling() error {
	t.logger.Info("long pooling started")
	for {
		updates, err := t.apiClient.GetUpdates(t.offset)
		if err != nil {
			return err
		}

		t.HandleUpdates(updates)
	}
}
