package handler

import (
	"log"
	"net/http"
	"time"

	def "git.pbiernat.dev/golang/rest-api-prototype/internal/app/definition"
	"git.pbiernat.dev/golang/rest-api-prototype/internal/app/entity"
)

var CreateArticleHandler *Handler

func init() {
	CreateArticleHandler = &Handler{
		Handle:   CreateArticleHandlerFunc,
		Request:  &def.CreateArticleRequest{},
		Response: &def.CreateArticleResponse{},
	}
}

func CreateArticleHandlerFunc(h *Handler, w http.ResponseWriter) (interface{}, int, error) {
	var art = h.Request.(*def.CreateArticleRequest)
	log.Println(art)

	return &entity.Article{
		ID:         1,
		CategoryID: 1,
		Title:      "Dummy article",
		Intro:      "Intro",
		Text:       "Text",
		CreateDate: time.Now(),
	}, http.StatusCreated, nil
}
