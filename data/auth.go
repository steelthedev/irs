package data

import "github.com/steelthedev/irs/utils"

type RegisterUser struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterUserValidator interface {
	Validate() error
}

func (h RegisterUser) Validate() error {
	if len(h.Email) < 1 {
		return &AppErr{
			Message: "Email cannot be empty",
		}
	}

	if !utils.EmailIsValid(h.Email) {
		return &AppErr{
			Message: "Invalid Email",
		}
	}

	if len(h.Password) < 1 {
		return &AppErr{
			Message: "Password cannot be empty",
		}
	}

	return nil
}
