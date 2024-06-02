package grpc

import (
	"fmt"
	"net"

	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
	"google.golang.org/grpc"
)

type MailGrpcServer struct {
	log    *logger.Logger
	server *grpc.Server
}

func NewMailGrpcServer(
	l *logger.Logger,
	t *zipkin.Tracer,
) *MailGrpcServer {
	server := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(t)))
	return &MailGrpcServer{log: l, server: server}
}

func (g *MailGrpcServer) Run(
	port string,
	SMTP_SERVER string,
	SMTP_PORT string,
	SMTP_LOGIN string,
	SMTP_PASSWORD string) {
	// Creating a listener on the specified port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		g.log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	defer listener.Close()

	// Registering the Mail service
	ms := NewMailHandler(SMTP_SERVER, SMTP_PORT, SMTP_LOGIN, SMTP_PASSWORD)

	// Registering the Mail service with the gRPC server
	proto_go.RegisterMailServiceServer(g.server, ms)

	g.log.Infof("Mail gRPC server listening on port %s", port)
	err = g.server.Serve(listener)
	if err != nil {
		g.log.Fatalf("Failed to serve Mail gRPC server: %v", err)
	}
}
