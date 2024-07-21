package validation

import (
	"github.com/samarthasthan/e-commerce/pkg/proto_go"
)

func (v *Validator) SignUp(e []Error, in *proto_go.SignUpRequest) []Error {
	// Validate FirstName
	if len(in.FirstName) < 5 {
		e = append(e, Error{
			Name: "first_name",
			Msg:  "First Name should be min 5 characters long",
		})
	}

	// Validate LastName
	if len(in.LastName) < 5 {
		e = append(e, Error{
			Name: "last_name",
			Msg:  "Last Name should be min 5 characters long",
		})
	}

	e = v.Email(e, in.Email)

	e = v.PhoneNo(e, in.PhoneNo)

	e = v.Password(e, in.Password)

	// Validate RoleName
	if len(in.RoleName) < 1 {
		e = append(e, Error{
			Name: "role_name",
			Msg:  "RoleName should be min 5 characters long",
		})
	}
	return e
}

func (v *Validator) OTPVerify(e []Error, in *proto_go.VerifyEmailOTPRequest) []Error {

	e = v.Email(e, in.Email)

	e = v.OTP(e, in.Otp)

	return e
}

func (v *Validator) SignIn(e []Error, in *proto_go.SignInRequest) []Error {

	e = v.Email(e, in.Email)

	e = v.Password(e, in.Password)

	return e
}
