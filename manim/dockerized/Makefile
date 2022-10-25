
# builds the animation (run inside the container)
.PHONY: build
build:
	manim -ql s2c.py SquareToCircle

# builds the docker image
.PHONY: dbuild
dbuild:
	docker build -t manim-dev .

# runs the container
.PHONY: drun
drun:
	docker run -it --name manim --rm \
    --mount type=bind,source="$(shell pwd)",target=/home/rtfb/dev \
    --net=host manim-dev:latest

