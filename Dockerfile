# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o=docker-go-app ./cmd

EXPOSE 8088

CMD [ "./docker-go-app" ]