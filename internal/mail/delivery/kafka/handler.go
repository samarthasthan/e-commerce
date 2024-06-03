package kafka

import (
	"bytes"
	"encoding/json"
	"strconv"
	"time"

	"github.com/samarthasthan/e-commerce/pkg/kafka"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/pkg/models"
	"gopkg.in/gomail.v2"
)

type MailHandler struct {
	consumer *kafka.Consumer
	log      *logger.Logger

	SMTP_SERVER   string
	SMTP_PORT     string
	SMTP_LOGIN    string
	SMTP_PASSWORD string
}

func NewMailHandler(
	consumer *kafka.Consumer,
	log *logger.Logger,
	SMTP_SERVER string,
	SMTP_PORT string,
	SMTP_LOGIN string,
	SMTP_PASSWORD string,
) *MailHandler {
	return &MailHandler{
		consumer:      consumer,
		log:           log,
		SMTP_SERVER:   SMTP_SERVER,
		SMTP_PORT:     SMTP_PORT,
		SMTP_LOGIN:    SMTP_LOGIN,
		SMTP_PASSWORD: SMTP_PASSWORD,
	}
}

func (h *MailHandler) SendMails() error {
	h.consumer.Subscribe([]string{"mail"})
	for {
		msg, err := h.consumer.ReadMessage(5 * time.Second)
		if err != nil {
			// h.log.Errorf("Error reading message: %v", err)
			continue
		}
		go func() {
			mail := models.Mail{}
			bytes := bytes.NewBuffer(msg.Value)
			json.NewDecoder(bytes).Decode(&mail)
			h.SendMail(&mail)
			h.log.Infof("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		}()
	}

}

func (h *MailHandler) SendMail(in *models.Mail) error {
	from := h.SMTP_LOGIN                   // Sender email
	to := in.To                            // Recipient email
	host := h.SMTP_SERVER                  // SMTP server
	port, err := strconv.Atoi(h.SMTP_PORT) // SMTP port

	if err != nil {
		return err
	}

	// Message
	msg := gomail.NewMessage()
	msg.SetHeader("From", "noreply@samarthasthan.com") // Sender email
	msg.SetHeader("To", to)                            // Recipient email
	msg.SetHeader("Subject", in.Subject)               // Subject of the email
	// text/html for a html email
	msg.SetBody("text/html", in.Body) // Body of the email

	n := gomail.NewDialer(host, port, from, h.SMTP_PASSWORD) // SMTP server details

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return err
	}

	h.log.Infof("Mail sent to %s", to)
	return nil
}
