package utils

import (
	"math/rand"
	"time"
)

func GenerateVerificationCode() int {
	return rand.Intn(999999-100000) + 100000
}

func CheckOTPExpiration(t time.Time) bool {
	return time.Now().Before(t)
}
