package grpc_clients

import (
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type MailClient struct {
	BaseClient
	Client proto_go.MailServiceClient
	log    *logger.Logger
}

func NewMailClient(log *logger.Logger) *MailClient {
	return &MailClient{
		BaseClient: BaseClient{
			Log: log,
		},
	}
}

func (m *MailClient) Connect(addr string) error {
	err := m.BaseClient.Connect(addr, m.log)
	if err != nil {
		return err
	}
	m.Client = proto_go.NewMailServiceClient(m.Conn)
	m.Log.Info("Successfully connected to Mail Service")
	return nil
}
