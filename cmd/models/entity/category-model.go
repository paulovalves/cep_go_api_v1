// Description: This file contains the model for the category.
// Autor: Paulo Alves
package models

import "github.com/google/uuid"

// Category model
type Category struct {
	Id     uuid.UUID `json:"id"     gorm:"primary_key"`
	Name   string    `json:"name"                      binding:"required"`
	Status string    `json:"status"`
}
