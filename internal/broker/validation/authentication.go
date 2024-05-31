package validation

import (
	"regexp"

	"github.com/samarthasthan/e-commerce/proto_go"
)

func (v *Validator) SignUp(e []Error, u *proto_go.SignUpRequest) []Error {
	// Validate FirstName
	if len(u.FirstName) < 5 {
		e = append(e, Error{
			Name: "FirstName",
			Msg:  "FirstName should be min 5 characters long",
		})
	}

	// Validate LastName
	if len(u.LastName) < 5 {
		e = append(e, Error{
			Name: "LastName",
			Msg:  "LastName should be min 5 characters long",
		})
	}

	// Validate Email
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		e = append(e, Error{
			Name: "Email",
			Msg:  "Invalid email format",
		})
	}

	// Validate PhoneNo
	phoneRegex := regexp.MustCompile(`^[0-9]{10}$`)
	if !phoneRegex.MatchString(u.PhoneNo) {
		e = append(e, Error{
			Name: "PhoneNo",
			Msg:  "Invalid phone number format",
		})
	}

	// Validate Password
	passwordRegex := regexp.MustCompile(`^[A-Za-z0-9!@#$%^&*]{8,}$`)
	if !passwordRegex.MatchString(u.Password) {
		e = append(e, Error{
			Name: "Password",
			Msg:  "Password should have at least 8 characters long",
		})
	}

	// Validate RoleName
	if len(u.RoleName) < 1 {
		e = append(e, Error{
			Name: "RoleName",
			Msg:  "RoleName should be min 5 characters long",
		})
	}

	return e
}
