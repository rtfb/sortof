
# Based on https://twitter.com/francoisfleuret/status/1725418041570173334

# builds the docker image
.PHONY: dbuild
dbuild:
	docker build -t llm-test .

# runs the container
.PHONY: drun
drun:
	docker run -it --name llm-test --rm \
    --net=host llm-test:latest
