package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samarthasthan/e-commerce/internal/authentication/database"
	"github.com/samarthasthan/e-commerce/internal/authentication/database/mysql/sqlc"
	"github.com/samarthasthan/e-commerce/pkg/bcrpyt"
	"github.com/samarthasthan/e-commerce/pkg/kafka"
	"github.com/samarthasthan/e-commerce/pkg/models"
	"github.com/samarthasthan/e-commerce/pkg/proto_go"
	"github.com/samarthasthan/e-commerce/pkg/utils"
)

type AuthenticationHandler struct {
	proto_go.UnimplementedAuthenticationServiceServer
	kp    *kafka.Producer
	mysql database.Database
	redis database.Database
}

func NewAuthenticationHandler(mysql database.Database, redis database.Database, kp *kafka.Producer) *AuthenticationHandler {
	if mysql == nil {
		panic("mysql dependency must not be nil")
	}
	if redis == nil {
		panic("redis dependency must not be nil")
	}
	if kp == nil {
		panic("kafka dependency must not be nil")
	}
	return &AuthenticationHandler{
		mysql: mysql,
		redis: redis,
		kp:    kp,
	}
}

// SignUp handles the SignUp gRPC request
func (h *AuthenticationHandler) SignUp(ctx context.Context, in *proto_go.SignUpRequest) (*proto_go.SignUpResponse, error) {
	// Type assertion for MySQL
	mysql, ok := h.mysql.(*database.MySQL)
	if !ok {
		return nil, fmt.Errorf("mysql is not of type *database.MySQL")
	}

	tx, err := mysql.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Generate hashed password
	hashedPassword, err := bcrpyt.HashPassword(in.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	// Generate UUID
	userID := uuid.New().String()

	// Execute CreateAccount query using sqlc
	err = mysql.Queries.CreateAccount(ctx, sqlc.CreateAccountParams{
		Userid:    userID,
		Firstname: in.FirstName,
		Lastname:  in.LastName,
		Email:     in.Email,
		Phoneno:   in.PhoneNo,
		Password:  hashedPassword,
		Rolename:  in.RoleName,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create account: %v", err)
	}

	// Generate UUID
	verificationID := uuid.New().String()

	OTP := utils.GenerateVerificationCode()

	err = mysql.Queries.CreateVerification(ctx, sqlc.CreateVerificationParams{
		Verificationid: verificationID,
		Userid:         userID,
		Otp:            int32(OTP),
		Expiresat:      time.Now().Add(time.Minute * 10),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create account: %v", err)
	}

	// Check if Kafka producer is nil
	if h.kp == nil {
		return nil, fmt.Errorf("kafka producer is nil")
	}

	// Create a new Mail struct
	mail := &models.Mail{
		To:      in.Email,
		Subject: "Welcome to E-commerce",
		Body:    fmt.Sprintf("Your OTP for e-commerce is %d", OTP),
	}

	// Produce a message to the mail topic
	err = h.kp.ProduceMsg(ctx, "mail", mail)
	if err != nil {
		return nil, fmt.Errorf("failed to produce message to Kafka: %v", err)
	}

	// Return a successful SignUpResponse
	return &proto_go.SignUpResponse{
		Success:        true,
		Message:        "Account has been created",
		UserId:         userID,
		VerificationId: verificationID,
	}, nil
}

func (h *AuthenticationHandler) VerifyEmailOTP(ctx context.Context, in *proto_go.VerifyEmailOTPRequest) (*proto_go.VerifyEmailOTPResponse, error) {
	// Type assertion for MySQL
	mysql, ok := h.mysql.(*database.MySQL)
	if !ok {
		return nil, fmt.Errorf("mysql is not of type *database.MySQL")
	}

	tx, err := mysql.DB.BeginTx(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	res, err := mysql.Queries.GetOTP(ctx, in.VerificationId)

	if err != nil {
		return nil, fmt.Errorf("failed to get OTP: %v", err)
	}

	if res.Otp != in.Otp {
		return nil, fmt.Errorf("OTP does not match")
	}

	if res.Expiresat.Before(time.Now()) {
		return nil, fmt.Errorf("OTP has expired")
	}

	err = mysql.Queries.DeleteVerification(ctx, in.VerificationId)

	if err != nil {
		return nil, fmt.Errorf("failed to delete verification: %v", err)
	}

	err = mysql.Queries.VerifyAccount(ctx, res.Userid)

	if err != nil {
		return nil, fmt.Errorf("failed to verify account: %v", err)
	}

	return &proto_go.VerifyEmailOTPResponse{
		Success: true,
		Message: "Account has been verified",
	}, nil
}
