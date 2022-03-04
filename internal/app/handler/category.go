package handler

import (
	"net/http"
	"time"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/entity"
)

var CreateCategoryHandler *Handler

func init() {
	CreateCategoryHandler = &Handler{
		Handle:   CreateCategoryHandlerFunc,
		Request:  &CreateCategoryRequest{},
		Response: &CreateCategoryResponse{},
	}
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type CreateCategoryResponse struct {
	Status string           `json:"status"`
	Data   *entity.Category `json:"data"`
	Err    string           `json:"err,omitempty"`
}

func CreateCategoryHandlerFunc(env *Env, req interface{}, w http.ResponseWriter) (interface{}, error) {
	// var cat = req.(*CreateCategoryRequest)
	// doSthWith(cat)

	return &CreateCategoryResponse{
		Status: http.StatusText(http.StatusOK),
		Data: &entity.Category{
			Name:       "Dummy category",
			CreateDate: time.Now(),
		},
	}, nil
}
