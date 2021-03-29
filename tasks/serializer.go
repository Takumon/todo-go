package tasks

import "github.com/gin-gonic/gin"

type TaskSerializer struct {
	c *gin.Context
	TaskModel
}

type TasksSerializer struct {
	c *gin.Context
	TaskModels
}

type TaskResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (s *TaskSerializer) Response() TaskResponse {
	response := TaskResponse{
		ID:        s.TaskModel.ID,
		Title:     s.TaskModel.Title,
		Content:   s.TaskModel.Content,
		CreatedAt: s.TaskModel.CreatedAt.Format("2006-01-02T15:04:05.999Z"),
		UpdatedAt: s.TaskModel.UpdatedAt.Format("2006-01-02T15:04:05.999Z"),
	}
	return response
}

func (s *TasksSerializer) Response() []TaskResponse {
	response := []TaskResponse{}
	for _, task := range s.TaskModels {
		serializer := TaskSerializer{s.c, task}
		response = append(response, serializer.Response())
	}
	return response
}
