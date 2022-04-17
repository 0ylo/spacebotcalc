package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	//  Global varible for BotToken.
	telegramBotToken string = "TOKENHERE"
)
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Calculate"),
		tgbotapi.NewKeyboardButton("More info"),
	),
)

func main() {
	// Use token for create instance bot
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - varible for get updates from Tg server
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Get new message
	updates := bot.GetUpdatesChan(u)

	// Reed a "update" messeges from "updates" channel
	for update := range updates {

		if update.Message == nil {
			continue
		}

		// Loging - from whom what message came
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Create answer message
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Default answer for any message
		msg.Text = "Sorry, unknown command. Just use butttons bellow, or type <Calculate> or <More Info>"

		// TG keyboard
		msg.ReplyMarkup = numericKeyboard

		// Switch for commands from bot, wich starts from "/"
		switch update.Message.Command() {
		case "start":
			msg.Text = "Hello! I help you to calculate your profit by staking in SpaceBot Project"
		case "calc":
			msg.Text = "Function in progress..."
		case "info":
			msg.Text = "Please, reed this instructions - https://teletype.in/@arousal/spacebot"
		}

		// Switch for buttons
		switch update.Message.Text {
		case "Calculate":
			msg.Text = "Coming soon"
		case "More info":
			msg.Text = "Please, reed this instructions - https://teletype.in/@arousal/spacebot"
		}

		// Send message
		bot.Send(msg)
	}
}
