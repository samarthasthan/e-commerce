package grpc

import (
	"fmt"
	"time"

	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	AUTHENTICATION_GRPC_PORT = env.GetEnv("AUTHENTICATION_GRPC_PORT", "8000")
	AUTHENTICATION_GRPC_HOST = env.GetEnv("AUTHENTICATION_GRPC_HOST", "localhost")
	ACCOUNT_GRPC_PORT        = env.GetEnv("ACCOUNT_GRPC_PORT", "8001")
	ACCOUNT_GRPC_HOST        = env.GetEnv("ACCOUNT_GRPC_HOST", "localhost")
	// Add more ports and hosts for other services as needed
)

// GRPCClients manages gRPC client connections.
type GRPCClients struct {
	AuthenticationClient proto_go.AuthenticationServiceClient
	// AccountClient        proto_go.AccountServiceClient
	// Add more client fields for other services

	log   *logger.Logger
	conns map[string]*grpc.ClientConn // Map of connections for different services
}

// NewGRPCClients creates a new instance of GRPCClients with a given logger.
func NewGRPCClients(log *logger.Logger) *GRPCClients {
	return &GRPCClients{
		log:   log,
		conns: make(map[string]*grpc.ClientConn),
	}
}

// ConnectToAuthenticationService connects the client to the Authentication Service and assigns the client.
func (c *GRPCClients) ConnectToAuthenticationService() error {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", AUTHENTICATION_GRPC_HOST, AUTHENTICATION_GRPC_PORT),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
	)
	if err != nil {
		c.log.Error("Failed to connect to Authentication Service: ", err)
		return err
	}
	defer conn.Close()

	c.AuthenticationClient = proto_go.NewAuthenticationServiceClient(conn)
	c.conns["authentication"] = conn
	c.log.Info("Successfully connected to Authentication Service")
	return nil
}

// // ConnectToAccountService connects the client to the Account Service and assigns the client.
// func (c *GRPCClients) ConnectToAccountService() error {
// 	conn, err := grpc.Dial(
// 		fmt.Sprintf("%s:%s", ACCOUNT_GRPC_HOST, ACCOUNT_GRPC_PORT),
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 		grpc.WithBlock(),
// 		grpc.WithTimeout(5*time.Second),
// 	)
// 	if err != nil {
// 		c.log.Error("Failed to connect to Account Service: ", err)
// 		return err
// 	}

// 	c.AccountClient = proto_go.NewAccountServiceClient(conn)
// 	c.conns["account"] = conn
// 	c.log.Info("Successfully connected to Account Service")
// 	return nil
// }

// Close gracefully closes all connections to the gRPC services.
func (c *GRPCClients) Close() error {
	for _, conn := range c.conns {
		if err := conn.Close(); err != nil {
			c.log.Error("Failed to close connection: ", err)
			return err
		}
	}
	return nil
}
