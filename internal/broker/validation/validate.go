package validation

import (
	"regexp"
	"strconv"
)

type Validator struct {
}

type Error struct {
	Name string
	Msg  string
}

type Errors struct {
	_ []Error
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) Password(e []Error, ps string) []Error {
	// Check password length
	if len(ps) < 8 || len(ps) > 16 {
		e = append(e, Error{
			Name: "password",
			Msg:  "Password should be 8-16 characters long",
		})
		return e
	}

	// Check for at least one digit
	if !regexp.MustCompile(`[0-9]`).MatchString(ps) {
		e = append(e, Error{
			Name: "password",
			Msg:  "Password should contain at least one digit",
		})
		return e
	}

	// Check for at least one lowercase letter
	if !regexp.MustCompile(`[a-z]`).MatchString(ps) {
		e = append(e, Error{
			Name: "password",
			Msg:  "Password should contain at least one lowercase letter",
		})
		return e
	}

	// Check for at least one uppercase letter
	if !regexp.MustCompile(`[A-Z]`).MatchString(ps) {
		e = append(e, Error{
			Name: "password",
			Msg:  "Password should contain at least one uppercase letter",
		})
		return e
	}

	// Check for at least one special character
	if !regexp.MustCompile(`[\W_]`).MatchString(ps) {
		e = append(e, Error{
			Name: "password",
			Msg:  "Password should contain at least one special character",
		})
		return e
	}

	return e
}

func (v *Validator) PhoneNo(e []Error, ph string) []Error {
	// Validate Phone Number
	phoneRegex := regexp.MustCompile(`^[0-9]{10}$`)
	if !phoneRegex.MatchString(ph) {
		e = append(e, Error{
			Name: "phone_no",
			Msg:  "PhoneNo should be 10 digits long",
		})
	}
	return e
}

func (v *Validator) Email(e []Error, em string) []Error {
	// Validate Email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(em) {
		e = append(e, Error{
			Name: "email",
			Msg:  "Invalid email format",
		})
	}
	return e
}

func (v *Validator) OTP(e []Error, otp int32) []Error {
	// Validate OTP
	otpRegex := regexp.MustCompile(`^[0-9]{6,6}$`)
	if !otpRegex.MatchString(strconv.Itoa(int(otp))) {
		e = append(e, Error{
			Name: "otp",
			Msg:  "Invalid OTP format",
		})
	}
	return e
}
