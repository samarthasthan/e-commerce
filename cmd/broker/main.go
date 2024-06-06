package main

import (
	"fmt"
	"net/http"

	"github.com/samarthasthan/e-commerce/internal/broker/delivery/rest"
	"github.com/samarthasthan/e-commerce/internal/broker/validation"
	"github.com/samarthasthan/e-commerce/pkg/env"
	"github.com/samarthasthan/e-commerce/pkg/grpc_clients"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	tracer "github.com/samarthasthan/e-commerce/pkg/zipkin"
)

var (
	BROKER_REST_PORT         string
	AUTHENTICATION_GRPC_HOST string
	AUTHENTICATION_GRPC_PORT string
)

func init() {
	BROKER_REST_PORT = env.GetEnv("BROKER_REST_PORT", "7000")
	AUTHENTICATION_GRPC_PORT = env.GetEnv("AUTHENTICATION_GRPC_PORT", "8000")
	AUTHENTICATION_GRPC_HOST = env.GetEnv("AUTHENTICATION_GRPC_HOST", "localhost")
}

func main() {
	log := logger.NewLogger("Broker")

	// create a new Zipkin tracer
	tracer, reporter, err := tracer.NewTracer("broker", 7000)
	if err != nil {
		log.Fatalf("failed to create tracer: %v", err)
	}
	defer reporter.Close()

	mux := http.NewServeMux()

	validator := validation.NewValidator()

	authentitcationClient := grpc_clients.NewAuthenticationClient(log, tracer)
	if ac_err := authentitcationClient.Connect(AUTHENTICATION_GRPC_HOST, AUTHENTICATION_GRPC_PORT); ac_err != nil {
		log.Errorf("Broker not able to connect to Authentication service, msg %v", ac_err.Error())

	}
	defer authentitcationClient.Close()

	handler := rest.NewRestHandler(authentitcationClient.Client, validator, log, mux, tracer)
	handler.Handle()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", BROKER_REST_PORT),
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Errorf("Broker could not listen on %s due to %s", server.Addr, err)
		panic(err)
	}

	defer server.Close()
}
