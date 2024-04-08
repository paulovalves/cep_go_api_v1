package service

import (
	"log"

	entity "models/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB    *gorm.DB
	table = "public.categories"
)

func SetDB(db *gorm.DB) {
	DB = db
}

func GetCategories() entity.ResponseModel {
	var data []entity.Category

	err := DB.Table(table).Find(&data).Error
	if err != nil {
		log.Fatalf("Error while fetching categories: %v", err)
		return entity.SetResponse(nil, err, "Error while fetching categories")
	}

	return entity.SetResponse(data, nil, "success")
}

func GetCategoryById(id string) entity.ResponseModel {
	var data entity.Category

	if err := DB.Table(table).Where("id = ?", id).Find(&data).Error; err != nil {
		log.Fatalf("Error while fetching category: %v", err)
		return entity.SetResponse(nil, err, "Error while fetching category")
	}
	return entity.SetResponse(data, nil, "success")
}

func GetCategoriesByStatus(status string) entity.ResponseModel {
	var data []entity.Category

	if err := DB.Table(table).Where("status = ?", status).Find(&data).Error; err != nil {
		log.Fatalf("Error while fetching categories: %v", err)
		return entity.SetResponse(nil, err, "Error while fetching categories")
	}

	return entity.SetResponse(data, nil, "success")
}

func CreateCategory(data entity.Category) entity.ResponseModel {
	if err := DB.Table(table).Create(&data).Error; err != nil {
		log.Fatalf("Error while creating category: %v", err)
		return entity.SetResponse(nil, err, "Error while creating category")
	}

	return entity.SetResponse(data, nil, "success")
}

func UpdateCategory(category entity.Category) entity.ResponseModel {
	err := DB.Table(table).Where("id = ?", category.Id).Update(map[string]interface{}{
		"name":   category.Name,
		"status": category.Status,
	}).Error
	if err != nil {
		log.Fatalf("Error while updating category: %v", err)
		return entity.SetResponse(category, err, "Error while updating category")
	}

	return entity.SetResponse(category, nil, "success")
}

func DeleteCategory(id string) entity.ResponseModel {
	if err := DB.Table(table).Delete(id).Error; err != nil {
		log.Fatalf("Error while deleting category: %v", err)
		return entity.SetResponse(false, err, "Error while deleting category")
	}

	return entity.SetResponse(true, nil, "success")
}
