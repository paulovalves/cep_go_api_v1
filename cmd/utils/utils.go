package utils

import (
	"strings"

	"github.com/google/uuid"
	entity "models/entity"
)

func IsValidUUID(uuidStr string) bool {
	_, err := uuid.Parse(uuidStr)
	return err == nil
}

func ValidateCategory(category entity.Category) bool {
	if strings.TrimRight(category.Name, "\n") == "" ||
		strings.TrimRight(category.Status, "\n") == "" {
		return false
	}

	return true
}
