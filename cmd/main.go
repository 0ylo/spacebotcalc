//*
package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
//webhook = "https://spacebotcalc.herokuapp.com/"
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)

func init() {
	// Hе статичный порт
	port := os.Getenv("PORT")
	//log.Print("Listening on :" + port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))

	//address := fmt.Sprintf("%s:%s", "0.0.0.0", port)
	//fmt.Println(address)

	// Рутина для инициализации соединения по порту
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

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
		case "help":
			reply = "Can't help right now..."
		case "calc":
			reply = "Let's see.."
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)

		// отправляем
		bot.Send(msg)
	}
}

/*
Use go get URLofRep - if some imports doesn't work
.
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
//_______________________________________

package main

import (
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"os"
)

var ()

const (
	webHook = "https://spacebotcalc.herokuapp.com/"
)

func main() {
	// Hе статичный порт
	port := os.Getenv("PORT")

	// Рутина для инициализации соединения по порту
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	// Падение в случае если не смогли создать бота + обработка ошибки
	var tgTOKEN string = "5281456176:AAH8pz8Rv-74_xUwKBwrujE8AxQ32O6zY-U"
	bot, err := tgbotapi.NewBotAPI(tgTOKEN)
	if err != nil {
		log.Fatal("creation bot", err)
	}
	log.Println("bot created")

	//webhook
	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatal("settening webhook %v; error: %v", webHook, err)
	}
	log.Println("webHook set")

	//слушаем url для принятия сообщений
	updates := bot.ListenForWebhook("/")
	//создаем канал для общения с го рутин, чтобы общаться с разными пользователями
	for update := range updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil {
			log.Print(err)
		}
	}
}

//Just clear Webhook setting with
//curl -F "url=" https://api.telegram.org/botYOUR_TOKEN/setWebhook

*/
