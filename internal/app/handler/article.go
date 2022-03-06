package handler

import (
	"log"
	"net/http"
	"time"

	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/entity"
)

var CreateArticleHandler *Handler

func init() {
	CreateArticleHandler = &Handler{
		Handle:   CreateArticleHandlerFunc,
		Request:  &CreateArticleRequest{},
		Response: &CreateArticleResponse{},
	}
}

type CreateArticleRequest struct {
	CategoryID int    `json:"category_id"`
	Title      string `json:"title"`
	Intro      string `json:"intro"`
	Text       string `json:"text"`
}

type CreateArticleResponse struct {
	Status string          `json:"status"`
	Data   *entity.Article `json:"data"`
	Err    string          `json:"err,omitempty"`
}

func CreateArticleHandlerFunc(h *Handler, w http.ResponseWriter) (interface{}, int, error) {
	var art = h.Request.(*CreateArticleRequest)
	log.Println(art)

	return &entity.Article{
		ID:         1,
		CategoryID: 1,
		Title:      "Dummy article",
		Intro:      "Intro",
		Text:       "Text",
		CreateDate: time.Now(),
	}, http.StatusCreated, nil

	// return &CreateArticleResponse{
	// 	Status: http.StatusText(http.StatusOK),
	// 	Data: &entity.Article{
	// 		ID:         1,
	// 		CategoryID: 1,
	// 		Title:      "Dummy article",
	// 		Intro:      "Intro",
	// 		Text:       "Text",
	// 		CreateDate: time.Now(),
	// 	},
	// }, http.StatusCreated, nil
}
