package grpc

// Import necessary packages
import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
	bcrpyt "github.com/samarthasthan/e-commerce/pkg/bcrpyt"
	"github.com/samarthasthan/e-commerce/proto_go"
)

// AuthenticationHandler implements the proto_go.AuthenticationServiceServer interface
type AuthenticationHandler struct {
	proto_go.UnimplementedAuthenticationServiceServer
	db *database.Database
}

// NewAuthenticationHandler initializes a new AuthenticationHandler
func NewAuthenticationHandler(db *database.Database) *AuthenticationHandler {
	return &AuthenticationHandler{
		db: db,
	}
}

// SignUp handles the SignUp gRPC request
func (c *AuthenticationHandler) SignUp(ctx context.Context, in *proto_go.SignUpRequest) (*proto_go.SignUpResponse, error) {

	role, err := c.db.Queries.GetRole(ctx, in.Role)

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
	_, err = c.db.Queries.CreateAccount(ctx, sqlc.CreateAccountParams{
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
