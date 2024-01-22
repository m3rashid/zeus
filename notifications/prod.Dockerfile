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

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /notifications


##############################################################################

FROM scratch

COPY --from=builder /notifications /notifications
COPY --from=builder /app/notifications/.env /.env
USER appuser:appuser

ENTRYPOINT ["/notifications"]
