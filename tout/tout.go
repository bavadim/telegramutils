package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"fmt"
	"os"
)

func main() {
	errl := log.New(os.Stderr, "ERROR: ", 0)
	warnl := log.New(os.Stderr, "WARNING: ", 0)
	infol := log.New(os.Stdout, "INFO: ", 0)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		errl.Panic(err)
	}

	infol.Printf("Authorized on account %s", bot.Self.UserName)

	for {
		var chatId int64
		var text string
		n, err := fmt.Scanf("%d,%q\n", &chatId, &text)
		if err != nil || n != 2 {
			warnl.Println("Input must be chatId, text.", err)
			continue
		}

		msg := tgbotapi.NewMessage(chatId, text)
		bot.Send(msg)
	}
}
