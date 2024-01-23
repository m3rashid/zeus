FROM golang:1.21.5-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

RUN cat <<EOF >> go.work
go 1.21.5

use (
	"./auth"
	"./common"
)
EOF
COPY ./go.work.sum go.work.sum
COPY ../auth auth
COPY ../common common

RUN cd common && go mod download

WORKDIR /app/auth

RUN go mod download
