package service

import (
	entity "models/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

var table = "public.categories"

func GetCategories() ([]entity.Category, error) {
	var data []entity.Category

	err := DB.Table(table).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
