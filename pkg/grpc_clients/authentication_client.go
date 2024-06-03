package grpc_clients

import (
	"fmt"

	"github.com/openzipkin/zipkin-go"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type AuthenticationClient struct {
	baseClient
	Client proto_go.AuthenticationServiceClient
}

func NewAuthenticationClient(log *logger.Logger, t *zipkin.Tracer) AuthenticationClient {
	return AuthenticationClient{
		baseClient: baseClient{
			log:    log,
			tracer: t,
		},
	}
}

func (a *AuthenticationClient) GetClient() *proto_go.AuthenticationServiceClient {
	return &a.Client
}

func (a *AuthenticationClient) Connect(addr string) error {
	err := a.baseClient.Connect(fmt.Sprintf("localhost:%v", addr))
	if err != nil {
		return err
	}
	a.Client = proto_go.NewAuthenticationServiceClient(a.conn)
	a.log.Info("Successfully connected to Authentication Service")
	return nil
}
