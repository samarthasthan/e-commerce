package main

import (
	mail "github.com/samarthasthan/e-commerce/internal/mail/delivery/kafka"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/kafka"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	tracer "github.com/samarthasthan/e-commerce/pkg/zipkin"
)

var (
	MAIL_GRPC_PORT string
	SMTP_SERVER    string
	SMTP_PORT      string
	SMTP_LOGIN     string
	SMTP_PASSWORD  string
	KAFKA_PORT     string
	KAFKA_HOST     string
)

func init() {
	MAIL_GRPC_PORT = env.GetEnv("MAIL_GRPC_PORT", "12000")
	SMTP_SERVER = env.GetEnv("SMTP_SERVER", "smtp-relay.sendinblue.com")
	SMTP_PORT = env.GetEnv("SMTP_PORT", "587")
	SMTP_LOGIN = env.GetEnv("SMTP_LOGIN", "use your own sender")
	SMTP_PASSWORD = env.GetEnv("SMTP_PASSWORD", "use your own key")
	KAFKA_PORT = env.GetEnv("KAFKA_PORT", "9092")
	KAFKA_HOST = env.GetEnv("KAFKA_HOST", "localhost")
}

func main() {
	// Initialising Custom Logger
	log := logger.NewLogger("Mail")

	// create a new Zipkin tracer
	_, err := tracer.NewTracer("mail", 12000)
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}

	// Initialising Kafka Consumer
	k := kafka.NewKafkaConsumer(KAFKA_HOST, KAFKA_PORT)

	// Initialising Mail Handler
	m := mail.NewMailHandler(k, log, SMTP_SERVER, SMTP_PORT, SMTP_LOGIN, SMTP_PASSWORD)
	// Start sending mails
	m.SendMails()
}
