package service

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	entity "models/entity"
)

var (
	DB    *gorm.DB
	table = "public.categories"
)

func SetDB(db *gorm.DB) {
	DB = db
}

func GetCategories() ([]entity.Category, error) {
	var data []entity.Category

	err := DB.Table(table).Find(&data).Error
	if err != nil {
		log.Fatalf("Error while fetching categories: %v", err)
		return nil, err
	}

	return data, nil
}

func GetCategoryById(id string) (entity.Category, error) {
	var data entity.Category

	if err := DB.Table(table).Where("id = ?", id).Find(&data).Error; err != nil {
		log.Fatalf("Error while fetching category: %v", err)
		return entity.Category{}, err
	}

	return data, nil
}

func GetCategoriesByStatus(status string) ([]entity.Category, error) {
	var data []entity.Category

	if err := DB.Table(table).Where("status = ?", status).Find(&data).Error; err != nil {
		log.Fatalf("Error while fetching categories: %v", err)
		return nil, err
	}

	return data, nil
}

func CreateCategory(category entity.Category) (entity.Category, error) {
	log.Printf("Category: %v", category)
	// if utils.ValidateCategory(category) {
	// 	log.Fatalf("Category is not valid %v", category)
	// 	return entity.Category{}, errors.New("Category is not valid")
	// }
	if err := DB.Table(table).Create(&category).Error; err != nil {
		log.Fatalf("Error while creating category: %v", err)
		return category, err
	}

	return category, nil
}

func UpdateCategory(category entity.Category) (entity.Category, error) {
	err := DB.Table(table).Where("id = ?", category.Id).Update(map[string]interface{}{
		"name":   category.Name,
		"status": category.Status,
	}).Error
	if err != nil {
		log.Fatalf("Error while updating category: %v", err)
		return category, err
	}

	return category, nil
}

func DeleteCategory(id string) (bool, error) {
	if err := DB.Table(table).Delete(id).Error; err != nil {
		log.Fatalf("Error while deleting category: %v", err)
		return false, err
	}

	return true, nil
}
