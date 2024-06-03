package utils

import (
	"math/rand"
)

func GenerateVerificationCode() int {
	return rand.Intn(999999-100000) + 100000
}
