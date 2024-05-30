package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type RestHandler struct {
	mux                  *http.ServeMux
	authenticationClient proto_go.AuthenticationServiceClient
	log                  *logger.Logger
}

func NewRestHandler(ac proto_go.AuthenticationServiceClient, l *logger.Logger) *RestHandler {
	return &RestHandler{
		mux:                  http.NewServeMux(),
		authenticationClient: ac,
		log:                  l,
	}
}

func (s *RestHandler) Handle() {
	s.mux.HandleFunc("/", s.CreateUser)
}

func (s *RestHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	res, err := s.authenticationClient.SignUp(ctx, &proto_go.SignUpRequest{
		FirstName: "Samarth",
		LastName:  "Asthan",
		Email:     "samarthasthan27@gmail.com",
		PhoneNo:   "91 9557030000",
		Password:  "password",
		RoleName:  "user",
	})

	if err != nil {
		fmt.Fprintf(w, "error: %v", err.Error())
		return
	}

	fmt.Fprintf(w, "%v", res)
}

func (s *RestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
