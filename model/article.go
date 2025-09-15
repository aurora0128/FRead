package model

import (
	"time"
)

type Article struct {
	ID        string    `json:"id"`
	PLATFORM  string    `json:"platform"`
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	Thumbnail string    `json:"thumbnail"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type ArticleList struct {
	//id->*article
}
