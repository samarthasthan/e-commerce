package validation

import (
	"github.com/samarthasthan/e-commerce/proto_go"
)

func (v *Validator) SignUp(e []Error, u *proto_go.SignUpRequest) []Error {
	if len(u.FirstName) < 5 {
		e = append(e, Error{
			Name: "FirstName",
			Msg:  "FirstName should be min 5 long",
		})
	}
	if len(u.LastName) < 5 {
		e = append(e, Error{
			Name: "LastName",
			Msg:  "LastName should be min 5 long",
		})
	}
	return e
}
