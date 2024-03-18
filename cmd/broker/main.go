package main

import "github.com/samarthasthan/e-commerce/internal/broker"

func main() {
	broker.NewBrokerServer(":8000")
}
