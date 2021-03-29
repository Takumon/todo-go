package main

import (
	"gin_test/container"
	"gin_test/controllers"
	"gin_test/db"
	"gin_test/repositories"
	"gin_test/router"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.Open()
	db.Init()

	c := container.NewContainer()
	c.Register(&container.Definition{
		Name: "TaskRepository",
		Builder: func(c *container.Container) interface{} {
			return repositories.NewTaskRepository(db.Get())
		},
	})

	c.Register(&container.Definition{
		Name: "TaskController",
		Builder: func(c *container.Container) interface{} {
			return controllers.NewTaskController(c.Get("TaskRepository").(repositories.TaskRepository))
		},
	})

	router.Router(c)
}
