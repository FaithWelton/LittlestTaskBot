package startup

import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func New() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		fmt.Printf("\nNewBotAPI Error:")
		panic(err)
	}

	bot.Debug = true

	fmt.Printf("\n[TASKBOT] Setup & Authorized on account: %s |", bot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := bot.GetUpdatesChan(updateConfig)
	return bot, updates
}
