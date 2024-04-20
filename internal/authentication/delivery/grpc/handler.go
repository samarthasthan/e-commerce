package grpc

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type AuthenticationServer struct {
	proto_go.UnimplementedAuthenticationServiceServer
	queries *sqlc.Queries
}

func NewAuthenticationServer() *AuthenticationServer {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:8001)/authentication") //mysql://root:password@vscode.samarthasthan.com:8001/authentication?statusColor=F8F8F8&env=local&name=Authentication&tLSMode=0&usePrivateKey=false&safeModeLevel=0&advancedSafeModeLevel=0&driverVersion=0&lazyload=true
	if err != nil {
		panic(err)
	}

	queries := sqlc.New(db)

	return &AuthenticationServer{
		queries: queries,
	}
}

func (c *AuthenticationServer) SignUp(ctx context.Context, in *proto_go.SignUpRequest) (*proto_go.SignUpResponse, error) {
	c.queries.CreateAccount(ctx, sqlc.CreateAccountParams{
		Userid:    "53264af3-bb81-4514-a6d4-afba340825cd",
		Firstname: in.FirstName,
		Lastname:  in.LastName,
		Email:     in.Email,
		Phoneno:   sql.NullString{String: in.PhoneNo, Valid: true},
		Password:  in.Password,
		Role:      sql.NullString{String: "83064af3-bb81-4514-a6d4-afba340825cd", Valid: true},
	})
	return &proto_go.SignUpResponse{
		Success: true,
		Message: "Account has been created",
	}, nil
}
