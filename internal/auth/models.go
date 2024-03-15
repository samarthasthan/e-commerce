package auth

import "net/url"

type User struct {
	UUID            string `json:"uuid"`
	Email           string `json:"email"`
	Phone           uint64 `json:"phone"`
	Hashed_password string `json:"hashed_password"`
	Country         string `json:"country"`
	Is_verified     bool   `json:"is_verified"`
}

func (u *User) Validate() url.Values {
	errs := url.Values{}
	if u.Email == "" {
		errs.Add("email", "Email must be a valid email")
	}

	if u.Phone == 0 {
		errs.Add("phone", "Phone must be a valid phone number")
	}

	if u.Hashed_password == "" {
		errs.Add("password", "Password must must contain a Upper case, One special charecter and must be minimum 8 chars")
	}

	if u.Country == "" {
		errs.Add("country", "Invalid country")
	}

	return errs
}
