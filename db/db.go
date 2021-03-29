package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const DB_SOURCE_NAME = "./data/sample.db"

func Get() *gorm.DB {
	return db
}

func Open() {
	_db, err := gorm.Open(sqlite.Open(DB_SOURCE_NAME), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = _db

}
