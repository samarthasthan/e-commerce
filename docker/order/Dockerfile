FROM golang

WORKDIR /usr/src/order

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o app ./cmd/order/