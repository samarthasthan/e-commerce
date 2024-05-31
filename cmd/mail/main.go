package main

import (
	"github.com/samarthasthan/e-commerce/internal/mail/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/logger"
)

var (
	MAIL_GRPC_PORT string
	SMTP_SERVER    string
	SMTP_PORT      string
	SMTP_LOGIN     string
	SMTP_PASSWORD  string
)

func init() {
	MAIL_GRPC_PORT = env.GetEnv("MAIL_GRPC_PORT", "12000")
	SMTP_SERVER = env.GetEnv("SMTP_SERVER", "smtp-relay.sendinblue.com")
	SMTP_PORT = env.GetEnv("SMTP_PORT", "587")
	SMTP_LOGIN = env.GetEnv("SMTP_LOGIN", "use your own sender")
	SMTP_PASSWORD = env.GetEnv("SMTP_PASSWORD", "use your own key")
}

func main() {
	log := logger.NewLogger("Mail")

	server := grpc.NewMailGrpcServer(log)
	server.Run(MAIL_GRPC_PORT, SMTP_SERVER, SMTP_PORT, SMTP_LOGIN, SMTP_PASSWORD)
}
