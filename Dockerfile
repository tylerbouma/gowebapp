# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY ./assets/index.html /app/assets/

RUN go build -o /birdpedia

CMD [ "/birdpedia" ]