package tasks

import (
	"errors"
	"gin_test/common"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskController interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
}

func NewTaskController(repo TaskRepository) TaskController {
	return &taskController{repo}
}

type taskController struct {
	repo TaskRepository
}

func (h *taskController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("tasks", errors.New("Invalid id")))
	}

	task, err := h.repo.FindOnTask(&TaskModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		// TODO エラー別のハンドリング
		c.JSON(http.StatusNotFound, common.NewError("tasks", errors.New("Invalid id")))
	}

	serializer := TaskSerializer{c, task}
	c.JSON(http.StatusOK, gin.H{"task": serializer.Response()})
}

func (h *taskController) GetAll(c *gin.Context) {

	tasks, err := h.repo.GetAll()
	if err != nil {
		// TODO エラー別のハンドリング
		c.JSON(http.StatusNotFound, common.NewError("tasks", errors.New("Invalid id")))
	}

	serializer := TasksSerializer{c, tasks}
	c.JSON(http.StatusOK, gin.H{"tasks": serializer.Response()})
}

func (h *taskController) Insert(c *gin.Context) {

	modelValidator := NewTaskModelValidator()
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}

	if err := h.repo.Insert(&modelValidator.taskModel); err != nil {
		// TODO エラー別のハンドリング
		c.JSON(http.StatusBadRequest, common.NewError("database", err))
		return
	}

	serializer := TaskSerializer{c, modelValidator.taskModel}
	c.JSON(http.StatusCreated, gin.H{"article": serializer.Response()})
}

func (h *taskController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("tasks", errors.New("Invalid id")))
		return
	}

	model, err := h.repo.FindOnTask(&TaskModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("tasks", errors.New("Invalid id")))
		return
	}

	modelValidator := NewTaskModelValidatorFillWith(model)
	if err := modelValidator.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, common.NewValidatorError(err))
		return
	}

	if err := h.repo.Update(&model, modelValidator.taskModel); err != nil {
		// TODO エラー別のハンドリング
		c.JSON(http.StatusBadRequest, common.NewError("database", err))
		return
	}

	serializer := TaskSerializer{c, modelValidator.taskModel}
	c.JSON(http.StatusCreated, gin.H{"task": serializer.Response()})
}

func (h *taskController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("tasks", errors.New("Invalid id")))
		return
	}

	model, err := h.repo.FindOnTask(&TaskModel{Model: gorm.Model{ID: uint(id)}})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("tasks", errors.New("Invalid id")))
		return
	}

	if err := h.repo.Delete(model.ID); err != nil {
		// TODO エラー別のハンドリング
		c.JSON(http.StatusBadRequest, common.NewError("database", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
