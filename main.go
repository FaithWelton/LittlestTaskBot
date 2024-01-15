package main

import (
	"fmt"
	"os"
	"taskbot/listener"
	"taskbot/weather"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		fmt.Printf("[ERROR] NewBotAPI Error:")
		panic(err)
	}

	bot.Debug = true

	fmt.Printf("[LTB]: Setup & Authorized on account: %s\n", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	weather, err := weather.New()
	if err != nil {
		fmt.Printf("\nWeather Error:")
		panic(err)
	}

	listener.Listen(bot, updates, weather)
}
