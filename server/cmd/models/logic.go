package models

import (
	"time"
)

type CachedExample struct {
	Id 	int `json:"id"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}