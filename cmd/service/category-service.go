// Description: This file contains the service functions for category entity.
// The service functions are responsible for interacting with the database and
// performing the necessary operations. The service functions are defined in the
// service package and are imported into the controller functions to be used by the server.
// The service functions are called by the controller functions to perform the CRUD operations on the database.
// The service functions return a response model which contains the data, error, and message.
package service

import (
	"log"

	entity "models/entity"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
* GetCategories function is responsible for fetching all categories from the database.
* It returns a response model which contains the data, error, and message.
* If the function encounters an error while fetching the categories, it returns an error response.
* If the function fetches the categories successfully, it returns a success response.
* return [entity.ResponseModel] - response model containing the data, error, and message
 */
func GetCategories() entity.ResponseModel {
	var data []entity.Category

	err := DB.Table(CategoriesTable).Find(&data).Error
	if err != nil {
		log.Printf("Error while fetching categories: %v", err)
		return entity.SetResponse(nil, err, "Error while fetching categories")
	}

	return entity.SetResponse(data, nil, "success")
}

// GetCategoryById function is responsible for fetching a category by id from the database.
// It returns a response model which contains the data, error, and message.
// If the function encounters an error while fetching the category, it returns an error response.
// If the function fetches the category successfully, it returns a success response.
// @param {string} - id
// return [entity.ResponseModel] - response model containing the data, error, and message
func GetCategoryById(id string) entity.ResponseModel {
	var data entity.Category

	if err := DB.Table(CategoriesTable).Where("id = ?", id).Find(&data).Error; err != nil {
		log.Printf("Error while fetching category: %v", err)
		return entity.SetResponse(nil, err, "Error while fetching category")
	}
	return entity.SetResponse(data, nil, "success")
}

// GetCategoriesByStatus function is responsible for fetching categories by status from the database.
// It returns a response model which contains the data, error, and message.
// If the function encounters an error while fetching the categories, it returns an error response.
// If the function fetches the categories successfully, it returns a success response.
// @param {string} - status
// return [entity.ResponseModel] - response model containing the data, error, and message
func GetCategoriesByStatus(status string) entity.ResponseModel {
	var data []entity.Category

	if err := DB.Table(CategoriesTable).Where("status = ?", status).Find(&data).Error; err != nil {
		log.Printf("Error while fetching categories: %v", err)
		return entity.SetResponse(nil, err, "Error while fetching categories")
	}

	return entity.SetResponse(data, nil, "success")
}

// CreateCategory function is responsible for creating a new category in the database.
// It returns a response model which contains the data, error, and message.
// If the function encounters an error while creating the category, it returns an error response.
// If the function creates the category successfully, it returns a success response.
// @param {entity.Category} - data
// return [entity.ResponseModel] - response model containing the data, error, and message
func CreateCategory(data entity.Category) entity.ResponseModel {
	if err := DB.Table(CategoriesTable).Create(&data).Error; err != nil {
		log.Printf("Error while creating category: %v", err)
		return entity.SetResponse(nil, err, "Error while creating category")
	}

	return entity.SetResponse(data, nil, "success")
}

// UpdateCategory function is responsible for updating a category in the database.
// It returns a response model which contains the data, error, and message.
// If the function encounters an error while updating the category, it returns an error response.
// If the function updates the category successfully, it returns a success response.
// @param {entity.Category} - category
// return [entity.ResponseModel] - response model containing the data, error, and message
func UpdateCategory(category entity.Category) entity.ResponseModel {
	err := DB.Table(CategoriesTable).Where("id = ?", category.Id).Update(map[string]interface{}{
		"name":   category.Name,
		"status": category.Status,
	}).Error
	if err != nil {
		log.Printf("Error while updating category: %v", err)
		return entity.SetResponse(category, err, "Error while updating category")
	}

	return entity.SetResponse(category, nil, "success")
}

// DeleteCategory function is responsible for soft deleting - deactivate - a category from the database.
// It returns a response model which contains the data, error, and message.
// If the function encounters an error while deleting the category, it returns an error response.
// If the function deletes the category successfully, it returns a success response.
// @param {string} - id
// return [entity.ResponseModel] - response model containing the data, error, and message
func DeleteCategory(id string) entity.ResponseModel {
	if err := DB.Table(CategoriesTable).Delete(id).Error; err != nil {
		log.Printf("Error while deleting category: %v", err)
		return entity.SetResponse(false, err, "Error while deleting category")
	}

	return entity.SetResponse(true, nil, "success")
}
