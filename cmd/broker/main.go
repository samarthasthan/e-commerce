package main

import (
	constants "github.com/samarthasthan/e-commerce"
	broker "github.com/samarthasthan/e-commerce/internal/broker/delivery/rest"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

var PORT string

func init() {
	PORT = env.GetEnv("BROKER_REST_PORT", "7000")
}

func main() {
	// Creating new custom logrus instance
	log := logger.NewLogger(constants.BROKER_LOGGER_NAME)

	s := broker.NewRestServer(log)
	s.RunServer(PORT)
}
