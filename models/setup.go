package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() error {
    database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    err = database.AutoMigrate(&Album{})
    if err != nil {
        return err
    }

    DB = database
    return nil
}
