run:
	go run ./cmd

check:
	curl http://localhost:8088

build:
	go build -o=app ./cmd
	chmod +x app

run-binary:
	./app

PHONY: build run	