FROM golang:1.21.5-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY ../auth /app/auth
COPY ../common /app/common

RUN cd common && go mod download

WORKDIR /app/auth

RUN go mod download
