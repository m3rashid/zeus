FROM golang:1.21.5-alpine as builder

RUN apk update && apk add --no-cache git

ENV USER=appuser
ENV UID=10001

RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home --uid "${UID}" "${USER}"

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
COPY ../common common
COPY ../auth auth

RUN cd common && go mod download && go mod verify

WORKDIR /app/auth

RUN go mod download && go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /auth


##############################################################################

FROM scratch

COPY --from=builder /auth /auth
COPY --from=builder /app/auth/.env /.env
USER appuser:appuser

ENTRYPOINT ["/auth"]
