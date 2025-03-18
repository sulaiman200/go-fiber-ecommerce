package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Database connection Successful!")
}
