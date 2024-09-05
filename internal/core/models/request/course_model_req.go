package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Courses struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	InstructorId string `json:"instructor_id"`
	Language     string `json:"language"`
	Level        string `json:"level"`
	Duration     int    `json:"duration"`
}

func (cv Courses) Validate() error {
	return validation.ValidateStruct(&cv,
		validation.Field(&cv.Title, validation.Required, validation.Length(5, 20)),
		validation.Field(&cv.Description, validation.Required, validation.Length(50, 300)),
		validation.Field(&cv.Language, validation.Required),
		validation.Field(&cv.Level, validation.Required),
		validation.Field(&cv.Duration, validation.Required, is.Int),
	)
}
