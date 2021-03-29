package controllers

import (
	"gin_test/models"
	"gin_test/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	DB      *gorm.DB
	Service *services.TaskService
}

func (handler *TaskHandler) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	task, err := handler.Service.GetById(handler.DB, id)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, task)
}

func (handler *TaskHandler) GetAll(c *gin.Context) {

	tasks, err := handler.Service.GetAll(handler.DB)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, tasks)
}

func (handler *TaskHandler) Insert(c *gin.Context) {
	var params models.Task
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Fatal(err.Error())
	}

	task, err := handler.Service.Insert(handler.DB, params.Title, params.Content)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, task)
}

func (handler *TaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	var params models.Task
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Fatal(err.Error())
	}

	task, err := handler.Service.Update(handler.DB, id, params.Title, params.Content, params.Done)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, task)
}

func (handler *TaskHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := handler.Service.Delete(handler.DB, id); err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
