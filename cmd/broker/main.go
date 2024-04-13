package main

import "github.com/samarthasthan/e-commerce/internal/broker"

func main() {
	// RestFull API broker server
	b := broker.NewBrokerServer()
	b.Connect()
}
