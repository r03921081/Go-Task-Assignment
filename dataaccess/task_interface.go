package dataaccess

import (
	"context"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/model"
)

type ITaskDataAccess interface {
	ListTasks(ctx context.Context) []model.Task
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError)
	PutTask(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError)
	DeleteTask(ctx context.Context, id int) common.ICodeError
}
