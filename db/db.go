package db

import (
	"log"

	"gin_test/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const DB_SOURCE_NAME = "./data/sample.db"

func addInitialTableData() {
	tasks := models.Tasks{
		{Title: "sample task 1", Content: "sample task content 1", Done: false},
		{Title: "sample task 2", Content: "sample task content 2", Done: true},
		{Title: "sample task 3", Content: "sample task content 3", Done: false},
		{Title: "sample task 4", Content: "sample task content 4", Done: false},
		{Title: "sample task 5", Content: "sample task content 5", Done: false},
	}

	db.Create(tasks)
	log.Println("has inserted initial data")
}

func Get() *gorm.DB {
	return db
}

func Open() {
	_db, err := gorm.Open(sqlite.Open(DB_SOURCE_NAME), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db = _db

	db.AutoMigrate(&models.Task{})
}

func Init() {
	addInitialTableData()
}
