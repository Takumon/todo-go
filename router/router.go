package router

import (
	"gin_test/controllers"
	"gin_test/db"
	"gin_test/services"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	taskService := services.TaskService{}
	taskHandler := controllers.TaskHandler{
		DB:      db.Get(),
		Service: &taskService,
	}

	r.GET("/tasks/:id", taskHandler.GetById)
	r.GET("/tasks", taskHandler.GetAll)
	r.POST("/tasks", taskHandler.Insert)
	r.PUT("/tasks/:id", taskHandler.Update)
	r.DELETE("/tasks/:id", taskHandler.Delete)
	r.Run(":8080")
}
