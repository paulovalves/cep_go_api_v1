package service

import (
	"log"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	entity "models/entity"
)

func GetAllImages() entity.ResponseModel {
	var data []entity.Image

	err := DB.Table(table).Find(&data).Error
	if err != nil {
		log.Fatalf("Error while getting images: %v", err)
		return entity.SetResponse(
			nil,
			err,
			"Error while fetching images",
		)
	}

	return entity.SetResponse(
		data,
		nil,
		"success",
	)
}

func GetImageById(id uuid.UUID) entity.ResponseModel {
	var data entity.Image

	err := DB.Table(table).Where("id = ?", id).Find(&data)
	if err != nil {
		log.Fatalf("Error while getting image: %v", err)
		return entity.SetResponse(
			nil,
			err.Error,
			"error",
		)
	}

	return entity.SetResponse(
		data,
		nil,
		"success",
	)
}
