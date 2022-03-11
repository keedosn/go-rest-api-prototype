package entity

import (
	"time"
)

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateDate time.Time `json:"create_date"`
	ModifyDate time.Time `json:"modify_date"` // FIXME: zero-value issue
}

// func (c Category) Validate() error {
// 	return validation.ValidateStruct(&c,
// 		validation.Field(&c.Name, validation.Required, validation.Length(3, 255)),
// 	)
// }
