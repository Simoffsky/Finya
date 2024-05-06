package main

import (
	"finance-bot/internal/config"
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot"
)

func main() {
	config := config.NewEnvConfig()

	logger := log.NewDefaultLogger(
		log.LevelFromString(config.LogLevel),
	).WithTimePrefix()

	bot := telebot.NewTelebot(config.BotToken, logger)

	err := bot.LongPooling()
	if err != nil {
		logger.Error("failed to long pool: " + err.Error())
	}
}
