package main

import (
	"finance-bot/internal/bot"
	"finance-bot/internal/config"
	"fmt"
)

func main() {
	config := config.NewEnvConfig()

	b := bot.NewBot(config)
	if err := b.Start(); err != nil {
		fmt.Println(err)
	}

}
