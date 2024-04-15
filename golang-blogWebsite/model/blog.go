package model

import "time"

type Blog struct {
	ID        uint      `json:"id"`
	AuthorID  uint      `json: "author_id"`
	CreatedAt time.Time `json: "createdat"`
	UpdatedAt time.Time `json: "updatedat"`
	Title     string    `json: "title"`
	Content   string    `json:"content"`
}
