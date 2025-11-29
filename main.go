package main

import (
	"fmt"
	"os"
	"time"

	"auf.party.bot/api"
	"auf.party.bot/logger"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func main() {
	//load logger
	logger := logger.SetupLogger()
	logger.Info("Starting telegram bot...")

	//load .env
	err := godotenv.Load()
	if err != nil {
		logger.Warn("No .env file found, using system environment variables")
	}

	pref, err := loadConfig()
	if err != nil {
		logger.Error("Can not load config")
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logger.Error("Can not create bot because:", err)
		return
	}

	logger.Info("Telegram bot created")
	// fmt.Println("bot created")

	handlers := api.NewHandlers(logger)
	handlers.SetupHandlers(bot)

	logger.Info("Telegram bot started")
	// fmt.Println("Telegram bot started")
	bot.Start()
}

func loadConfig() (telebot.Settings, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" { //todo create normal env
		return telebot.Settings{}, fmt.Errorf("Environment variable 'BOT_TOKEN' should not be empty")
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	return pref, nil
}
