package utils

import (
	"testing"
	"time"
)

func TestGenerateVerificationCode(t *testing.T) {
	code := GenerateVerificationCode()
	if code < 100000 || code > 999999 {
		t.Errorf("GenerateVerificationCode() = %v; want a value between 100000 and 999999", code)
	}
}

func TestCheckOTPExpiration(t *testing.T) {
	tests := []struct {
		dbTime  time.Time
		expired bool
	}{
		{time.Now().Add(-1 * time.Minute), false},
		{time.Now().Add(-6 * time.Minute), false},
		{time.Now().Add(-5 * time.Minute), false},
		{time.Now().Add(1 * time.Minute), true},
		{time.Now().Add(6 * time.Minute), true},
		{time.Now().Add(5 * time.Minute), true},
	}

	for _, tt := range tests {
		if got := CheckOTPExpiration(tt.dbTime); got != tt.expired {
			t.Errorf("CheckOTPExpiration(%v) = %v, want %v", tt.dbTime, got, tt.expired)
		}
	}
}
