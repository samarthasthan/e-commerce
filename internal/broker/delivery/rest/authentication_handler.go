package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/samarthasthan/e-commerce/internal/broker/validation"
	"github.com/samarthasthan/e-commerce/pkg/proto_go"
	"github.com/sirupsen/logrus"
)

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
		s.log.WithFields(logrus.Fields{
			"email": user.Email,
		}).Errorf("Error creating user: %v", err)
		errs = append(errs, validation.Error{Name: "Service", Msg: "Something went wrong. Please try again later"})
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs)
		return
	}

	s.log.WithFields(logrus.Fields{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
	}).Infof("A new user created")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (s *RestHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUser"))
}

func (s *RestHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UpdateUser"))
}

func (s *RestHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DeleteUser"))
}

func (s *RestHandler) DisableUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DisableUser"))
}

func (s *RestHandler) OTPVerify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var otp proto_go.VerifyEmailOTPRequest
	var errs []validation.Error

	if err := json.NewDecoder(r.Body).Decode(&otp); err != nil {
		errs = append(errs, validation.Error{Name: "Input", Msg: "Invalid input"})
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	errs = s.validator.OTPVerify(errs, &otp)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	ctx := context.Background()
	res, err := s.authenticationClient.VerifyEmailOTP(ctx, &otp)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email": otp.Email,
		}).Errorf("Error verifying OTP: %v", err)
		errs = append(errs, validation.Error{Name: "Service", Msg: "Something went wrong. Please try again later"})
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *RestHandler) ResendOTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ResendOTP"))
}

func (s *RestHandler) SignInUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user proto_go.SignInRequest
	var errs []validation.Error

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errs = append(errs, validation.Error{Name: "Input", Msg: "Invalid input"})
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	errs = s.validator.SignIn(errs, &user)
	if len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errs)
		return
	}

	ctx := context.Background()
	res, err := s.authenticationClient.SignIn(ctx, &user)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"email": user.Email,
		}).Errorf("Error signing in user: %v", err)
		errs = append(errs, validation.Error{Name: "Service", Msg: "Something went wrong. Please try again later"})
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errs)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *RestHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ForgotPassword"))
}
