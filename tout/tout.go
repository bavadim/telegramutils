package main

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
	"encoding/csv"
	"bufio"
	"strconv"
	"io"
	"flag"
)

func sendFile(chatId int64, url string, filename string, bot *tgbotapi.BotAPI){
                f, _ := os.Open(url)		 				  
	        reader := tgbotapi.FileReader{Name: filename, Reader: f, Size: -1}                
		msg := tgbotapi.NewDocumentUpload(chatId, reader)			
		bot.Send(msg)	           	  
	        f.Close()	
}

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

func send(chatId int64,
          text string,
          buttons tgbotapi.ReplyKeyboardMarkup,
          warnl *log.Logger,
          bot *tgbotapi.BotAPI,
          mdSupport bool) {
  const maxSize = 4096

  chunks := split(text, maxSize)
  lastIndex := len(chunks) - 1
  for i, m := range chunks {
    msg := tgbotapi.NewMessage(chatId, m)
    if mdSupport == true {
      msg.ParseMode = tgbotapi.ModeMarkdown
    }
    if i == lastIndex && len(buttons.Keyboard) > 0 {
      msg.ReplyMarkup = buttons
    } else {
      msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(false)
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
    if err != nil || len(record) != 5 {
      warnl.Println("Input must have csv format, for example: chatId,text,button1,button2,button3.", err)
      continue
    }
    chatId, err := strconv.ParseInt(record[0], 10, 64)
    if err != nil {
      warnl.Println("ChatId must be integer.", err)
      continue
    }
    text := record[1]

    buttons := make([][]tgbotapi.KeyboardButton, 3)
    k := 0
    for i := 2; i < 5; i++ {
      if record[i] != "" {
        buttons[k] = tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton(record[i]))
        k++
      }
    }
    buttons = buttons[0:k]

    send(chatId, text, tgbotapi.NewReplyKeyboard(buttons...), warnl, bot, *isMarkdownPtr)
  }
}
