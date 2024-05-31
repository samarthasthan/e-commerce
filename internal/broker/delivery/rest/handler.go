package rest

import (
	"context"
	"encoding/json"
	"net/http"

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
	s.mux.HandleFunc("POST /create-user", s.CreateUser)
}

func (s *RestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user proto_go.SignUpRequest
	var errs []validation.Error

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errs = append(errs, validation.Error{Name: "Input", Msg: "Invalid input"})
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	errs = s.validator.SignUp(errs, &user)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	ctx := context.Background()
	res, err := s.authenticationClient.SignUp(ctx, &user)
	if err != nil {
		errs = append(errs, validation.Error{Name: "Service", Msg: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs)
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (s *RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
