package data

// Description: This file contains the database connection logic using GORM.

import (
	"service"

	entity "models/entity"

	"github.com/jinzhu/gorm"
)

func Init(DB *gorm.DB) *gorm.DB {
	service.SetDB(DB)
	logs := entity.LogModel{}

	logs.Print("Connecting to database...", "info")
	// log.Printf("%v - Connecting to database...", time.Now().Format(time.RFC3339))
	var err error
	DB, err = gorm.Open(
		"postgres",
		"host=postgres port=5432 user=postgres dbname=postgres password=secret sslmode=disable",
	)
	if err != nil {
		logs.Print("Error while connecting to database", "danger")
		// stop the application if the connection to the database fails
		panic(err)
	}

	DB = DB.Unscoped()
	DB.LogMode(true)
	logs.Print("Connected to database on port 5432", "success")
	// log.Printf("%v - Connected to database on port 5432", time.Now().Format(time.RFC3339))

	return DB
}
