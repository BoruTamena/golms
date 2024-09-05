package params

import validation "github.com/go-ozzo/ozzo-validation"

type CourseParams struct {
	// list of the query paramaters
	Category string `form:"category,omitempty"`
	Plan     string `form:"plan,omitempty"`
	Language string `form:"language,omitempty"`
	Level    string `form:"level,omitempty"`
}

func (cp CourseParams) Validate() error {
	return validation.ValidateStruct(&cp,

		validation.Field(&cp.Category, validation.Length(5, 10)),
		validation.Field(&cp.Plan, validation.Length(4, 7)),
		validation.Field(&cp.Language, validation.Length(4, 10)),
		validation.Field(&cp.Language, validation.In("beginner", "intermidate", "advanced")),
	)
}
