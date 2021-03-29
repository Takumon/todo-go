package tasks

import (
	"gin_test/common"

	"github.com/gin-gonic/gin"
)

type TaskModelValidator struct {
	Task struct {
		Title   string `form:"title" json:"title" binding:"required,max=50"`
		Content string `form:"content" json:"content" binding:"required,max=1024"`
		Done    bool   `form:"done" json:"done"`
	} `json:"task"`
	taskModel TaskModel `json:"-"`
}

func (self *TaskModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.taskModel.Title = self.Task.Title
	self.taskModel.Content = self.Task.Content
	self.taskModel.Done = self.Task.Done

	return nil
}

func NewTaskModelValidator() TaskModelValidator {
	v := TaskModelValidator{}
	return v
}

func NewTaskModelValidatorFillWith(model TaskModel) TaskModelValidator {
	v := NewTaskModelValidator()
	v.Task.Title = model.Title
	v.Task.Content = model.Content
	v.Task.Done = model.Done
	return v
}
