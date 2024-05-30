package grpc_clients

import (
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"google.golang.org/grpc"
)

type BaseClient struct {
	Conn *grpc.ClientConn
	Log  *logger.Logger
}

func (b *BaseClient) Connect(addr string, l *logger.Logger) error {
	var err error
	b.Conn, err = grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseClient) Close() error {
	if b.Conn != nil {
		err := b.Conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
