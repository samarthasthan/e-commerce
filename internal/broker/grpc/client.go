package grpc

import (
	"time"

	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient interface {
	Connect(string) error
	Close()
}

type AuthenticationClient struct {
	Client proto_go.AuthenticationServiceClient
	Conn   *grpc.ClientConn
	Log    *logger.Logger
}

func NewAuthenticationClient() *AuthenticationClient {
	return &AuthenticationClient{}
}

func (a *AuthenticationClient) Connect(addr string) error {
	var err error
	a.Conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		a.Log.Error("Failed to connect to Authentication Service: ", err)
		return err
	}
	a.Client = proto_go.NewAuthenticationServiceClient(a.Conn)
	a.Log.Info("Successfully connected to Authentication Service")
	return nil
}

func (a *AuthenticationClient) Close() error {
	err := a.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
