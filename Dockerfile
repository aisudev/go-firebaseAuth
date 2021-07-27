FROM golang:alpine

WORKDIR /app

ADD . /app

CMD go mod tidy

CMD go run main.go
