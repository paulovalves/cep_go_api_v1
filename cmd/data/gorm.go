package data

// Description: This file contains the database connection logic using GORM.

import (
	"log"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres dbname=postgres password=secret sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	DB = DB.Unscoped()
}
