package grpc_clients

import (
	"fmt"

	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type AuthenticationClient struct {
	BaseClient
	Client proto_go.AuthenticationServiceClient
	log    *logger.Logger
}

func NewAuthenticationClient(log *logger.Logger) AuthenticationClient {
	return AuthenticationClient{
		BaseClient: BaseClient{
			Log: log,
		},
	}
}

func (a *AuthenticationClient) GetClient() *proto_go.AuthenticationServiceClient {
	return &a.Client
}

func (a *AuthenticationClient) Connect(addr string) error {
	err := a.BaseClient.Connect(fmt.Sprintf("localhost:%v", addr), a.log)
	if err != nil {
		return err
	}
	a.Client = proto_go.NewAuthenticationServiceClient(a.Conn)
	a.Log.Info("Successfully connected to Authentication Service")
	return nil
}
