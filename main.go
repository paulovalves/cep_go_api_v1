package main

import (
	"service"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	service.Start(":8080")
}
