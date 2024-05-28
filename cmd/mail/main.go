package main

import (
	"github.com/samarthasthan/e-commerce/internal/mail/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

var (
	MAIL_GRPC_PORT string
)

func init() {
	MAIL_GRPC_PORT = env.GetEnv("MAIL_GRPC_PORT", "12000")
}

func main() {
	log := logger.NewLogger("Mail")

	server := grpc.NewMailGrpcServer(log)
	server.Run(MAIL_GRPC_PORT)
}
