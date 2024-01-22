FROM golang:1.21.5-alpine as base

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . /app/
COPY ../common /app/

RUN cd common && go mod download && cd ..

WORKDIR /app/auth

RUN go mod download
