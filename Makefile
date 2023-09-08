run:
	go run ./cmd

check:
	curl http://localhost:8088

build:
	go build -o=app ./cmd
	chmod +x app

run-binary:
	./app

dk-build:
	docker build --tag docker-go-app .
dk-run:
	docker run --publish=8088:8088 --name=docker-go-app docker-go-app 

PHONY: build run	