package controllers

import (
	"database/sql"
	"gin_test/models"
	"gin_test/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	DB      *sql.DB
	Service *services.TaskService
}

func (handler *TaskHandler) GetById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	todo, err := handler.Service.GetById(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, todo)
}

func (handler *TaskHandler) GetAll(c *gin.Context) {

	todos, err := handler.Service.GetAll()
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, todos)
}

func (handler *TaskHandler) Insert(c *gin.Context) {
	var params models.RequestParamsInsertTodo
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Fatal(err.Error())
	}

	tx, err := handler.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	name := params.Name
	id, err := handler.Service.Insert(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	todo, err := handler.Service.GetById(int(id))
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, todo)
}

func (handler *TaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	var params models.Todo
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Fatal(err.Error())
	}

	tx, err := handler.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	if err := handler.Service.Update(id, params); err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	todo, err := handler.Service.GetById(id)
	if err != nil {
		log.Fatal(err.Error())
	}

	c.JSON(http.StatusOK, todo)
}

func (handler *TaskHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	tx, err := handler.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}

	if err := handler.Service.Delete(id); err != nil {
		log.Fatal(err.Error())
	}
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
