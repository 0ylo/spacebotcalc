package main

import (
	"fmt"
	//tgbotapi "gopkg.in/telegram-bot-api.v4"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

var ()

const (
	webhook = "https://spacebotcalc.herokuapp.com/"
)

func main() {
	//не статичный порт
	port := os.Getenv( key: "PORT")

	//Рутина для инициализации соединения по порту
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, handler: nil))
	}()

	// Падение в случае если не смогли создать бота + обработка ошибки
	bot, err := tgbotapi.NewBotAPI(tgTOKEN(/*nen*/)
	if err != nil{
		log.Fatal(/*nen*/ "creation bot", err)
	}
	log.Println((/*nen*/ "bot created")

	//webhook
	if _, err := bot.SetWebhook(tgbotapi.NewWebHook(webHook)); err != nil {log.Fatal(format: "settening webhook %v; error: %v", webHook, err)
	}
	log.Println((/*nen*/ "webHook set")
	
	//слушаем url для принятия сообщений
	updates := bot.ListenForWebhook(pattern: "/")
	//создаем канал для общения с го рутин, чтобы общаться с разными пользователями
	for update := renge updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.From.Chat.ID, /*возвращаемое сообщение*/ update.Message.Text)); err != nil {
			log.Print(err)
		}
	}

}
