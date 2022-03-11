package entity

import "time"

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreateDate time.Time `json:"create_date"`
	ModifyDate time.Time `json:"modify_date"` // FIXME: zero-value issue
}

var TestUser = &User{
	ID:         1,
	Username:   "test",
	Password:   "test",
	CreateDate: time.Now(),
}
