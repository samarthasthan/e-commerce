package grpc

import (
	"fmt"
	"net"

	"github.com/openzipkin/zipkin-go"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/pkg/kafka"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/pkg/proto_go"
	"google.golang.org/grpc"
)

type AuthenticationGrpcServer struct {
	log    *logger.Logger
	mysql  database.Database
	redis  database.Database
	server *grpc.Server
	kp     *kafka.Producer
}

func NewAuthenticationGrpcServer(log *logger.Logger, mysql, redis database.Database, kp *kafka.Producer, t *zipkin.Tracer) *AuthenticationGrpcServer {
	// Creating a new gRPC server
	server := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(t)))
	return &AuthenticationGrpcServer{log: log, mysql: mysql, redis: redis, server: server, kp: kp}
}

func (g *AuthenticationGrpcServer) Run(port string) {
	// Creating a listener on the specified port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		g.log.Fatalf("Failed to listen on port %s: %v", port, err)
	}
	defer listener.Close()

	// Registering the Authentication service
	as := NewAuthenticationHandler(g.mysql, g.redis, g.kp)

	// Registering the Authentication service with the gRPC server
	proto_go.RegisterAuthenticationServiceServer(g.server, as)

	g.log.Infof("Authentication gRPC server listening on port %s", port)
	err = g.server.Serve(listener)
	if err != nil {
		g.log.Fatalf("Failed to serve Authentication gRPC server: %v", err)
	}

	defer g.server.GracefulStop()
	g.log.Info("Authentication gRPC server stopped")
}
