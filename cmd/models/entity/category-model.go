// package models
//
// import "github.com/google/uuid"
//
//	type Category struct {
//		id     uuid.UUID `json:"id"     gorm:"primary_key"`
//		name   string    `json:"name"   gorm:"unique"         validate:"required" binding:"required"`
//		status string    `json:"status" gorm:"default:active" validate:"required" binding:"required"`
//	}
//
//	type CategoryRequest struct {
//		Name   string `json:"name"   validate:"required" binding:"required"`
//		Status string `json:"status" validate:"required" binding:"required"`
//	}
//
//	type CategoryResponse struct {
//		ID     uuid.UUID `json:"id"`
//		Name   string    `json:"name"`
//		Status string    `json:"status"`
//	}
package models

import "github.com/google/uuid"

type Category struct {
	Id     uuid.UUID `json:"id"     gorm:"primary_key"`
	Name   string    `json:"name"                      binding:"required"`
	Status string    `json:"status"`
}
