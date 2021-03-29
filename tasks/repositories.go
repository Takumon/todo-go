package tasks

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	FindOnTask(interface{}) (TaskModel, error)
	GetAll() (TaskModels, error)
	Insert(model *TaskModel) error
	Update(model *TaskModel, value TaskModel) error
	Delete(id uint) error
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{*db}
}

type taskRepository struct {
	db gorm.DB
}

func (repo *taskRepository) FindOnTask(condition interface{}) (TaskModel, error) {
	var model TaskModel
	tx := repo.db.Begin()
	tx.Where(condition).First(&model)
	if tx.Error != nil {
		tx.Rollback()
		return model, tx.Error
	}
	tx.Commit()
	return model, nil
}

func (repo *taskRepository) GetAll() (TaskModels, error) {
	var tasks TaskModels
	repo.db.Find(&tasks)
	return tasks, nil
}

func (repo *taskRepository) Insert(model *TaskModel) error {
	tx := repo.db.Begin()

	if err := tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (repo *taskRepository) Update(model *TaskModel, values TaskModel) error {
	tx := repo.db.Begin()

	if err := tx.Model(&model).Updates(values).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (repo *taskRepository) Delete(id uint) error {
	tx := repo.db.Begin()

	if err := tx.Delete(&TaskModel{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
