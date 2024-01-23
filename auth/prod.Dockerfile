FROM golang:1.21.5-alpine as builder

RUN apk update && apk add git ca-certificates curl gnupg nodejs npm && npm i -g yarn

ENV USER=genos
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

RUN cd common && go mod download

WORKDIR /app/auth
RUN go mod download && cd web && yarn install && yarn build
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /auth


##############################################################################

FROM scratch

COPY --from=builder /auth /auth
COPY --from=builder /app/auth/.env /.env
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER genos:genos

ENTRYPOINT ["/auth"]
