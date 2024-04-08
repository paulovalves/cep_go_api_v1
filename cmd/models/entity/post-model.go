package models

import "github.com/google/uuid"

type Post struct {
	Id        uuid.UUID
	Title     string
	Text      string
	CreatedAt string
	UpdatedAt string
	Image     Image
	Category  Category
	CreatedBy User
	UpdatedBy User
}
