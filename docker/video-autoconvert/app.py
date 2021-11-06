import logging
import os
import random
import shutil
import subprocess
import sys
import time
from datetime import datetime, timedelta

from watchdog.events import FileCreatedEvent, FileSystemEventHandler
from watchdog.observers import Observer
from flask import Flask


DIR_TO_WATCH = '/autoconvert/'
recent_events = []


def ensure_dir(dir):
    if not os.path.exists(dir):
        os.mkdir(dir)


def ffmpeg(input_path):
    logging.info('inside ffmpeg')
    start = datetime.now()
    basename = os.path.basename(input_path)
    tmp_dir = os.path.join(DIR_TO_WATCH, 'tmp')
    ensure_dir(tmp_dir)
    tmp_path = os.path.join(tmp_dir, basename)
    vf = '"pad=width=3412:height=1920:x=1166:y=0:color=black"'
    other_flags = '-nostats -preset veryfast -codec:a copy'
    cmd = 'ffmpeg -i {} -vf {} {} {}'.format(
        input_path, vf, other_flags, tmp_path)
    logging.info(cmd)
    process = subprocess.Popen(cmd,
                               stdout=subprocess.PIPE, stderr=subprocess.PIPE,
                               shell=True)
    stdout, stderr = process.communicate()
    if stdout:
        logging.info('ffmpeg stdout: ' + str(stdout))
    if stderr:
        logging.info('ffmpeg stderr: ' + str(stderr))
    logging.info('duration: ' + str(datetime.now() - start))
    output_path = os.path.join('/video', basename)
    output_tmp = output_path + '.tmp'
    shutil.move(tmp_path, output_tmp)
    os.rename(output_tmp, output_path)


class Watcher(FileSystemEventHandler):
    def on_created(self, event):
        if type(event) == FileCreatedEvent:
            file = event.src_path
            recent_events.append('New file: ' + file + ', ' + os.path.basename(file))
            if file.endswith('.mp4'):
                ffmpeg(file)
        else:
            recent_events.append(str((type(event))))


def monitor():
    path = DIR_TO_WATCH
    e_handler = Watcher()
    watch = Observer()
    watch.schedule(e_handler, path, recursive=False)
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
    logging.basicConfig(level=logging.INFO)
    logging.info('inside main()')
    recent_events.append('this is a test')
    watch = monitor()
    # watch.stop()
    # watch.join()


main()
