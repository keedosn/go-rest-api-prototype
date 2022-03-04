package entity

import (
	"time"
)

type Category struct {
	ID         int
	Name       string
	CreateDate time.Time
	ModifyDate time.Time
}
