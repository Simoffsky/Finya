package telebot

import (
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/api"
	"finance-bot/pkg/telebot/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterCommand(t *testing.T) {
	bot := NewTelebot("test", nil, nil)

	handler := func(message models.Message) error { return nil }

	bot.RegisterCommand("test", handler)

	commandHandlers := bot.GetCommandHandlers()
	assert.Contains(t, commandHandlers, "test")
}

func TestHandleUpdates(t *testing.T) {
	tbot := NewTelebot("test", &log.MockLogger{}, &api.TelegramMockAPI{})

	updates := []models.Update{}

	tbot.HandleUpdates(updates)

	if tbot.offset != 0 {
		t.Errorf("Expected offset to be 0, but got %v", tbot.offset)
	}

	updates = append(updates,
		models.Update{
			UpdateID: 1,
			Message: &models.Message{
				MessageID: 1,
				Text:      "test",
			},
		},
		models.Update{
			UpdateID: 2,
			Message: &models.Message{
				MessageID: 2,
				Text:      "test",
			},
		},
	)

	tbot.HandleUpdates(updates)

	if tbot.offset != 3 {
		t.Errorf("Expected offset to be 3, but got %v", tbot.offset)
	}

	updates = []models.Update{}

	tbot.HandleUpdates(updates)

	if tbot.offset != 3 {
		t.Errorf("Expected offset to remain 3, but got %v", tbot.offset)
	}

}

func TestHandleUpdate(t *testing.T) {
	tbot := NewTelebot("test", &log.MockLogger{}, &api.TelegramMockAPI{})

	tbot.RegisterCommand("/test", func(message models.Message) error {
		if message.Text != "/test command" {
			t.Errorf("Expected '/test command', got '%s'", message.Text)
		}
		return nil
	})
	testCases := []struct {
		name        string
		update      models.Update
		expectError bool
	}{
		{
			name: "SuccessCommand",
			update: models.Update{
				UpdateID: 1,
				Message: &models.Message{
					MessageID: 1,
					Text:      "/test command",
					Entities: []models.MessageEntity{
						{
							Type: models.EntityTypeBotCommand,
						},
					},
				},
			},
			expectError: false,
		},
		{
			name: "FailureDueToEmptyMessage",
			update: models.Update{
				UpdateID: 1,
				Message:  nil,
			},
			expectError: true,
		},
		{
			name: "SuccesNonCommand",
			update: models.Update{
				UpdateID: 1,
				Message: &models.Message{
					MessageID: 1,
					Text:      "non-command text",
					Entities: []models.MessageEntity{
						{
							Type: models.EntityTypeMention,
						},
					},
				},
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tbot.handleUpdate(tc.update)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestHandleCommand(t *testing.T) {
	handler := func(message models.Message) error {
		if message.Text != "/test command" {
			t.Errorf("Expected '/test command', got '%s'", message.Text)
		}
		return nil
	}

	bot := NewTelebot("test", &log.MockLogger{}, &api.TelegramMockAPI{})
	bot.RegisterCommand("/test", handler)

	msg := models.Message{
		Text: "/test command",
	}

	err := bot.handleCommand(msg)

	require.NoError(t, err)

	// Test command not found
	msg.Text = "/nonexistent command"
	err = bot.handleCommand(msg)
	assert.ErrorIs(t, err, models.ErrCommandNotFound)
}
