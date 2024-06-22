package dataaccess

import (
	"context"
	"fmt"
	"sort"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/patrickmn/go-cache"
)

var TaskStorage ITaskDataAccess

type taskInMemory struct {
	c cache.Cache
}

func NewTaskStorage(ctx context.Context) *taskInMemory {
	c := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	return &taskInMemory{
		c: *c,
	}
}

func (t *taskInMemory) ListTasks(ctx context.Context) []model.Task {
	tasks := []model.Task{}
	for _, v := range t.c.Items() {
		task := v.Object.(*model.Task)
		tasks = append(tasks, *task)
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks
}

func (t *taskInMemory) CreateTask(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError) {
	t.c.Set(fmt.Sprint(task.ID), task, cache.DefaultExpiration)
	return task, nil
}

func (t *taskInMemory) PutTask(ctx context.Context, task *model.Task) (*model.Task, common.ICodeError) {
	if _, found := t.c.Get(fmt.Sprint(task.ID)); !found {
		return nil, common.NewCodeError(constant.ErrCodeTaskNotFound, constant.ErrTaskNotFound)
	}
	t.c.Set(fmt.Sprint(task.ID), task, cache.DefaultExpiration)
	return task, nil
}

func (t *taskInMemory) DeleteTask(ctx context.Context, id int) common.ICodeError {
	t.c.Delete(fmt.Sprint(id))
	return nil
}
