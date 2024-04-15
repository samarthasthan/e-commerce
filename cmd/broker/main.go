package main

import (
	broker "github.com/samarthasthan/e-commerce/internal/broker/delivery/rest"
	"github.com/samarthasthan/e-commerce/pkg/env"
)

var PORT string

func init() {
	PORT = env.GetEnv("BROKER_REST_PORT", "7000")
}

func main() {
	s := broker.NewRestServer()
	s.RunServer(PORT)
}
