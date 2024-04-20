package grpc

import (
	"fmt"
	"net"

	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
	"google.golang.org/grpc"
)

var (
	AUTHENTICATION_GRPC_PORT = env.GetEnv("AUTHENTICATION_GRPC_PORT", "8000")
)

type AuthenticationGrpcServer struct {
	Server grpc.Server
	log    *logger.Logger
	db     *database.Database
}

func NewAuthenticationGrpcServer(log *logger.Logger, db *database.Database) *AuthenticationGrpcServer {
	return &AuthenticationGrpcServer{log: log, db: db}
}

func (g *AuthenticationGrpcServer) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", AUTHENTICATION_GRPC_PORT))
	if err != nil {
		g.log.Error(err)
		return
	}

	s := grpc.NewServer()
	as := NewAuthenticationHandler(g.db)
	proto_go.RegisterAuthenticationServiceServer(s, as)

	g.log.Infof("Authentication gRPC server listening on port %s", AUTHENTICATION_GRPC_PORT)
	err = s.Serve(listener)
	if err != nil {
		g.log.Error("Failed to serve gRPC server:", err)
	}
}
