package entity

import (
	"time"
)

type Article struct {
	ID         int       `json:"id"`
	CategoryID int       `json:"category_id"`
	Title      string    `json:"title"`
	Intro      string    `json:"intro"`
	Text       string    `json:"text"`
	CreateDate time.Time `json:"create_date"`
	ModifyDate time.Time `json:"modify_date"`
}
