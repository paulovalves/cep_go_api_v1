package models

import "github.com/google/uuid"

type Post struct {
	Id        uuid.UUID `json:"id"         gorm:"primary_key"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Image     Image     `json:"image"`
	Category  Category  `json:"category"`
	CreatedBy User      `json:"created_by"`
	UpdatedBy User      `json:"updated_by"`
}
