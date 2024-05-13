package api

import (
	"finance-bot/pkg/telebot/models"

	"github.com/stretchr/testify/mock"
)

type TelegramMockAPI struct {
	mock.Mock
}

func (t *TelegramMockAPI) GetMe() (*models.User, error) {
	args := t.Called()
	return args.Get(0).(*models.User), args.Error(1)
}

func (t *TelegramMockAPI) SendMessageText(chatID int, text string) (models.Message, error) {
	args := t.Called(chatID, text)
	return args.Get(0).(models.Message), args.Error(1)
}

func (t *TelegramMockAPI) GetUpdates(offset int) ([]models.Update, error) {
	args := t.Called(offset)
	return args.Get(0).([]models.Update), args.Error(1)
}
