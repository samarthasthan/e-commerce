package main

import (
	"fmt"
	"log"

	"github.com/samarthasthan/e-commerce-backend/internal/broker"
)

func init() {
	fmt.Println("Broker listening on PORT 8000")
}

func main() {
	// Initialize the broker server
	brokerServer := broker.NewBrokerServer()

	//Routes
	brokerServer.HandleRoutes()

	// Start the server
	err := brokerServer.Run(":8000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
