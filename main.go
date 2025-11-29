package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"auf.party.bot/logger"
	"gopkg.in/telebot.v3"
)

func main() {
	logger := logger.SetupLogger()
	logger.Info("Starting telegram bot...")

	pref, err := loadConfig()
	if err != nil {
		logger.Error("Can not load config")
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal("Can not create bot because:", err)
		return
	}

	fmt.Println("bot created")

	bot.Handle("/start", func(c telebot.Context) error {
		return c.Send("hello my friend")
	})

	fmt.Println("Bot started...")
	bot.Start()
}

func loadConfig() (telebot.Settings, error) {
	token := os.Getenv("BOT_TOKEN")
	if token == "" { //todo create normal env
		token = ""
		//return telebot.Settings{}, fmt.Errorf("Environment variable 'BOT_TOKEN' should not be empty") // todo error
	}

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	return pref, nil
}
