package bcrpyt

import "testing"

func TestHashPassword(t *testing.T) {
	password := "password"
	hashed, err := HashPassword(password)
	if err != nil {
		t.Errorf("HashPassword() error = %v", err)
	}
	if ValidatePassword(hashed, password) == false {
		t.Errorf("ValidatePassword() = false")
	}
}
