package telebot

import (
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/api"
	"finance-bot/pkg/telebot/models"
	"fmt"
)

type Telebot struct {
	apiClient *api.TelegramBotAPI
	logger    log.Logger
	offset    int
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

func (t *Telebot) LongPooling() error {
	for {

		updates, err := t.apiClient.GetUpdates(t.offset)
		if err != nil {
			return err
		}
		for _, update := range updates {
			t.offset = update.UpdateID + 1
			t.logger.Debug("new update: \n" + "\tID: " + fmt.Sprint(update.UpdateID) + "\n" + "\tUser: " + update.Message.From.Username + "\n" + "\tMessage: " + update.Message.Text)
		}

	}
}
