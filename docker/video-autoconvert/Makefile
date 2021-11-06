
.PHONY: all
all: clean
	docker-compose build
	docker save -o autoconvert-image.tar video-autoconvert_convert

.PHONY: clean
clean:
	rm -f autoconvert-image.tar
