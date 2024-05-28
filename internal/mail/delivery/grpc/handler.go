package grpc

import (
	"context"

	"github.com/samarthasthan/e-commerce/proto_go"
	"gopkg.in/gomail.v2"
)

type MailHandler struct {
	proto_go.UnimplementedMailServiceServer
}

func NewMailHandler() *MailHandler {
	return &MailHandler{}
}

func (h *MailHandler) SendMail(ctx context.Context, in *proto_go.MailRequest) (*proto_go.MailResponse, error) {
	from := "use your own sender"
	to := in.Email
	host := "smtp-relay.sendinblue.com"
	port := 587
	msg := gomail.NewMessage()
	msg.SetHeader("From", "noreply@samarthasthan.com")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", "Test mail")
	// text/html for a html email
	msg.SetBody("text/html", in.Body)

	n := gomail.NewDialer(host, port, from, "use your own key")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		return &proto_go.MailResponse{Success: false, Message: err.Error()}, err
	}

	return &proto_go.MailResponse{Success: true, Message: "Mail send successfully!"}, nil
}
