package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
	"os"
)

func main() {
	errl := log.New(os.Stderr, "ERROR: ", 0)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		errl.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		fmt.Printf("%d,%q\n", update.Message.Chat.ID, update.Message.Text)
	}
}
