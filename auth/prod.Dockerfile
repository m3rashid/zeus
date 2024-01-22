FROM golang:1.21.5-alpine as builder

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . /app/auth/
COPY ../common /app/common

RUN cd common && go mod download && cd ..

WORKDIR /app/auth

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /auth


##########################
# Making a smaller image #
##########################

FROM scratch

COPY --from=builder /auth /auth

ENTRYPOINT ["/auth"]
