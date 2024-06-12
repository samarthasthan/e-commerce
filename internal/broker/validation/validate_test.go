package validation

import (
	"testing"
)

var validation *Validator

func init() {
	validation = NewValidator()
}
func TestPassword(t *testing.T) {
	var e []Error
	e = validation.Password(e, "MyP@ssw0rd")
	if len(e) != 0 {
		t.Errorf("Expected 0 error, got %v", len(e))
	}

	clear(e)
	e = validation.Password(e, "password")
	if len(e) == 0 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}
}

func TestPhone(t *testing.T) {
	var e []Error
	e = validation.PhoneNo(e, "9411950169")
	if len(e) != 0 {
		t.Errorf("Expected 0 error, got %v", len(e))
	}

	clear(e)
	e = validation.PhoneNo(e, "0812345678901")
	if len(e) == 0 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}

	clear(e)
	e = validation.PhoneNo(e, "941195")
	if len(e) == 0 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}
}

func TestEmail(t *testing.T) {
	var e []Error
	e = validation.Email(e, "samarthasthan9411@gmail.com")
	if len(e) != 0 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}

	clear(e)
	e = validation.Email(e, "samarthasthan5gmail.com")
	if len(e) == 0 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}

	clear(e)
	e = validation.Email(e, "9411gmail.i")
	if len(e) == 0 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}
}

func TestOTP(t *testing.T) {
	var e []Error
	e = validation.OTP(e, 903466)
	if len(e) != 0 {
		t.Errorf("Expected 0 error, got %v", len(e))
	}

	clear(e)
	e = validation.OTP(e, 2345)
	if len(e) == 0 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}
}
