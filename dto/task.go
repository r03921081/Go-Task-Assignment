package dto

import "github.com/andyliao/task-homework/model"

// CreateTaskRequest example
type CreateTaskRequest struct {
	Name string `json:"name" example:"task name"`
}

// PutTaskRequest example
type PutTaskRequest struct {
	ID     int              `json:"id" example:"1"`
	Name   string           `json:"name" example:"task name"`
	Status model.TaskStatus `json:"status" example:"1"`
}
