package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	//  Global varible for BotToken
	telegramBotToken string
)
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Calculate"),
		tgbotapi.NewKeyboardButton("More info"),
	),
)

func init() {
	// Don't static port
	port := os.Getenv("PORT")

	// Goroutine for initialize port connection
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	// Get flag - telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "5281456176:AAH8pz8Rv-74_xUwKBwrujE8AxQ32O6zY-U", "Telegram Bot Token")
	flag.Parse()

	// Don't run without telegrambottoken
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

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
		// Default answer for any message
		reply := "Sorry, unknown command. Just use butttons bellow, or type <Calculate> or <More Info>"
		if update.Message == nil {
			continue
		}

		// Loging - from whom what message came
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Switch for commands from bot, wich starts from "/"
		switch update.Message.Command() {
		case "start":
			reply = "Hello! I help you to calculate your profit by staking in SpaceBot Project"
		case "calc":
			reply = "Function in progress..."
		case "info":
			reply = "Please, reed this instructions - https://teletype.in/@arousal/spacebot"
		}

		// Switch for buttons
		switch update.Message.Text {
		case "Calculate":
			reply = "Coming soon"
		case "More info":
			reply = "Please, reed this instructions - https://teletype.in/@arousal/spacebot"
		}

		// Create answer message
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		// TG keyboard
		msg.ReplyMarkup = numericKeyboard

		// Send message
		bot.Send(msg)
	}
}
