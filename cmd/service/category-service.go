package service

import (
	"fmt"
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
	fmt.Println("DB", DB)

	err := DB.Table("public.categories").Find(&data).Error
	if err != nil {
		log.Fatalf("Error while fetching categories: %v", err)
		return nil, err
	}

	return data, nil
}

func GetCategoryById(id string) (entity.Category, error) {
	var data entity.Category
	err := DB.Table(table).Where("id = ?", id).Find(&data).Error
	if err != nil {
		log.Fatalf("Error while fetching category: %v", err)
		return entity.Category{}, err
	}

	return data, nil
}
