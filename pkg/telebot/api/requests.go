package api

import (
	"finance-bot/pkg/telebot/models"
)

// /getUpdates
type UpdateResponse struct {
	Ok     bool            `json:"ok"`
	Result []models.Update `json:"result"`
}

// /getMe
type GetMeResponse struct {
	Ok   bool         `json:"ok"`
	User *models.User `json:"result"`
}
