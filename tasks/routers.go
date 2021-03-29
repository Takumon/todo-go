package tasks

import (
	"gin_test/container"

	"github.com/gin-gonic/gin"
)

func TasksRegister(r *gin.RouterGroup, container *container.Container) {
	c := container.Get("TaskController").(TaskController)

	r.GET("", c.GetAll)
	r.GET("/:id", c.GetById)
	r.POST("", c.Insert)
	r.PUT("/:id", c.Update)
	r.DELETE("/:id", c.Delete)
}
