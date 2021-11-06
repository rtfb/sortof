import logging
import os
import random
import sys
import time

from watchdog.events import FileCreatedEvent, FileSystemEventHandler
from watchdog.observers import Observer
from flask import Flask


DIR_TO_WATCH = '/autoconvert/'
recent_events = []


class Watcher(FileSystemEventHandler):
    def on_created(self, event):
        if type(event) == FileCreatedEvent:
            file = event.src_path
            recent_events.append('New file: ' + file + ', ' + os.path.basename(file))
        else:
            recent_events.append(str((type(event))))


def monitor():
    path = DIR_TO_WATCH
    e_handler = Watcher()
    watch = Observer()
    watch.schedule(e_handler, path, recursive=True)
    watch.start()
    return watch


app = Flask(__name__)


def all_events(events):
    return '<br>'.join(sorted(events))


@app.route('/')
def hello():
    num = random.randint(1, 30)
    hello = 'Hello World! Here\'s a random number for you: {}.\n'.format(num)
    events = all_events(recent_events)
    return hello + '<p>Here are all recent events at {}:<br>{}</p>'.format(
        DIR_TO_WATCH, events)


def main():
    logging.log(logging.WARN, 'inside main()')
    recent_events.append('this is a test')
    watch = monitor()
    # watch.stop()
    # watch.join()


main()
