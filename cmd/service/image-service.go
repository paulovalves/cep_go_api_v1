package service

import (
	"log"

	entity "models/entity"

	_ "github.com/jinzhu/gorm/dialects/postgres"
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
