package models

import "github.com/google/uuid"

type Category struct {
	id     uuid.UUID `json:"id"     gorm:"primary_key"`
	name   string    `json:"name"   gorm:"unique"         validate:"required" binding:"required"`
	status string    `json:"status" gorm:"default:active" validate:"required" binding:"required"`
}
