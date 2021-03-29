package controllers

import (
	model "gin_test/models"
	repositories "gin_test/repositories"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Insert(c *gin.Context)
	Delete(c *gin.Context)
}

func NewTaskController(repo repositories.TaskRepository) TaskController {
	return &taskController{repo}
}

type taskController struct {
	repo repositories.TaskRepository
}

func (h *taskController) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	task, err := h.repo.GetById(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, task)
}

func (h *taskController) GetAll(c *gin.Context) {

	tasks, err := h.repo.GetAll()
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *taskController) Insert(c *gin.Context) {
	var params model.Task
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Fatal(err.Error())
	}

	task, err := h.repo.Insert(params.Title, params.Content)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, task)
}

func (h *taskController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	var params model.Task
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Fatal(err.Error())
	}

	task, err := h.repo.Update(id, params.Title, params.Content, params.Done)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, task)
}

func (h *taskController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := h.repo.Delete(id); err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
