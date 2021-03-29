package main

import (
	"gin_test/container"
	"gin_test/db"
	"gin_test/tasks"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&tasks.TaskModel{})
}
func addInitialTableData(db *gorm.DB) {
	tasks := tasks.TaskModels{
		{Title: "sample task 1", Content: "sample task content 1", Done: false},
		{Title: "sample task 2", Content: "sample task content 2", Done: true},
		{Title: "sample task 3", Content: "sample task content 3", Done: false},
		{Title: "sample task 4", Content: "sample task content 4", Done: false},
		{Title: "sample task 5", Content: "sample task content 5", Done: false},
	}

	db.Create(tasks)
	log.Println("has inserted initial data")
}

func main() {
	db.Open()
	Migrate(db.Get())
	addInitialTableData(db.Get())

	r := gin.Default()
	v1 := r.Group("/api/v1")

	c := container.NewContainer()
	tasks.ResisterToContainer(c)
	tasks.TasksRegister(v1.Group("/tasks"), c)

	r.Run(":8080")

}
