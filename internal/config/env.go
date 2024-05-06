package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewEnvConfig() *Config {
	if godotenv.Load() != nil {
		log.Println("No .env file found, setting default config values")
	}
	config := &Config{
		BotToken: getEnv("TELEGRAM_BOT_KEY", ""),
		LogLevel: getEnv("LOG_LEVEL", "DEBUG"),
	}
	return config
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
