package service

import (
	"log"
	"utils"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	entity "models/entity"
)

func GetAllImages() entity.ResponseModel {
	var data []entity.Image

	err := DB.Table(ImagesTable).Find(&data).Error
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

func GetImageById(id string) entity.ResponseModel {
	var data entity.Image
	if !utils.IsValidUUID(id) {
		return entity.SetResponse(
			nil,
			"Invalid UUID",
			"error",
		)
	}

	err := DB.Table(ImagesTable).Where("id = ?", id).Find(&data).Error
	if err != nil {
		log.Fatalf("Error while getting image: %v", err)
		return entity.SetResponse(
			nil,
			err.Error(),
			"error",
		)
	}

	return entity.SetResponse(
		data,
		nil,
		"success",
	)
}
