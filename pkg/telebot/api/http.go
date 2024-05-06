package api

import (
	"encoding/json"
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/models"
	"fmt"
	"net/http"
)

const (
	telegramUrl = "https://api.telegram.org/bot"
)

type TelegramBotAPI struct {
	Token  string
	logger log.Logger
}

func NewTelegramBotAPI(token string, logger log.Logger) *TelegramBotAPI {
	return &TelegramBotAPI{
		Token:  token,
		logger: logger,
	}
}

func (t *TelegramBotAPI) GetMe() (*models.User, error) {
	url := telegramUrl + t.Token + "/getMe"
	t.logger.Debug("http: request to telegram api with url: " + url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		t.logger.Error("http: failed to make http request with url: " + url)
		return nil, err
	}
	defer resp.Body.Close()

	body := GetMeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	if !body.Ok {
		return nil, fmt.Errorf("http: failed to get information about the bot")
	}
	return body.User, nil
}

func (t *TelegramBotAPI) GetUpdates(offset int) ([]models.Update, error) {
	url := telegramUrl + t.Token + "/getUpdates?timeout=60&offset=" + fmt.Sprint(offset)

	t.logger.Debug("http: request to telegram api with url: " + url)

	resp, err := http.DefaultClient.Get(url)
	if err != nil {

		t.logger.Error("http: failed to make http request with url: " + url)
		return nil, err
	}
	defer resp.Body.Close()

	body := GetUpdatesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}
	if !body.Ok {
		return nil, fmt.Errorf("http: failed to get updates")
	}

	return body.Result, nil
}
