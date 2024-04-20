package main

import (
	constants "github.com/samarthasthan/e-commerce"
	grpc "github.com/samarthasthan/e-commerce/internal/authentication/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

func main() {
	// Creating new custom logrus instance
	log := logger.NewLogger(constants.AUTHENTICATION_LOGGER_NAME)

	

	// Creating new Authentication gRPC server
	s := grpc.NewAuthenticationGrpcServer(log)
	s.Run()
}
