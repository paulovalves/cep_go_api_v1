package models

import "github.com/google/uuid"

type Image struct {
	Id       uuid.UUID `json:"id"        gorm:"primary_key"`
	PublicId uuid.UUID `json:"public_id"`
	Url      string    `json:"url"`
	Category Category  `json:"category"`
}
