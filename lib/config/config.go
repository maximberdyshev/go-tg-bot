package config

import (
	"go-tg-bot/lib/wrappers"
	"os"

	"github.com/joho/godotenv"
)

type TelegramConfig struct {
	BotToken string
	Host     string
}

type Config struct {
	Telegram TelegramConfig
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, wrappers.WrapErr("can't load .env", err)
	}

	token, ok := os.LookupEnv("TG_BOT_TOKEN")
	if !ok {
		return nil, wrappers.NewErr("can't load TG_BOT_TOKEN")
	}

	host, ok := os.LookupEnv("TG_BOT_HOST")
	if !ok {
		return nil, wrappers.NewErr("can't load TG_BOT_HOST")
	}

	return &Config{
		Telegram: TelegramConfig{
			BotToken: token,
			Host:     host,
		},
	}, nil
}
