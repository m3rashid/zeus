FROM golang:1.21.5-alpine as builder

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

RUN cd common && go mod download

WORKDIR /app/auth

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /auth


##########################
# Making a smaller image #
##########################

FROM scratch

COPY --from=builder /auth /auth
COPY --from=builder /app/auth/.env /.env

ENTRYPOINT ["/auth"]
