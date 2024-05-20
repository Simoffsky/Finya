package telebot

import (
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot/api"
	"finance-bot/pkg/telebot/models"
	"sync"
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

	tbot.handleUpdates(updates)

	assert.Equal(t, 0, tbot.offset)

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

	tbot.handleUpdates(updates)

	assert.Equal(t, 3, tbot.offset)

	updates = append(updates, models.Update{
		UpdateID: 3,
		Message:  nil,
	})

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := <-tbot.Errors()
		assert.EqualError(t, err.Error, models.ErrUpdateEmptyMessage.Error())
	}()

	tbot.handleUpdates(updates) // assert error in channel
	wg.Wait()
	assert.Equal(t, 4, tbot.offset)

	updates = []models.Update{}

	tbot.handleUpdates(updates)
	assert.Equal(t, 4, tbot.offset)

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
