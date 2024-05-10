package api

import (
	"finance-bot/pkg/telebot/models"
)

// /getUpdates
type GetUpdatesResponse struct {
	Ok     bool            `json:"ok"`
	Result []models.Update `json:"result"`
}

// /sendMessage
type SendMessageResponse struct {
	Ok     bool           `json:"ok"`
	Result models.Message `json:"result"`
}

// /getMe
type GetMeResponse struct {
	Ok   bool         `json:"ok"`
	User *models.User `json:"result"`
}
