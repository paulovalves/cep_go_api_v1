// Description: This file contains the service functions for the image model.
// The service functions are responsible for interacting with the database and
// performing the necessary operations. The service functions are defined in the
// service package and are imported into the controller functions to be used by
// the server. The service functions are called by the controller functions to
// perform the CRUD operations on the database.
package service

import (
	"log"

	entity "models/entity"
	"utils"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
* GetAllImages function is responsible for getting all images from the database.
* It returns a response model containing the images if the operation is successful.
* If the operation fails, it returns an error response.
* Return: [entity.ResponseModel] - response model
 */
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

/*
* GetImageById function is responsible for getting an image by id from the database.
* It takes the id of the image as a parameter and returns a response model containing
* the image if the operation is successful. If the operation fails, it returns an error response.
* Param: id - image id
* Return: [entity.ResponseModel] - response model
 */
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

/*
* GetImagesByCategory function is responsible for getting images by category from the database.
* It takes the category id as a parameter and returns a response model containing the images
* if the operation is successful. If the operation fails, it returns an error response.
* Param: id - category id
* Return: [entity.ResponseModel] - response model
 */
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

/*
* GetImagesByStatus function is responsible for getting images by status from the database.
* It takes the status as a parameter and returns a response model containing the images
* if the operation is successful. If the operation fails, it returns an error response.
* Param: status - image status
* Return: [entity.ResponseModel] - response model
 */
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

/*
* GetImagesByDescription function is responsible for getting images by description from the database.
* It takes the description as a parameter and returns a response model containing the images
* if the operation is successful. If the operation fails, it returns an error response.
* Param: description - image description
* Return: [entity.ResponseModel] - response model
 */
func GetImagesByDescription(description string) entity.ResponseModel {
	var data []entity.Image
	err := DB.Table(ImagesTable).
		Preload("Category").
		Where("description LIKE ?", "%"+description+"%").
		Find(&data).
		Error
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

/*
* CreateImage function is responsible for creating an image in the database.
* It takes the image data as a parameter and returns a response model containing
* the image if the operation is successful. If the operation fails, it returns an error response.
* Param: data - image data
* Return: [entity.ResponseModel] - response model
 */
func CreateImage(data entity.Image) entity.ResponseModel {
	err := DB.Table(ImagesTable).Create(&data).Error
	if err != nil {
		log.Printf("Error while creating image: %v", err)
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
