package model

import "time"

type Blog struct {
	ID uint `json:"id"`
	// AuthorID  uint      `json: "author_id"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json: "createdat"`
	UpdatedAt time.Time `json: "updatedat"`
	Title     string    `json: "title"`
	Content   string    `json:"content"`
}
