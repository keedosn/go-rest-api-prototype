package definition

import "git.pbiernat.dev/golang/rest-api-prototype/internal/app/entity"

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
