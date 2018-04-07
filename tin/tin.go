package main

import (
	"log"
	"os"
	"encoding/csv"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	errl := log.New(os.Stderr, "ERROR: ", 0)
	warnl := log.New(os.Stderr, "WARNING: ", 0)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		errl.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		errl.Panic(err)
	}
	w := csv.NewWriter(os.Stdout)
	w.UseCRLF = false

	for update := range updates {
		if update.Message == nil {
			continue
		}

		var data [4]string
		data[0] = strconv.FormatInt(update.Message.Chat.ID, 10)
		data[1] = strconv.Itoa(update.Message.From.ID)
		data[2] = update.Message.From.UserName
		data[3] = update.Message.Text
		if err := w.Write(data[:]); err != nil {
			warnl.Println("Failed to write message", err)
		}

		w.Flush()

		if err := w.Error(); err != nil {
			log.Fatal(err)
		}
	}
}
