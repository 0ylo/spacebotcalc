package main

import (
	"flag"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	webhook = "https://spacebotcalc.herokuapp.com/"
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)

func init() {
	// принимаем на входе флаг -telegrambottoken
	flag.StringVar(&telegramBotToken, "telegrambottoken", "5281456176:AAH8pz8Rv-74_xUwKBwrujE8AxQ32O6zY-U", "Telegram Bot Token")
	flag.Parse()

	// без него не запускаемся
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		reply := "Не знаю что сказать!"
		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот"
		case "hello":
			reply = "world"
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}

/*
Read heroku logs - "heroku logs --tail"

Deploy:
$ heroku login
$ cd /Users/anton/Documents/GoCode/GitHub/spacebotcalc
$ git init
$ heroku git:remote -a spacebotcalc

Deploy your application
$ git add .
$ git commit -am "initial commit"
$ git push heroku main
*/

/*
package main

import (
	"fmt"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

var ()

const (
	webhook = "https://spacebotcalc.herokuapp.com/"
)

func main() {
	// Hе статичный порт
	port := os.Getenv( key: "PORT")

	// Рутина для инициализации соединения по порту
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, handler: nil))
	}()

	// Падение в случае если не смогли создать бота + обработка ошибки
	bot, err := tgbotapi.NewBotAPI(tgTOKEN()
	if err != nil{
		log.Fatal("creation bot", err)
	}
	log.Println(("bot created")

	//webhook
	if _, err := bot.SetWebhook(tgbotapi.NewWebHook(webHook)); err != nil {log.Fatal(format: "settening webhook %v; error: %v", webHook, err)
	}
	log.Println(("webHook set")

	//слушаем url для принятия сообщений
	updates := bot.ListenForWebhook(pattern: "/")
	//создаем канал для общения с го рутин, чтобы общаться с разными пользователями
	for update := renge updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.From.Chat.ID, update.Message.Text)); err != nil {
			log.Print(err)
		}
	}

}
*/
