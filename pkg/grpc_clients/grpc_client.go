package grpc_clients

import (
	"time"

	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type baseClient struct {
	conn   *grpc.ClientConn
	log    *logger.Logger
	tracer *zipkin.Tracer
}

func (b *baseClient) Connect(addr string) error {
	var err error
	b.conn, err = grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second),
		grpc.WithStatsHandler(zipkingrpc.NewClientHandler(b.tracer)))
	if err != nil {
		return err
	}
	return nil
}

func (b *baseClient) Close() error {
	if b.conn != nil {
		err := b.conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
