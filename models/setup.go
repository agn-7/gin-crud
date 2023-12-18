package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
)

var SQLiteDB   *gorm.DB
var PostgresDB *gorm.DB


func ConnectSQLite() error {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        return err
    }

    err = db.AutoMigrate(&Album{})
    if err != nil {
        return err
    }

    SQLiteDB = db
    return nil
}


func ConnectPostgres() error {
    dsn := "host=localhost " +
        "user=ifsguid_usr " +
        "password=root " +
        "dbname=ifsguid_db " +
        "port=5432 " +
        "sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    PostgresDB = db
    return nil
}