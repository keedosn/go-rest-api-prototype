package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/entity"
	validation "github.com/go-ozzo/ozzo-validation"
)

var CreateCategoryHandler *Handler
var DeleteCategoryHandler *Handler

func init() {
	CreateCategoryHandler = &Handler{
		Handle:   CreateCategoryHandlerFunc,
		Request:  &CreateCategoryRequest{},
		Response: &CreateCategoryResponse{},
	}

	DeleteCategoryHandler = &Handler{
		Handle:   DeleteCategoryHandlerFunc,
		Request:  &DeleteCategoryRequest{},
		Response: &DeleteCategoryResponse{},
	}
}

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

func CreateCategoryHandlerFunc(h *Handler, w http.ResponseWriter) (interface{}, int, error) {
	var cat = h.Request.(*CreateCategoryRequest)
	log.Println("Cat input:", cat)

	if err := cat.Validate(); err != nil {
		log.Println("Create category validation errors:", err)
		return nil, http.StatusUnprocessableEntity, err
	}

	return &entity.Category{
		Name:       cat.Name,
		CreateDate: time.Now(),
	}, http.StatusCreated, nil

	// return &CreateCategoryResponse{
	// 	Data: &entity.Category{
	// 		Name:       "Dummy category",
	// 		CreateDate: time.Now(),
	// 	},
	// }, http.StatusCreated, nil
}

func DeleteCategoryHandlerFunc(h *Handler, w http.ResponseWriter) (interface{}, int, error) {
	var cat = h.Request.(*DeleteCategoryRequest)
	log.Println(cat)

	id, _ := strconv.Atoi(h.Params["id"])
	log.Println(h.Params)

	if id != 1 {
		return &DeleteCategoryResponse{}, http.StatusNotFound, nil
	}

	return &DeleteCategoryResponse{}, http.StatusNoContent, nil
}
