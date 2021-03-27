package router

import (
	"gin_test/controllers"
	"gin_test/db"
	"gin_test/services"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	taskService := services.TaskService{
		DB: db.Get(),
	}
	taskHandler := controllers.TaskHandler{
		DB:      db.Get(),
		Service: &taskService,
	}

	r.GET("/todos/:id", taskHandler.GetById)
	r.GET("/todos", taskHandler.GetAll)
	r.POST("/todos", taskHandler.Insert)
	r.PUT("/todos/:id", taskHandler.Update)
	r.DELETE("/todos/:id", taskHandler.Delete)
	r.Run(":8080")
}
