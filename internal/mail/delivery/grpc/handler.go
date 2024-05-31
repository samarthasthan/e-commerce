package grpc

import (
	"context"
	"strconv"

	"github.com/samarthasthan/e-commerce/proto_go"
	"gopkg.in/gomail.v2"
)

type MailHandler struct {
	proto_go.UnimplementedMailServiceServer
	SMTP_SERVER   string
	SMTP_PORT     string
	SMTP_LOGIN    string
	SMTP_PASSWORD string
}

func NewMailHandler(
	SMTP_SERVER string,
	SMTP_PORT string,
	SMTP_LOGIN string,
	SMTP_PASSWORD string,
) *MailHandler {
	return &MailHandler{
		SMTP_SERVER:   SMTP_SERVER,
		SMTP_PORT:     SMTP_PORT,
		SMTP_LOGIN:    SMTP_LOGIN,
		SMTP_PASSWORD: SMTP_PASSWORD,
	}
}

func (h *MailHandler) SendMail(ctx context.Context, in *proto_go.MailRequest) (*proto_go.MailResponse, error) {
	from := h.SMTP_LOGIN
	to := in.Email
	host := h.SMTP_SERVER
	port, err := strconv.Atoi(h.SMTP_PORT)
	if err != nil {
		return nil, err
	}
	msg := gomail.NewMessage()
	msg.SetHeader("From", "noreply@samarthasthan.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", in.Subject)
	// text/html for a html email
	msg.SetBody("text/html", in.Body)

	n := gomail.NewDialer(host, port, from, h.SMTP_PASSWORD)

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return &proto_go.MailResponse{Success: false, Message: err.Error()}, err
	}

	return &proto_go.MailResponse{Success: true, Message: "Mail send successfully!"}, nil
}
