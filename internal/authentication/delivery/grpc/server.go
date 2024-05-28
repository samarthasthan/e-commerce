package grpc

import (
	"fmt"
	"net"

	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
	"google.golang.org/grpc"
)

type AuthenticationGrpcServer struct {
	log    *logger.Logger
	mysql  database.Database
	redis  database.Database
	server *grpc.Server
}

func NewAuthenticationGrpcServer(log *logger.Logger, mysql, redis database.Database) *AuthenticationGrpcServer {
	server := grpc.NewServer()
	return &AuthenticationGrpcServer{log: log, mysql: mysql, redis: redis, server: server}
}

func (g *AuthenticationGrpcServer) Run(port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		g.log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	defer listener.Close()

	as := NewAuthenticationHandler(g.mysql, g.redis)
	proto_go.RegisterAuthenticationServiceServer(g.server, as)

	g.log.Infof("Authentication gRPC server listening on port %s", port)
	err = g.server.Serve(listener)
	if err != nil {
		g.log.Fatalf("Failed to serve Authentication gRPC server: %v", err)
	}
}
