package telegrambot

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func checkError(src string, e error) {
	if e != nil {
		log.Fatal(src, e)
	}
}

func getToken() string {
	pwd, _ := os.Getwd()

	tok, err := ioutil.ReadFile(pwd + "\\conf\\token.conf")
	checkError("Getting token from configuration file: ", err)

	return string(tok)
}

func createNewBot() *tgbotapi.BotAPI {
	token := getToken()
	bot, err := tgbotapi.NewBotAPI(token)
	checkError("Creating new bot using token: ", err)
	return bot
}

func SendMessage(message string) {
	bot := createNewBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	checkError("Sending Telegram message: ", err)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		bot.Send(msg)
	}
}
