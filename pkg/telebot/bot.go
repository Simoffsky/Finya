package telebot

import (
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/api"
	"finance-bot/pkg/telebot/models"
)

type Telebot struct {
	apiClient *api.TelegramBotAPI
	logger    log.Logger
}

func NewTelebot(token string, logger log.Logger) *Telebot {
	return &Telebot{
		apiClient: api.NewTelegramBotAPI(token, logger),
		logger:    logger,
	}
}

func (t *Telebot) GetMe() (*models.User, error) {
	return t.apiClient.GetMe()
}
