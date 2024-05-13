package api

import (
	"encoding/json"
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/models"
	"fmt"
	"net/http"
)

type TelegramBotAPI struct {
	Token       string
	telegramURL string
	logger      log.Logger
}

func NewTelegramBotAPI(token, telegramURL string, logger log.Logger) *TelegramBotAPI {
	return &TelegramBotAPI{
		Token:       token,
		logger:      logger,
		telegramURL: telegramURL,
	}
}

func (t *TelegramBotAPI) GetMe() (*models.User, error) {
	url := t.telegramURL + t.Token + "/getMe"
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

func (t *TelegramBotAPI) SendMessageText(chatID int, text string) (models.Message, error) {
	messageUrl := t.telegramURL + t.Token + "/sendMessage?chat_id=" + fmt.Sprint(chatID) + "&text=" + text

	t.logger.Debug("http: request to telegram api with url: " + messageUrl)

	resp, err := http.DefaultClient.Get(messageUrl)
	if err != nil {
		t.logger.Error("http: failed to make http request with url: " + messageUrl)
		return models.Message{}, err
	}
	defer resp.Body.Close()

	body := SendMessageResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return models.Message{}, err
	}
	if !body.Ok {
		return models.Message{}, fmt.Errorf("http: failed to send message")
	}

	t.logger.Debug("message sent:\n" + formatDebugMessageModel(body.Result))
	return body.Result, nil
}

func (t *TelegramBotAPI) GetUpdates(offset int) ([]models.Update, error) {
	url := t.telegramURL + t.Token + "/getUpdates?timeout=60&offset=" + fmt.Sprint(offset)

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

func formatDebugMessageModel(msg models.Message) string {
	var name string
	if msg.From.Username == "" {
		name = msg.Chat.Title
	} else {
		name = msg.Chat.Username
	}

	return fmt.Sprintf("\tID: %d\n\tDate: %d\n\tChat: %s\n\tText: %s", msg.MessageID, msg.Date, name, msg.Text)
}
