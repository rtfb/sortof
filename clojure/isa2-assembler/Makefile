
# runs the container
all: dbuild
	docker run -it --name clj-nrepl --rm \
    --mount type=bind,source="$(shell pwd)",target=/home/rtfb \
    --net=host clj-nrepl:latest

# builds the docker image
.PHONY: dbuild
dbuild:
	docker build -t clj-nrepl .
