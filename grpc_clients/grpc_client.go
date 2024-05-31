package grpc_clients

import (
	"time"

	"github.com/samarthasthan/e-commerce/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type BaseClient struct {
	Conn *grpc.ClientConn
	Log  *logger.Logger
}

func (b *BaseClient) Connect(addr string, l *logger.Logger) error {
	var err error
	b.Conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
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
