package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	def "git.pbiernat.dev/golang/rest-api-prototype/internal/app/definition"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/entity"
)

var CreateCategoryHandler *Handler
var DeleteCategoryHandler *Handler

func init() {
	CreateCategoryHandler = &Handler{
		Handle:   CreateCategoryHandlerFunc,
		Request:  &def.CreateCategoryRequest{},
		Response: &def.CreateCategoryResponse{},
	}

	DeleteCategoryHandler = &Handler{
		Handle:   DeleteCategoryHandlerFunc,
		Request:  &def.DeleteCategoryRequest{},
		Response: &def.DeleteCategoryResponse{},
	}
}

func CreateCategoryHandlerFunc(h *Handler, w http.ResponseWriter) (interface{}, int, error) {
	var cat = h.Request.(*def.CreateCategoryRequest)
	log.Println("Cat input:", cat)

	if err := cat.Validate(); err != nil {
		log.Println("Create category validation errors:", err)
		return nil, http.StatusUnprocessableEntity, err
	}

	return &entity.Category{
		Name:       cat.Name,
		CreateDate: time.Now(),
	}, http.StatusCreated, nil
}

func DeleteCategoryHandlerFunc(h *Handler, w http.ResponseWriter) (interface{}, int, error) {
	var cat = h.Request.(*def.DeleteCategoryRequest)
	log.Println(cat)

	id, _ := strconv.Atoi(h.Params["id"])
	log.Println(h.Params)

	if id != 1 {
		return nil, http.StatusNotFound, nil
	}

	return nil, http.StatusNoContent, nil
}
