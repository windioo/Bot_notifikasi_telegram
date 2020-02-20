package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	t "golang/notifikasi/config"
	"log"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err:= godotenv.Load("config.env")
	if  err != nil {
		log.Panic("Error loading .env file")
	}

	Token:= os.Getenv("Token")
	bot, err := tgbotapi.NewBotAPI(Token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	resp, err := http.Get(t.Parameter())
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
}
