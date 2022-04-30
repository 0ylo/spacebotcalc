package main

import (
	"log"

	"github.com/0ylo/spacebotcalc/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	version string = "unknown"
	build   string = "unknown"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Calculate"),
		tgbotapi.NewKeyboardButton("More info"),
	),
)

func main() {
	cfg, err := config.Init(version, build)
	if err != nil {
		log.Fatalf("Read config: %v", err)
	}
	if cfg == nil {
		log.Fatal("Config is empty")
	}

	if len(cfg.Bot.Token) == 0 {
		log.Fatal("Please, configure the Telegram token either in the config or in environment variable SPACEBOTCALC_TOKEN")
	}

	log.Printf("Version: %s, build: %s", cfg.Version, cfg.Build)
	log.Println("Starting Spacebotcalc")
	// Use token for create instance bot
	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - varible for get updates from Tg server
	u := tgbotapi.NewUpdate(0)
	u.Timeout = cfg.Bot.Timeout

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
