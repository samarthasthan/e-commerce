package grpc_clients

import (
	"fmt"

	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type MailClient struct {
	BaseClient
	Client proto_go.MailServiceClient
	log    *logger.Logger
}

func NewMailClient(log *logger.Logger) MailClient {
	return MailClient{
		BaseClient: BaseClient{
			Log: log,
		},
	}
}

func (a *MailClient) GetClient() *proto_go.MailServiceClient {
	return &a.Client
}

func (a *MailClient) Connect(addr string) error {
	err := a.BaseClient.Connect(fmt.Sprintf("localhost:%v", addr), a.log)
	if err != nil {
		return err
	}
	a.Client = proto_go.NewMailServiceClient(a.Conn)
	a.Log.Info("Successfully connected to Mail Service")
	return nil
}
