package main

import (
	"log"
	//"github.com/go-telegram-bot-api/telegram-bot-api"
	"gopkg.in/telegram-bot-api.v4"
	"os"
	"encoding/csv"
	"bufio"
	"strconv"
	"io"
	"flag"
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
		chunks = append(chunks, string(buf[:len(buf)]))
	}
	return []string(chunks)
}

func send(chatId int64, text string, warnl *log.Logger, bot *tgbotapi.BotAPI, mdSupport bool) {
  const maxSize = 4096

  for _, m := range split(text, maxSize) {
    msg := tgbotapi.NewMessage(chatId, m)
    if mdSupport == true {
    	msg.ParseMode = tgbotapi.ModeMarkdown
    }
    _, err := bot.Send(msg)
    if err != nil {
      warnl.Println("Can't to send message.", err)
      break
    }
  }
}

func main() {
  errl := log.New(os.Stderr, "ERROR: ", 0)
  warnl := log.New(os.Stderr, "WARNING: ", 0)

  bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
  if err != nil {
    errl.Panic(err)
  }

  isMarkdownPtr := flag.Bool("m", false, "markdawn support")
  flag.Parse()

  reader := bufio.NewReader(os.Stdin)
  r := csv.NewReader(reader)
  for {
    record, err := r.Read()
    if err == io.EOF {
      break
    }
	  
    chatId, err := strconv.ParseInt(record[0], 10, 64)
    if err != nil {
      warnl.Println("ChatId must be integer.", err)
      continue
    }
	  
    if err != nil || len(record) == 2 {
	    text := record[1]
           send(chatId, text, warnl, bot, *isMarkdownPtr)
    }

    if err != nil || len(record) == 3 {
	    text := record[1]
  	    tbtn1 := record[2]
	    msg := tgbotapi.NewMessage(chatId, text)
            msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
	                              tgbotapi.NewKeyboardButtonRow(
		                               tgbotapi.NewKeyboardButton(tbtn1),		               		               
                                                ),
                                       )
		bot.Send(msg)
    }
	
    if err != nil || len(record) == 4 {
	    text := record[1]
            tbtn1 := record[2]	
	    tbtn2 := record[3]			
	    msg := tgbotapi.NewMessage(chatId, text)
            msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
	                               tgbotapi.NewKeyboardButtonRow(
		                                tgbotapi.NewKeyboardButton(tbtn1),
		                                tgbotapi.NewKeyboardButton(tbtn2),		               
                                        ),
                                     )
		bot.Send(msg)
	}
	
	if err != nil || len(record) == 5 {
	        text := record[1]	
                tbtn1 := record[2]	
	        tbtn2 := record[3]	
                tbtn3 := record[4]		
		msg := tgbotapi.NewMessage(chatId, text)
                msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
	        tgbotapi.NewKeyboardButtonRow(
		               tgbotapi.NewKeyboardButton(tbtn1),
		               tgbotapi.NewKeyboardButton(tbtn2),
		               tgbotapi.NewKeyboardButton(tbtn3),
                         	),
                          )
		bot.Send(msg)
	}
  }
}
