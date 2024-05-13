package bot

import (
	"context"
	"finance-bot/internal/config"
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot"
	"finance-bot/pkg/telebot/api"
	"finance-bot/pkg/telebot/models"
	"os"
	"os/signal"
)

type Bot struct {
	telebot *telebot.Telebot
	logger  log.Logger
	config  *config.Config
}

func NewBot(cfg *config.Config) *Bot {
	return &Bot{
		config: cfg,
	}
}

func (b *Bot) configure() error {
	b.logger = log.NewDefaultLogger(
		log.LevelFromString(b.config.LogLevel),
	).WithTimePrefix()

	api := api.NewTelegramBotAPI(b.config.BotToken, "https://api.telegram.org/bot", b.logger)
	b.telebot = telebot.NewTelebot(b.config.BotToken, b.logger, api)

	return nil
}
func (bot *Bot) Start() error {
	if err := bot.configure(); err != nil {
		return err
	}

	bot.telebot.RegisterCommand("/start", func(msg models.Message) error {
		_, err := bot.telebot.SendMessage(msg.Chat.ID, "Hello! I'm a finance bot. I can help you with finance stuff.")
		if err != nil {
			bot.logger.Error("failed to send message: " + err.Error())
			return err
		}
		return nil
	})

	bot.telebot.RegisterCommand("/echo", func(msg models.Message) error {
		_, err := bot.telebot.SendMessage(msg.Chat.ID, msg.Text[5:])
		if err != nil {
			bot.logger.Error("failed to send message: " + err.Error())
			return err
		}
		return nil
	})

	go func() {
		for failedUpdate := range bot.telebot.Errors() {
			if failedUpdate.Error == models.ErrCommandNotFound {
				_, _ = bot.telebot.SendMessage(failedUpdate.Update.Message.Chat.ID, "Command not found") // TODO: handle error
			}
			bot.logger.Error("received error from telebot: " + failedUpdate.Error.Error())
		}
	}()

	errCh := make(chan error, 1)
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	go func() {
		err := bot.telebot.LongPooling(context.Background())
		if err != nil {
			bot.logger.Error("failed to long pool: " + err.Error())
		}
		errCh <- err

	}()

	select {
	case <-sigint:
		bot.logger.Info("received SIGINT, stopping...")
	case err := <-errCh:
		if err != nil {
			bot.logger.Error(err.Error())
		}
		return err
	}

	bot.logger.Info("stopped gracefully")
	return nil
}
