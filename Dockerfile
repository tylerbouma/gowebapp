# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY assets/ ./assets/

RUN go mod download

COPY *.go ./

RUN go build -o /birdpedia

CMD [ "/birdpedia" ]