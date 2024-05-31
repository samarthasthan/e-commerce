FROM golang:1.22.1-alpine3.19

WORKDIR /usr/src/mail

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app ./cmd/mail/