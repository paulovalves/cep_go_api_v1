package service

import (
	"log"

	entity "models/entity"
	"utils"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllImages() entity.ResponseModel {
	var data []entity.Image

	err := DB.Table(ImagesTable).Preload("Category").Find(&data).Error
	if err != nil {
		log.Printf("Error while getting images: %v", err)
		return entity.SetResponse(
			nil,
			err.Error(),
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

	err := DB.Table(ImagesTable).Preload("Category").Where("id = ?", id).Find(&data).Error
	if err != nil {
		log.Printf("Error while getting image: %v", err)
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

func GetImagesByCategory(id string) entity.ResponseModel {
	var data []entity.Image
	if !utils.IsValidUUID(id) {
		return entity.SetResponse(
			nil,
			"Invalid UUID",
			"error",
		)
	}

	err := DB.Table(ImagesTable).
		Preload("Category").
		Where("category_id = ?", id).
		Find(&data).
		Error
	if err != nil {
		log.Printf("Error while getting images: %v", err)
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

func GetImagesByStatus(status string) entity.ResponseModel {
	var data []entity.Image
	err := DB.Table(ImagesTable).Preload("Category").Where("status = ?", status).Find(&data).Error
	if err != nil {
		log.Printf("Error while getting image: %v", err)
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
