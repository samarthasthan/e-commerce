package rest

import (
	"encoding/json"
	"net/http"

	"github.com/samarthasthan/e-commerce/internal/broker/delivery/grpc"
	"github.com/samarthasthan/e-commerce/pkg/logger"
	"github.com/samarthasthan/e-commerce/proto_go"
)

type Handler struct {
	log        *logger.Logger
	grpcClient *grpc.GRPCClients
}

func NewHandler(log *logger.Logger, grpcClient *grpc.GRPCClients) *Handler {
	return &Handler{log: log, grpcClient: grpcClient}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user proto_go.SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	h.log.Info(&user)

	validationErrors := make(map[string]string)

	if len(user.FirstName) < 5 {
		validationErrors["FirstName"] = "First name should be at least 5 characters"
	}
	if len(user.LastName) < 5 {
		validationErrors["LastName"] = "Last name should be at least 5 characters"
	}
	if len(user.Email) == 0 {
		validationErrors["Email"] = "Email is required"
	}
	if len(user.PhoneNo) < 10 {
		validationErrors["PhoneNo"] = "Phone number should be at least 10 characters"
	}
	if len(user.Password) < 8 {
		validationErrors["Password"] = "Password should be at least 8 characters"
	}

	if len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": validationErrors,
		})
		return
	}

	json.NewEncoder(w).Encode(&user)
}
