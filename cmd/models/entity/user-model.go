package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"       gorm:"primary_key"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
}
