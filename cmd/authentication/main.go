package main

import grpc "github.com/samarthasthan/e-commerce/internal/authentication/delivery/grpc"

func main() {
	s := grpc.NewAuthenticationGrpcServer()
	s.Run()
}
