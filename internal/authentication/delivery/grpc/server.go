package grpc

import (
	"fmt"
	"net"

	// Importing the logger from the authentication package
	"github.com/samarthasthan/e-commerce/internal/authentication"

	// Importing other dependencies
	"github.com/samarthasthan/e-commerce/pkg/env"
	proto_go "github.com/samarthasthan/e-commerce/proto_go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	// Define a shorter name for the logger for easier access
	log = authentication.Logrusc

	// Define the port variable
	AUTHENTICATION_GRPC_PORT = env.GetEnv("AUTHENTICATION_GRPC_PORT", "8000")
)

// AuthenticationGrpcServer represents the gRPC server for authentication
type AuthenticationGrpcServer struct {
	Server grpc.Server
}

// NewAuthenticationGrpcServer initializes a new gRPC server
func NewAuthenticationGrpcServer() *AuthenticationGrpcServer {
	return &AuthenticationGrpcServer{}
}

// Run starts the gRPC server
func (g *AuthenticationGrpcServer) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", AUTHENTICATION_GRPC_PORT))

	if err != nil {
		log.Error(err)
	}

	s := grpc.NewServer()

	proto_go.RegisterAuthenticationServiceServer(s, &AuthenticationServer{})
	log.WithFields(
		logrus.Fields{
			"Listening on PORT: ": lis.Addr(),
		},
	).Info("Authentication gRPC server started")

	if err := s.Serve(lis); err != nil {
		log.WithFields(
			logrus.Fields{
				"Error: ": err,
			},
		).Error("Authentication gRPC server failed to serve")
	}
}
