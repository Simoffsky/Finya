package main

import (
	"finance-bot/internal/config"
	"finance-bot/pkg/log"
	"finance-bot/pkg/telebot"
	"fmt"
)

func main() {
	config := config.NewEnvConfig()

	logger := log.NewDefaultLogger(
		log.LevelFromString(config.LogLevel),
	).WithTimePrefix()

	bot := telebot.NewTelebot(config.BotToken, logger)

	user, err := bot.GetMe()
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Println(user)

}
