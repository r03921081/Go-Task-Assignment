package module

import (
	"context"
	"testing"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/model"
	"github.com/stretchr/testify/assert"
)

func TestTaskModule_ListTasks(t *testing.T) {
	mockListTasks := func(ctx context.Context) []model.Task {
		return []model.Task{{ID: 1, Name: "task1"}, {ID: 2, Name: "task2"}}
	}
	ListTasks = mockListTasks

	taskModule := &taskModuleImpl{}

	ctx := context.Background()
	tasks := taskModule.ListTasks(ctx)

	expectedTasks := []model.Task{{ID: 1, Name: "task1"}, {ID: 2, Name: "task2"}}
	assert.Equal(t, expectedTasks, tasks)
}

func TestTaskModule_CreateTask(t *testing.T) {
	mockCreateTask := func(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError) {
		task.ID = 1
		return task, nil
	}
	CreateTask = mockCreateTask

	taskModule := &taskModuleImpl{}

	ctx := context.Background()
	task, err := taskModule.CreateTask(ctx, "new task")

	expectedTask := &model.Task{ID: 1, Name: "new task"}
	assert.Equal(t, expectedTask, task)
	assert.Nil(t, err)
}

func TestTaskModule_PutTask(t *testing.T) {
	mockPutTask := func(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError) {
		return task, nil
	}
	PutTask = mockPutTask

	taskModule := &taskModuleImpl{}

	task := &model.Task{ID: 1, Name: "updated task"}
	ctx := context.Background()
	updatedTask, err := taskModule.PutTask(ctx, task)

	assert.Equal(t, task, updatedTask)
	assert.Equal(t, nil, err)
}

func TestTaskModule_DeleteTask(t *testing.T) {
	mockDeleteTask := func(ctx context.Context, id int) common.ICodeError {
		return nil
	}
	DeleteTask = mockDeleteTask

	taskModule := &taskModuleImpl{}

	ctx := context.Background()
	err := taskModule.DeleteTask(ctx, 1)
	assert.Nil(t, err)
}
