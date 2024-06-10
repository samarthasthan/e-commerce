package utils

import "testing"

func TestGenerateVerificationCode(t *testing.T) {
	code := GenerateVerificationCode()
	if code < 100000 || code > 999999 {
		t.Errorf("GenerateVerificationCode() = %v; want a value between 100000 and 999999", code)
	}
}
