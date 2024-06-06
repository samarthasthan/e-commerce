package validation

import (
	"regexp"
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
	// Validate Password
	passwordRegex := regexp.MustCompile(`^[A-Za-z0-9!@#$%^&*]{8,}$`)
	if !passwordRegex.MatchString(ps) {
		e = append(e, Error{
			Name: "Password",
			Msg:  "Password should have at least 8 characters long",
		})
	}
	return e
}

func (v *Validator) PhoneNo(e []Error, ph string) []Error {
	// Validate Phone Number
	phoneRegex := regexp.MustCompile(`^[0-9]{10}$`)
	if !phoneRegex.MatchString(ph) {
		e = append(e, Error{
			Name: "PhoneNo",
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
			Name: "Email",
			Msg:  "Invalid email format",
		})
	}
	return e
}

func (v *Validator) OTP(e []Error, otp int32) []Error {
	// Validate OTP
	otpRegex := regexp.MustCompile(`^[0-9]{6,6}$`)
	if !otpRegex.Match([]byte{byte(otp)}) {
		e = append(e, Error{
			Name: "OTP",
			Msg:  "Invalid OTP format",
		})
	}
	return e
}
