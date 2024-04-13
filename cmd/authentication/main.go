package main

import "github.com/samarthasthan/e-commerce/internal/authentication"

func main() {
	// Authentication gRPC server
	as := authentication.NewAuthServer()
	as.Connect()
}
