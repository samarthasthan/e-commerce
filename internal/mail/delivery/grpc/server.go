package grpc

import (
	"fmt"
	"net"

	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
	"google.golang.org/grpc"
)

type MailGrpcServer struct {
	log    *logger.Logger
	server *grpc.Server
}

func NewMailGrpcServer(l *logger.Logger) *MailGrpcServer {
	server := grpc.NewServer()
	return &MailGrpcServer{log: l, server: server}
}

func (g *MailGrpcServer) Run(port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		g.log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	defer listener.Close()

	ms := NewMailHandler()
	proto_go.RegisterMailServiceServer(g.server, ms)

	g.log.Infof("Mail gRPC server listening on port %s", port)
	err = g.server.Serve(listener)
	if err != nil {
		g.log.Fatalf("Failed to serve Mail gRPC server: %v", err)
	}
}
