package grpc

// Import necessary packages
import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
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

	uuid := uuid.New()
	// Execute CreateAccount query using sqlc
	_, err := c.db.Queries.CreateAccount(ctx, sqlc.CreateAccountParams{
		Userid:    uuid.String(),
		Firstname: in.FirstName,
		Lastname:  in.LastName,
		Email:     in.Email,
		Phoneno:   sql.NullString{String: in.PhoneNo, Valid: true},
		Password:  in.Password,
		Role:      sql.NullString{String: "83064af3-bb81-4514-a6d4-afba340825cd", Valid: true},
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
