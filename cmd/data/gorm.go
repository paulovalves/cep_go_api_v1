package data

// Description: This file contains the database connection logic using GORM.

import (
	"log"
	"service"

	"github.com/jinzhu/gorm"
)

func Init(DB *gorm.DB) *gorm.DB {
	service.SetDB(DB)
	log.Println("Connecting to database...")
	var err error
	DB, err = gorm.Open(
		"postgres",
		"host=postgres port=5432 user=postgres dbname=postgres password=secret sslmode=disable",
	)
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	DB = DB.Unscoped()
	DB.LogMode(true)
	log.Println("Connected to database on port 5432")

	return DB
}
