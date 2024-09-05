package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Lessons struct {
	ModuleId int    `jsonL:"module_id"`
	Title    string `json:"title"`
	VideoUrl string `json:"video_url" `
	Order    int    `json:"order"`
}

func (lv Lessons) Validate() error {
	return validation.ValidateStruct(&lv,
		validation.Field(&lv.Title, validation.Required, validation.Length(5, 20)),
		validation.Field(&lv.VideoUrl, validation.Required),
		validation.Field(&lv.Order, validation.Required, is.Int),
	)
}
