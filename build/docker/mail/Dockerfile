FROM golang:latest

WORKDIR /usr/src/mail

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app ./cmd/mail/