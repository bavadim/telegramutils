# Telegramutils

This project inspired by linux utils. There are two simple programs in this repository: 

- tin - a program for reading messages sent by user to the bot. Messages are written to `stdout` in the following format: `chatId,"text"`.
- tout - program for sending messages from the bot to the telegram user. Messages are read from `stdin` and must have the following format: `chatId,"text"`. 

## Instalation from sources

The project is written on GO so you must have a customized GO environment. Clone this repo and run `install.sh`. After that, `tin` and `tout` appear in the `{GOPATH}/bin`.

For tout.go :

go get "github.com/go-telegram-bot-api/telegram-bot-api"

## Usage

You mast have a telegram token in `TELEGRAM_TOKEN` env variable before you start using the bot. This is example echo bot:

```
	export TELEGRAM_TOKEN=<my token> 
	tin | tout
```
