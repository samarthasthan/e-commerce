package grpc

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
	"github.com/samarthasthan/e-commerce/pkg/bcrpyt"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type AuthenticationHandler struct {
	proto_go.UnimplementedAuthenticationServiceServer
	mysql database.Database
	redis database.Database
}

func NewAuthenticationHandler(mysql database.Database, redis database.Database) *AuthenticationHandler {
	return &AuthenticationHandler{
		mysql: mysql,
		redis: redis,
	}
}

// SignUp handles the SignUp gRPC request
func (h *AuthenticationHandler) SignUp(ctx context.Context, in *proto_go.SignUpRequest) (*proto_go.SignUpResponse, error) {

	mysql, ok := h.mysql.(*database.MySQL)
	if !ok {
		return nil, fmt.Errorf("mysql is not of type *database.MySQL")
	}

	role, err := mysql.Queries.GetRole(ctx, in.Role)

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
		Role:      sql.NullString{String: role, Valid: true},
	})

	// Handle any errors that occurred during the CreateAccount query
	if err != nil {
		return nil, err
	}

	// Return a successful SignUpResponse
	return &proto_go.SignUpResponse{
		Success: true,
		Message: "Account has been created",
	}, nil
}
