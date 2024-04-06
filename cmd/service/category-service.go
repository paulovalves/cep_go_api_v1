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
		return category, nil
	}

	return category, nil
}
