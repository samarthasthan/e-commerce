package rest

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/samarthasthan/e-commerce/internal/broker/validation"
	"github.com/samarthasthan/e-commerce/proto_go"
	"github.com/sirupsen/logrus"
)

// randomDelay introduces a random delay between 0 and maxDelay milliseconds
func randomDelay(maxDelay int) {
	delay := time.Duration(rand.Intn(maxDelay)) * time.Millisecond
	time.Sleep(delay)
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
	randomDelay(1000)

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

	s.log.WithFields(logrus.Fields{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	}).Infof("A new user created")
	json.NewEncoder(w).Encode(res)
}

func (s *RestHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	randomDelay(1000)
	w.Write([]byte("GetUser"))
}

func (s *RestHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func (s *RestHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	randomDelay(1000)
	w.Write([]byte("DeleteUser"))
}

func (s *RestHandler) DisableUser(w http.ResponseWriter, r *http.Request) {
	randomDelay(1000)
	w.Write([]byte("DisableUser"))
}

func (s *RestHandler) EmailVerify(w http.ResponseWriter, r *http.Request) {
	randomDelay(1000)
	w.Write([]byte("EmailVerify"))
}

func (s *RestHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	randomDelay(1000)
	w.Write([]byte("LoginUser"))
}

func (s *RestHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	randomDelay(1000)
	w.Write([]byte("ForgotPassword"))
}
