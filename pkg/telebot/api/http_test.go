package api

import (
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T) {
	testCases := []struct {
		name           string
		serverResponse string
		expectError    bool
		expectedUser   *models.User
	}{
		{
			name:           "Success",
			serverResponse: `{"ok": true, "result": {"id": 123, "is_bot": true, "first_name": "TestBot"}}`,
			expectError:    false,
			expectedUser:   &models.User{ID: 123, IsBot: true, FirstName: "TestBot"},
		},
		{
			name:           "FailureDueToParsing",
			serverResponse: `{"ok": true, "result": {"id": 123, "is_bot": true, "first_name": "TestBot",}`,
			expectError:    true,
		},
		{
			name:           "FailureDueToResponse",
			serverResponse: `{"ok": false, "result": {"id": 123, "is_bot": true, "first_name": "TestBot"}}`,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.Write([]byte(tc.serverResponse))
			}))

			defer server.Close()

			bot := &TelegramBotAPI{
				telegramURL: server.URL + "/",
				Token:       "testToken",
				logger:      &log.MockLogger{},
			}

			user, err := bot.GetMe()

			if tc.expectError {
				assert.Error(t, err, "expected an error")
			} else {
				assert.NoErrorf(t, err, "expected no error, got %v", err)
				assert.Equal(t, tc.expectedUser.ID, user.ID)
				assert.Equal(t, tc.expectedUser.IsBot, user.IsBot)
				assert.Equal(t, tc.expectedUser.FirstName, user.FirstName)
			}
		})
	}
}

func TestSendMessageText(t *testing.T) {

	testCases := []struct {
		name           string
		serverResponse string
		expectError    bool
		expectedResult *models.Message
	}{
		{
			name: "Success",
			serverResponse: `{
                "ok": true,
                "result": {
                    "message_id": 1,
                    "from": {"id": 123, "is_bot": true, "first_name": "TestBot"},
                    "chat": {"id": 456},
                    "date": 1616161616,
                    "text": "Test_message"
                }
            }`,
			expectError: false,
			expectedResult: &models.Message{
				MessageID: 1,
				From: &models.User{
					ID:        123,
					IsBot:     true,
					FirstName: "TestBot",
				},
				Chat: &models.Chat{
					ID: 456,
				},
				Date: 1616161616,
				Text: "Test_message",
			},
		},
		{
			name:           "FailureDueToParsing",
			serverResponse: `{"ok": true, "result": {"message_id": 1, "from": {"id": 123, "is_bot": true, "first_name": "TestBot",}`,
			expectError:    true,
		},
		{
			name:           "FailureDueToResponse",
			serverResponse: `{"ok": false, "result": {"message_id": 1, "from": {"id": 123, "is_bot": true, "first_name": "TestBot"}}}`,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				rw.Write([]byte(tc.serverResponse))
			}))

			defer server.Close()

			bot := &TelegramBotAPI{
				telegramURL: server.URL + "/",
				Token:       "testToken",
				logger:      &log.MockLogger{},
			}

			message, err := bot.SendMessageText(456, "Test_message")

			if tc.expectError {
				assert.Error(t, err, "expected an error")
			} else {
				assert.NoErrorf(t, err, "expected no error, got %v", err)
				assert.Equal(t, *tc.expectedResult, message)
			}
		})
	}
}
