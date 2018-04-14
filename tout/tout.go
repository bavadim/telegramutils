package main

import (
	"log"
	"os"
	"encoding/csv"
	"bufio"
	"strconv"
	"io"
	"flag"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func split(str string, lim int) []string {
	buf := []rune(str)
	var chunk []rune
	chunks := make([]string, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, string(chunk))
	}
	if len(buf) > 0 {
		chunks = append(chunks, string(buf))
	}
	return []string(chunks)
}

func send(id int64, text string, buttons tgbotapi.ReplyKeyboardMarkup,
	warnl *log.Logger, bot *tgbotapi.BotAPI, md bool) {
	const maxSize = 4096
	chunks := split(text, maxSize)
	for i, m := range chunks {
		msg := tgbotapi.NewMessage(id, m)
		if md == true {
			msg.ParseMode = tgbotapi.ModeMarkdown
		}
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
		if i == len(chunks)-1 && len(buttons.Keyboard) > 0 {
			msg.ReplyMarkup = buttons
		}
		_, err := bot.Send(msg)
		if err != nil {
			if md == true {
				send(id, text, buttons, warnl, bot, false)
				break
			} else {
				warnl.Println("Failed to send message", err)
				break
			}
		}
	}
}

func main() {
	errl := log.New(os.Stderr, "ERROR: ", 0)
	warnl := log.New(os.Stderr, "WARNING: ", 0)
	bot, err := tgbotapi.NewBotAPI("567524377:AAHfNjID00ESGRJ4vodDA05-gDmmcvY1mYM")
	if err != nil {
		errl.Panic(err)
	}
	md := flag.Bool("m", false, "Enables Markdown support")
	flag.Parse()
	reader := csv.NewReader(bufio.NewReader(os.Stdin))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(record) != 6 {
			warnl.Println("Input must have following format: chatId,text,file,button1,button2,button3", err)
			continue
		}
		chatId, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			warnl.Println("chatId must be integer", err)
			continue
		}
		text := record[1]
		buttons := make([][]tgbotapi.KeyboardButton, 3)
		k := 0
		for i := 3; i < 6; i++ {
			if record[i] != "" {
				buttons[k] = tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(record[i]))
				k++
			}
		}
		buttons = buttons[0:k]
		send(chatId, text, tgbotapi.NewReplyKeyboard(buttons...), warnl, bot, *md)
		if record[2] != "" {
			fDescr, err := os.Open(record[3])
			if err != nil {
				warnl.Println("Check path to the file", err)
				continue
			}
			fReader := tgbotapi.FileReader{Name: record[3], Reader: fDescr, Size: -1}
			fMsg := tgbotapi.NewDocumentUpload(chatId, fReader)
			_, err = bot.Send(fMsg)
			if err != nil {
				warnl.Println("Failed to send file", err)
				fDescr.Close()
				break
			}
			fDescr.Close()
		}
	}
}
