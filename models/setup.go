package models

import (
	"gorm.io/gorm"
	_ "gorm.io/driver/sqlite"
)


var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
			panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Album{})
	if err != nil {
			return
	}

	DB = database
}
