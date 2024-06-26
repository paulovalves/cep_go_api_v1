// Description: This file contains the model for the image entity.
// Autor: Paulo Alves
package models

import "github.com/google/uuid"

type Image struct {
	Id          uuid.UUID `json:"id"          gorm:"primary_key"`
	Alt         string    `json:"alt"`
	Description string    `json:"description"`
	Filename    string    `json:"filename"`
	PublicId    string    `json:"public_id"`
	Status      string    `json:"status"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	CategoryID  uuid.UUID `json:"category_id"`
	Category    *Category `json:"category"    gorm:"foreignKey:CategoryID"`
}
