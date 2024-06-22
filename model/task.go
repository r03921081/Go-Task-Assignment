package model

type TaskStatus int

const (
	Incomplete TaskStatus = iota
	Complete
)

var taskID autoIncr

// Task example
type Task struct {
	ID     int        `json:"id" example:"1"`
	Name   string     `json:"name" example:"task name"`
	Status TaskStatus `json:"status" example:"1"`
}

// NewTask example
func NewTask(name string) *Task {
	return &Task{
		ID:     taskID.GetID(),
		Name:   name,
		Status: Incomplete,
	}
}
