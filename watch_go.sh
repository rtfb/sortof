#! /usr/bin/env sh

inotifywait -m -r -e close_write *.go | while read line
do
	go test
done
