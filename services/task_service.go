package services

import (
	"gin_test/models"

	"gorm.io/gorm"
)

type TaskService struct {
}

func (service *TaskService) GetById(db *gorm.DB, id int) (models.Task, error) {
	var task models.Task
	db.First(&task, id)
	return task, nil
}

func (service *TaskService) GetAll(db *gorm.DB) (models.Tasks, error) {
	var tasks models.Tasks
	db.Find(&tasks)
	return tasks, nil
}

func (service *TaskService) Insert(db *gorm.DB, title string, content string) (*models.Task, error) {
	tx := db.Begin()

	task := models.Task{
		Title:   title,
		Content: content,
	}
	if err := tx.Create(&task).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &task, nil
}

func (service *TaskService) Update(db *gorm.DB, id int, title string, content string, done bool) (*models.Task, error) {
	tx := db.Begin()

	var task models.Task
	tx.First(&task, id)
	newTask := models.Task{
		Title:   title,
		Content: content,
		Done:    done,
	}
	if err := tx.Model(&task).Updates(newTask).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &task, nil
}

func (service *TaskService) Delete(db *gorm.DB, id int) error {
	tx := db.Begin()

	if err := tx.Delete(&models.Task{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
