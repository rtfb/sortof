# syntax=docker/dockerfile:1
FROM python:3.7-alpine
WORKDIR /code
ARG group=users
ARG gid=100
ARG user=rtfb
ARG uid=1026
ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0
RUN apk add --no-cache gcc musl-dev linux-headers ffmpeg
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
# RUN addgroup -g ${gid} ${group}
RUN adduser -u ${uid} -G ${group} -s /bin/sh -D ${user}
EXPOSE 5000
COPY . .
CMD ["flask", "run"]
