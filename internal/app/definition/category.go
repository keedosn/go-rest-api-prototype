package definition

import (
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/entity"
	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

func (c CreateCategoryRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(3, 255)),
	)
}

type CreateCategoryResponse struct {
	Data *entity.Category `json:"data"`
	Err  string           `json:"err,omitempty"` // FIXME: omitempty on/off?
}

type DeleteCategoryRequest struct {
}

type DeleteCategoryResponse struct {
}
