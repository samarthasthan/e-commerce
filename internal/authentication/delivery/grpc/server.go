package grpc

import (
	"fmt"
	"net"

	// Importing other dependencies
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (

	// Define the port variable
	AUTHENTICATION_GRPC_PORT = env.GetEnv("AUTHENTICATION_GRPC_PORT", "8000")
)

// AuthenticationGrpcServer represents the gRPC server for authentication
type AuthenticationGrpcServer struct {
	Server grpc.Server
	log    *logger.Logger
}

// NewAuthenticationGrpcServer initializes a new gRPC server
func NewAuthenticationGrpcServer(log *logger.Logger) *AuthenticationGrpcServer {
	return &AuthenticationGrpcServer{log: log}
}

// Run starts the gRPC server
func (g *AuthenticationGrpcServer) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", AUTHENTICATION_GRPC_PORT))

	if err != nil {
		g.log.Error(err)
	}

	s := grpc.NewServer()

	as := NewAuthenticationServer()

	proto_go.RegisterAuthenticationServiceServer(s, as)
	g.log.WithFields(
		logrus.Fields{
			"Listening on PORT: ": lis.Addr(),
		},
	).Info("Authentication gRPC server started")

	if err := s.Serve(lis); err != nil {
		g.log.WithFields(
			logrus.Fields{
				"Error: ": err,
			},
		).Error("Authentication gRPC server failed to serve")
	}
}
