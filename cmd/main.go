package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type application struct{
	client *tbot.Client
}

var (
	app application
	bot *tbot.Server
	token string
)

func init() {
	e := godotenv.Load()
	if e != nil{
		log.Println(e)
	}
	token = os.Getenv("5281456176:AAH8pz8Rv-74_xUwKBwrujE8AxQ32O6zY-U
	")
}

func main() {
	bot = tbot.New(token)
	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	log.Fatal(bot.Start())
}

func (a *application) startHandler(m *tbot.Message){
	msg := "work!" a.client.SendMessage(m.Chat.ID, msg)
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