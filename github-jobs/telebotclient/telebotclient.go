package telebotclient

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func check(src string, e error) {
	if e != nil {
		log.Fatal(src, e)
	}
}

func getToken() string {
	pwd, _ := os.Getwd()

	tok, err := ioutil.ReadFile(pwd + "/conf/token.conf")
	check("getToken: ", err)

	return string(tok)
}

func createNewBot() *tgbotapi.BotAPI {
	token := getToken()
	bot, err := tgbotapi.NewBotAPI(token)
	check("Creating new bot using token: ", err)
	return bot
}

func SendTelegramMessage(message string) {
	bot := createNewBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	check("Sending message: ", err)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "update.Message.Text")
		//		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		bot.Send(msg)
	}
}
