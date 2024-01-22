FROM golang:1.21.5-alpine as builder

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.work go.work
COPY go.work.sum go.work.sum
COPY ../auth auth
COPY ../common common
COPY ../notifications notifications

RUN cd common && go mod download

WORKDIR /app/notifications

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /notifications


##########################
# Making a smaller image #
##########################

FROM scratch

COPY --from=builder /notifications /notifications
COPY --from=builder /app/notifications/.env /.env

ENTRYPOINT ["/notifications"]
