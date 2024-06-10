package validation

import (
	"github.com/samarthasthan/e-commerce/pkg/proto_go"
)

func (v *Validator) SignUp(e []Error, in *proto_go.SignUpRequest) []Error {
	// Validate FirstName
	if len(in.FirstName) < 5 {
		e = append(e, Error{
			Name: "FirstName",
			Msg:  "FirstName should be min 5 characters long",
		})
	}

	// Validate LastName
	if len(in.LastName) < 5 {
		e = append(e, Error{
			Name: "LastName",
			Msg:  "LastName should be min 5 characters long",
		})
	}

	v.Email(e, in.Email)

	v.PhoneNo(e, in.PhoneNo)

	v.Password(e, in.Password)

	// Validate RoleName
	if len(in.RoleName) < 1 {
		e = append(e, Error{
			Name: "RoleName",
			Msg:  "RoleName should be min 5 characters long",
		})
	}

	return e
}

func (v *Validator) OTPVerify(e []Error, in *proto_go.VerifyEmailOTPRequest) []Error {

	v.Email(e, in.Email)

	v.OTP(e, in.Otp)

	return e
}

func (v *Validator) SignIn(e []Error, in *proto_go.SignInRequest) []Error {

	v.Email(e, in.Email)

	v.Password(e, in.Password)

	return e
}
