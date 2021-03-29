package tasks

import (
	"gin_test/container"
	"gin_test/db"
)

func ResisterToContainer(c *container.Container) {
	c.Register(&container.Definition{
		Name: "TaskRepository",
		Builder: func(c *container.Container) interface{} {
			return NewTaskRepository(db.Get())
		},
	})

	c.Register(&container.Definition{
		Name: "TaskController",
		Builder: func(c *container.Container) interface{} {
			return NewTaskController(c.Get("TaskRepository").(TaskRepository))
		},
	})

}
