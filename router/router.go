package router

import (
	"gin_test/container"
	"gin_test/controllers"

	"github.com/gin-gonic/gin"
)

func Router(c *container.Container) {
	r := gin.Default()

	controller := c.Get("TaskController").(controllers.TaskController)

	r.GET("/tasks/:id", controller.GetById)
	r.GET("/tasks", controller.GetAll)
	r.POST("/tasks", controller.Insert)
	r.PUT("/tasks/:id", controller.Update)
	r.DELETE("/tasks/:id", controller.Delete)
	r.Run(":8080")
}
