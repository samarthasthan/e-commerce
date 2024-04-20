package main

import (
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

func main() {
	log := logger.NewLogger("authentication")

	db, err := database.NewDatabase("root:password@tcp(localhost:8001)/authentication")
	if err != nil {
		log.Error("Failed to initialize the database:", err)
		return
	}
	defer db.Close()

	server := grpc.NewAuthenticationGrpcServer(log, db)
	server.Run()
}
