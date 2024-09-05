package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var NameRule = []validation.Rule{
	validation.Required,
	validation.Length(4, 6),
}

type User struct {
	FristName string `json:"fristname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Password  string `json:"password"`
}

func (uv User) Validate() error {
	return validation.ValidateStruct(&uv,
		validation.Field(&uv.FristName, NameRule...),
		validation.Field(&uv.LastName, NameRule...),
		validation.Field(&uv.Email, validation.Required, is.Email),
		validation.Field(&uv.Role, validation.Required),
		validation.Field(&uv.Password, validation.Required, validation.Length(1, 7)),
	)
}
