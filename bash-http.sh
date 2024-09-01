#!/usr/bin/env bash

# From here: https://www.youtube.com/watch?v=tSoFTD9Y8UU

req=(
    'GET / HTTP/1.1'
    'Host: rtfb.lt'
    'Connection: close'
    ''
)

exec 3<>/dev/tcp/rtfb.lt/80

printf '%s\r\n' "${req[@]}" >&3

while read -r data <&3; do
    echo "R: $data"
done
