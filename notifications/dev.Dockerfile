FROM golang:1.21.5-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY ../notifications /app/notifications
COPY ../common /app/common

RUN cd common && go mod download

WORKDIR /app/notifications

RUN go mod download
