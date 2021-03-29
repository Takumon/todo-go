package repositories

import (
	"gin_test/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetById(id int) (models.Task, error)
	GetAll() (models.Tasks, error)
	Insert(title string, content string) (*models.Task, error)
	Update(id int, title string, content string, done bool) (*models.Task, error)
	Delete(id int) error
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{*db}
}

type taskRepository struct {
	db gorm.DB
}

func (repo *taskRepository) GetById(id int) (models.Task, error) {
	var task models.Task
	if err := repo.db.First(&task, id).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (repo *taskRepository) GetAll() (models.Tasks, error) {
	var tasks models.Tasks
	repo.db.Find(&tasks)
	return tasks, nil
}

func (repo *taskRepository) Insert(title string, content string) (*models.Task, error) {
	tx := repo.db.Begin()

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

func (repo *taskRepository) Update(id int, title string, content string, done bool) (*models.Task, error) {
	tx := repo.db.Begin()

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

func (repo *taskRepository) Delete(id int) error {
	tx := repo.db.Begin()

	if err := tx.Delete(&models.Task{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
