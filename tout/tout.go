package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"encoding/csv"
	"bufio"
	"strconv"
	"io"
)

func main() {
	errl := log.New(os.Stderr, "ERROR: ", 0)
	warnl := log.New(os.Stderr, "WARNING: ", 0)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		errl.Panic(err)
	}

	reader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(reader)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) != 2 {
			warnl.Println("Input must have csv format ([chatId, text]).", err)
			continue
		}
		chatId, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			warnl.Println("ChatId must be integer.", err)
			continue
		}
		text := record[1]
		warnl.Println(text)
		msg := tgbotapi.NewMessage(chatId, text)
		_, err = bot.Send(msg)
		if err != nil {
			warnl.Println("Can't to send message.", err)
			continue
		}
	}
}
