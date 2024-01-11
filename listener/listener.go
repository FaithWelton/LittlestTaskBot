package listener

import (
	"fmt"
	"ltb/dice"
	"ltb/greetings"
	"ltb/help"
	"ltb/weather"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func processMessage(update tgbotapi.Update) (string, string) {
	return update.Message.From.UserName, update.Message.Text
}

func Listen(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel, weather *weather.Weather) {
	fmt.Printf("[LTB]: Listening...\n")

	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("1", "1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
		),
	)

	for update := range updates {
		if update.Message == nil {
			if update.CallbackQuery != nil {
				// Part of Test Command - testing what I can do with the callback, this just replies to the user with the number they choose on the keyboard
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := bot.Request(callback); err != nil {
					panic(err)
				}

				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				if _, err := bot.Send(msg); err != nil {
					panic(err)
				}
				continue
			}

			// handle NON MESSAGE updates
			fmt.Println("[LTB]: NEW NON MESSAGE RECEIVED")
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I AM A DEFAULT MESSAGE")
		username, text := processMessage(update)
		fmt.Printf("[LTB]: New Message Received From %s:\n%s\n", username, text)

		if update.Message.IsCommand() {
			// Handle Commands
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
				greetingMessage, err := greetings.Hello(username)
				if err != nil {
					fmt.Println(err)
					msg.Text = "Sorry I seem to have lost my greeting!"
					break
				}

				msg.Text = greetingMessage
			case "help":
				helpMessage, err := help.New()
				if err != nil {
					fmt.Println(err)
					msg.Text = "Oh no the help is broken!"
					break
				}

				msg.Text = helpMessage
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
			case "weather":
				msg.Text = "Sure Thing! What is your location so I can find the nearest weather station"
				btn := tgbotapi.NewKeyboardButtonLocation("Send Location to Taskbot 📍")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
			case "test": // Currently testing inline keyboard things
				msg.ReplyMarkup = numericKeyboard
			case "settings":
				msg.Text = "Settings have not been implemented yet, sorry!"
			default:
				msg.Text = fmt.Sprintf("Sorry I don't know the command %s 😢", text)
			}
		} else {
			// Handle NON COMMANDS

			// Weather / Location
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
			fmt.Printf("[ERROR]: SEND ERROR\n")
			panic(err)
		}

		continue
	}
}
