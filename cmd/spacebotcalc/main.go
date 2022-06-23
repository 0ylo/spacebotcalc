package main

import (
	"log"
	"strconv"

	"github.com/0ylo/spacebotcalc/internal/commands"
	"github.com/0ylo/spacebotcalc/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DataCalc struct {
	State    int //1-Deposit, 2-Duration
	Name     string
	Deposit  string
	Duration string
}

var (
	dataCalcMap map[int64]*DataCalc
	dc          *DataCalc
	version     string = "unknown"
	build       string = "unknown"
)

func init() {
	dataCalcMap = make(map[int64]*DataCalc)
}

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💰 Calculate"),
		tgbotapi.NewKeyboardButton("📢 More info"),
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

		// ignore any non-Message updates
		if update.Message == nil {
			continue
		}

		// For commands
		if update.Message.IsCommand() {
			cmdText := update.Message.Command()
			if cmdText == "start" {
				msgConfig := tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"Hello! I help you to calculate your profit by staking in SpaceBot Project")
				msgConfig.ReplyMarkup = mainMenu
				bot.Send(msgConfig)
			} else if cmdText == "info" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, reed this instructions - https://teletype.in/@arousal/spacebot")
				bot.Send(msg)
			}
		} else {
			// For messages

			if update.Message.Text == mainMenu.Keyboard[0][0].Text {

				dataCalcMap[update.Message.From.ID] = new(DataCalc)
				dataCalcMap[update.Message.From.ID].State = 1

				msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите, пожалуйста, сумму депозита 👇")
				bot.Send(msgConfig)
			} else {
				dc, ok := dataCalcMap[update.Message.From.ID]
				if ok {
					if dc.State == 1 {
						dc.Deposit = update.Message.Text
						msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите, пожалуйста, срок депозита 👇")
						bot.Send(msgConfig)
						dc.State = 2
					} else if dc.State == 2 {
						dc.Duration = update.Message.Text

						// ----
						de, err := strconv.ParseFloat(dc.Deposit, 64)
						log.Printf("Convert f, g:", de, err)
						du, err := strconv.Atoi(dc.Duration)
						log.Printf("Convert f, g:", du, err)

						res1, res2 := commands.Calculate(de, du)
						log.Print("res1&2 is ", res1, res2)
						// ----

						msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, "Считаю...")
						bot.Send(msgConfig)
						delete(dataCalcMap, update.Message.From.ID)
					}
				} else {
					// other messages
					if update.Message.Text != mainMenu.Keyboard[0][1].Text {
						msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, "Нет такой команды, нажмите, пожалуйста, на одну из кнопок ниже👇")
						bot.Send(msgConfig)
					}
				}
			}

			if update.Message.Text == mainMenu.Keyboard[0][1].Text {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, reed this instructions - https://teletype.in/@arousal/spacebot")
				bot.Send(msg)
			}
		}

		/*
			// Switch for commands from bot, wich starts from "/"
			switch update.Message.Command() {
			case "start":
				msgConfig := tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"Hello! I help you to calculate your profit by staking in SpaceBot Project")
				msgConfig.ReplyMarkup = mainMenu
				bot.Send(msgConfig)
			case "info":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please, reed this instructions - https://teletype.in/@arousal/spacebot")
				bot.Send(msg)
			}
		*/

	}
}
