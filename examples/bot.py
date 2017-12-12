#!/usr/bin/env python3

import csv
import sys


writer = csv.writer(sys.stdout, quoting=csv.QUOTE_NONNUMERIC)
for row in csv.reader(iter(sys.stdin.readline, '')):
        chat_id, question = row[0], row[1]
        writer.writerow([chat_id, "echo: {}".format(question), "my.file", "button 1", "button 2", "button 3"])
        sys.stdout.flush()
