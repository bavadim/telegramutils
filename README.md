# Telegramutils

The project implements the telegram messenger I/O interface through unix I/O streams. There are two simple programs in this repository: 

- tin - a program for reading messages sent by user to the bot. Messages are written to `stdout` in the following format: `chatId,"text",,,,`.
- tout - program for sending messages from the bot to the telegram user. Messages are read from `stdin` and must have the following format: `chatId,"text","file_path","button1 text","button2 text","button3 text"`. 

## Instalation

You can download binaries or install from sources.

The project is written on GO so you must have a customized GO environment for assembly a program from sources. Clone this repo and run `install.sh`. After that, `tin` and `tout` appear in the `{GOPATH}/bin`.

## Usage

You mast have a telegram token in `TELEGRAM_TOKEN` env variable before you start using the bot. This is example echo bot:

```
export TELEGRAM_TOKEN=<my token> 
tin | tout
```

This is example simple python echo bot.

In bot.py:

```
#!/usr/bin/env python3

import csv
import sys


writer = csv.writer(sys.stdout, quoting=csv.QUOTE_NONNUMERIC)
for row in csv.reader(iter(sys.stdin.readline, '')):
        chat_id, question = row[0], row[1]
        writer.writerow([chat_id, "echo: {}".format(question), "my.file", \
		"button 1", "button 2", "button 3"])
        sys.stdout.flush()                     
```

In command line:
```
export TELEGRAM_TOKEN=<my token>
tin | python3 ./bot.py | tout
```
	
