package module

import (
	"context"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/model"
)

type ITaskModule interface {
	ListTasks(ctx context.Context) []model.Task
	CreateTask(ctx context.Context, name string) (*model.Task, common.ICodeError)
	PutTask(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError)
	DeleteTask(ctx context.Context, id int) common.ICodeError
}

var TaskModule ITaskModule = &taskModuleImpl{}

type taskModuleImpl struct{}

func (t *taskModuleImpl) ListTasks(ctx context.Context) []model.Task {
	return ListTasks(ctx)
}
func (t *taskModuleImpl) CreateTask(ctx context.Context, name string) (*model.Task, common.ICodeError) {
	task := model.NewTask(name)
	return CreateTask(ctx, task)
}
func (t *taskModuleImpl) PutTask(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError) {
	return PutTask(ctx, task)
}
func (t *taskModuleImpl) DeleteTask(ctx context.Context, id int) common.ICodeError {
	return DeleteTask(ctx, id)
}
