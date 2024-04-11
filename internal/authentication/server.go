package authentication

import (
	"log"
	"net"

	proto_go "github.com/samarthasthan/e-commerce/proto_go"
	"google.golang.org/grpc"
)

const (
	PORT = ":7000"
)

type AuthServer struct {
	proto_go.UnimplementedAuthenticationServiceServer
}

func NewAuthServer() AuthServer {
	return AuthServer{}
}

func (a *AuthServer) Connect() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto_go.RegisterAuthenticationServiceServer(s, &AuthServer{})

	log.Printf("Authentication gRPC listing at %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
