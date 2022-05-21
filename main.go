package main

import (
	"log"
	"os"
	"tb_photo_loc/translit"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token, exist := os.LookupEnv("TRANSLIT_TOKEN")
	if !exist || len(token) < 10 {
		log.Fatalln("no env var | TRANSLIT_TOKEN")
	}
	bot, err := tg.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tg.NewUpdate(0)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			text := update.Message.Text
			chat := update.Message.Chat
			msg := tg.NewMessage(chat.ID, translit.EngToRu(text))
			_, err := bot.Send(msg)
			if err != nil {
				log.Printf("can't send message | %v", err)
			}
		}
	}
}
