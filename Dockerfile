FROM golang:alpine

WORKDIR /app

ADD . /app

CMD go run main.go
