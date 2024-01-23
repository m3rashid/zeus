FROM golang:1.21.5-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

RUN cat <<EOF >> go.work
go 1.21.5

use (
	"./notifications"
	"./common"
)
EOF
COPY go.work.sum go.work.sum
COPY ../common common
COPY ../notifications notifications

RUN cd common && go mod download

WORKDIR /app/notifications

RUN go mod download
