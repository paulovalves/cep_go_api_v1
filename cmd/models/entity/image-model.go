package models

import "github.com/google/uuid"

type Image struct {
	Id       uuid.UUID
	PublicId uuid.UUID
	Url      string
	Category Category
}
