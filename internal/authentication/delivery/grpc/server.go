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
	Server grpc.Server
	log    *logger.Logger
	mysql  database.Database
	redis  database.Database
}

func NewAuthenticationGrpcServer(log *logger.Logger, mysql database.Database, redis database.Database) *AuthenticationGrpcServer {
	return &AuthenticationGrpcServer{log: log, mysql: mysql, redis: redis}
}

func (g *AuthenticationGrpcServer) Run(port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		g.log.Error(err)
		return
	}

	s := grpc.NewServer()
	as := NewAuthenticationHandler(g.mysql, g.redis)
	proto_go.RegisterAuthenticationServiceServer(s, as)

	g.log.Infof("Authentication gRPC server listening on port %s", port)
	err = s.Serve(listener)
	if err != nil {
		g.log.Error("Failed to serve gRPC server:", err)
	}
}
