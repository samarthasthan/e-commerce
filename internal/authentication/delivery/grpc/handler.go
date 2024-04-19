package grpc

import (
	"context"

	"github.com/samarthasthan/e-commerce/proto_go"
)

type AuthenticationServer struct {
	proto_go.UnimplementedAuthenticationServiceServer
}

func (c *AuthenticationServer) SignUp(ctx context.Context, in *proto_go.SignUpRequest) (*proto_go.SignUpResponse, error) {
	return &proto_go.SignUpResponse{
		Success: true,
		Message: "Account has been created",
	}, nil
}
