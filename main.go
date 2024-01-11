package main

import (
	"fmt"
	"strconv"
	"strings"
	"telebot/dice"
	"telebot/greetings"
	"telebot/help"
	"telebot/startup"
	"telebot/weather"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func processMessage(update tgbotapi.Update) (string, string, string) {
	return update.Message.From.UserName, update.Message.Text, update.Message.From.LanguageCode
}

func main() {
	bot, updates := startup.New()

	for update := range updates {
		if update.Message == nil {
			fmt.Println("[LTB]: NON MESSAGE RECEIVED")
			continue // Ignore NON MESSAGES
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		username, text, languageCode := processMessage(update)
		fmt.Printf("\n[LTB] New Message Received From %s: %s ", username, text)

		weather, err := weather.New(languageCode)
		if err != nil {
			fmt.Printf("\nWeather Error:")
			panic(err)
		}

		if update.Message.IsCommand() { // HANDLE COMMANDS
			command := update.Message.Command()

			switch command {
			case "start":
				greetingMessage, err := greetings.Hello(username)
				if err != nil {
					fmt.Println(err)
					msg.Text = "Sorry I seem to have lost my greeting!"
					break
				}

				helpMessage, err := help.New()
				if err != nil {
					fmt.Println(err)
					msg.Text = "Oh no the help is broken!"
					break
				}

				msg.Text = fmt.Sprintf("%s\n\n%s", greetingMessage, helpMessage)
			case "hello":
				message, err := greetings.Hello(username)
				if err != nil {
					fmt.Println(err)
					msg.Text = "Sorry I seem to have lost my greeting!"
					break
				}

				msg.Text = message
			case "weather":
				msg.Text = "Sure Thing! What is your location so I can find the nearest weather station"
				btn := tgbotapi.NewKeyboardButtonLocation("Send Location to Taskbot")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
			case "roll":
				parts := strings.Split(text, " ")
				if len(parts) != 2 {
					msg.Text = "Sorry I don't understand, Please format your command like: /dice #d#. \nWith the first # being the number of die, and the second # being the number of sides."
					break
				}

				diceParts := strings.Split(parts[1], "d")
				numDice, err := strconv.Atoi(diceParts[0])
				if err != nil {
					msg.Text = "You've given an invalid number of dice, please try again with the format: /dice #d#"
					break
				}

				numSides, err := strconv.Atoi(diceParts[1])
				if err != nil {
					msg.Text = "You've given an invalid number of sides for your dice, please try again with the format: /dice #d#"
					break
				}

				message, err := dice.Roll(numDice, numSides)
				if err != nil {
					fmt.Println(err)
					msg.Text = "OOPS The dice fell off the table!"
					break
				}

				msg.Text = message
			case "test":
				msg.Text = "Testing testing"
			case "help":
				helpMessage, err := help.New()
				if err != nil {
					fmt.Println(err)
					msg.Text = "Oh no the help is broken!"
					break
				}

				msg.Text = helpMessage
			default:
				msg.Text = fmt.Sprintf("Sorry I don't know the command %s 😢", text)
			}
		} else { // HANDLE NON COMMANDS
			fmt.Println("\n[LTB] NON COMMAND MESSAGE: ")
			fmt.Println(update.Message)
			fmt.Println(update.Message.Text)

			// Handle Weather / Location
			if update.Message.Location != nil {
				language := update.Message.From.LanguageCode
				location := map[string]float64{
					"lon": update.Message.Location.Longitude,
					"lat": update.Message.Location.Latitude,
				}

				weather, err := weather.Get(location, language)
				if err != nil {
					fmt.Println(err)
					msg.Text = "Sorry I'm having trouble finding the weather for your location, try again!"
					break
				}

				msg.Text = weather
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			}
		}

		if _, err := bot.Send(msg); err != nil {
			fmt.Printf("\nBot.Send Error:")
			panic(err)
		}

		continue
	}
}
