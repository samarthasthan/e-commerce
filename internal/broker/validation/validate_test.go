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
	e = validation.Password(e, "1234")
	if len(e) != 0 {
		t.Errorf("Expected 0 error, got %v", len(e))
	}
	e = validation.Password(e, "password")
	if len(e) != 1 {
		t.Errorf("Expected 1 error, got %v", len(e))
	}
}
