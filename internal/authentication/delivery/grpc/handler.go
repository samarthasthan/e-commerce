package grpc

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
	"github.com/samarthasthan/e-commerce/pkg/bcrpyt"
	"github.com/samarthasthan/e-commerce/pkg/kafka"
	"github.com/samarthasthan/e-commerce/pkg/models"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type AuthenticationHandler struct {
	proto_go.UnimplementedAuthenticationServiceServer
	kp    *kafka.Producer
	mysql database.Database
	redis database.Database
}

func NewAuthenticationHandler(mysql database.Database, redis database.Database, kp *kafka.Producer) *AuthenticationHandler {
	if mysql == nil {
		panic("mysql dependency must not be nil")
	}
	if redis == nil {
		panic("redis dependency must not be nil")
	}
	if kp == nil {
		panic("kafka dependency must not be nil")
	}
	return &AuthenticationHandler{
		mysql: mysql,
		redis: redis,
		kp:    kp,
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

	if h.kp == nil {
		return nil, fmt.Errorf("kafka is nil")
	}

	// Create a new Mail struct
	mail := &models.Mail{To: in.Email, Subject: "Welcome to E-commerce", Body: "<h1>Your account has been created successfully</h1>"}

	// Produce a message to the mail topic
	h.kp.ProduceMsg([]string{"mail"}, mail)
	// Handle any errors that occurred during the ProduceMsg function
	if err != nil {
		return nil, err
	}

	// Return a successful SignUpResponse
	return &proto_go.SignUpResponse{
		Success: true,
		Message: "Account has been created",
	}, nil
}
