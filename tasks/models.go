package tasks

import "gorm.io/gorm"

type TaskModel struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Done    bool   `json:"done" gorm:"default:false"`
}

type TaskModels []TaskModel
