package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
	"github.com/samarthasthan/e-commerce/pkg/bcrpyt"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type AuthenticationHandler struct {
	proto_go.UnimplementedAuthenticationServiceServer
	mailClient proto_go.MailServiceClient
	mysql      database.Database
	redis      database.Database
}

func NewAuthenticationHandler(mysql database.Database, redis database.Database, mailClient proto_go.MailServiceClient) *AuthenticationHandler {
	if mysql == nil {
		panic("mysql dependency must not be nil")
	}
	if redis == nil {
		panic("redis dependency must not be nil")
	}
	if mailClient == nil {
		panic("mailClient dependency must not be nil")
	}
	return &AuthenticationHandler{
		mysql:      mysql,
		redis:      redis,
		mailClient: mailClient,
	}
}

// SignUp handles the SignUp gRPC request
func (h *AuthenticationHandler) SignUp(ctx context.Context, in *proto_go.SignUpRequest) (*proto_go.SignUpResponse, error) {

	mysql, ok := h.mysql.(*database.MySQL)
	if !ok {
		return nil, fmt.Errorf("mysql is not of type *database.MySQL")
	}

	role, err := mysql.Queries.GetRole(ctx, in.RoleName)

	if err != nil {
		return nil, err
	}

	// generate hashed password
	pass, err := bcrpyt.HashPassword(in.Password)

	if err != nil {
		return nil, err
	}

	uuid := uuid.New()
	// Execute CreateAccount query using sqlc
	_, err = mysql.Queries.CreateAccount(ctx, sqlc.CreateAccountParams{
		Userid:    uuid.String(),
		Firstname: in.FirstName,
		Lastname:  in.LastName,
		Email:     in.Email,
		Phoneno:   in.PhoneNo,
		Password:  pass,
		Roleid:    role,
	})

	// Handle any errors that occurred during the CreateAccount query
	if err != nil {
		return nil, err
	}

	if h.mailClient == nil {
		return nil, fmt.Errorf("mailClient is nil")
	}

	_, err = h.mailClient.SendMail(ctx, &proto_go.MailRequest{
		Email:   in.Email,
		Subject: "OTP for account verification",
		Body:    "<h1>OTP is 9045</h1>",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send email: %v", err)
	}

	// Return a successful SignUpResponse
	return &proto_go.SignUpResponse{
		Success: true,
		Message: "Account has been created",
	}, nil
}
