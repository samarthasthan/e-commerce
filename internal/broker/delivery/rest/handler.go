package rest

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/samarthasthan/e-commerce/internal/broker/validation"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type RestHandler struct {
	mux                  *http.ServeMux
	authenticationClient proto_go.AuthenticationServiceClient
	validator            *validation.Validator
	log                  *logger.Logger
}

func NewRestHandler(ac proto_go.AuthenticationServiceClient, v *validation.Validator, l *logger.Logger) *RestHandler {
	return &RestHandler{
		mux:                  http.NewServeMux(),
		authenticationClient: ac,
		validator:            v,
		log:                  l,
	}
}

func (s *RestHandler) Handle() {
	s.mux.Handle("/metrics", promhttp.Handler())
	s.mux.HandleFunc("POST /create-user", s.CreateUser)
}

func (s *RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
