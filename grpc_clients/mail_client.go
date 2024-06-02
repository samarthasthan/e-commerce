package grpc_clients

import (
	"fmt"

	"github.com/openzipkin/zipkin-go"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type MailClient struct {
	baseClient
	Client proto_go.MailServiceClient
}

func NewMailClient(log *logger.Logger, t *zipkin.Tracer) MailClient {
	return MailClient{
		baseClient: baseClient{
			log:    log,
			tracer: t,
		},
	}
}

func (a *MailClient) GetClient() *proto_go.MailServiceClient {
	return &a.Client
}

func (a *MailClient) Connect(addr string) error {
	err := a.baseClient.Connect(fmt.Sprintf("localhost:%v", addr))
	if err != nil {
		return err
	}
	a.Client = proto_go.NewMailServiceClient(a.conn)
	a.log.Info("Successfully connected to Mail Service")
	return nil
}
