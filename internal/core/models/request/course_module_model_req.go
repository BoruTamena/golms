package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Modules struct {
	CourseId    int    `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}

func (m Modules) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.CourseId, is.Int),
		validation.Field(&m.Title, validation.Required, validation.Length(5, 20)),
		validation.Field(&m.Description, validation.Required, validation.Length(100, 500)),
		validation.Field(&m.Order, validation.Required, is.Int),
	)
}
