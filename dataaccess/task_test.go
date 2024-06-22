package dataaccess

import (
	"context"
	"testing"

	"github.com/andyliao/task-homework/model"
	"github.com/stretchr/testify/assert"
)

func TestTaskInMemory(t *testing.T) {
	// Initialize the task storage
	ctx := context.Background()
	taskStorage := NewTaskStorage(ctx)

	// Test CreateTask
	task1 := &model.Task{ID: 1, Name: "Task 1"}
	createdTask, err := taskStorage.CreateTask(ctx, task1)
	assert.NotNil(t, createdTask, "Created task should not be nil")
	assert.Equal(t, task1.ID, createdTask.ID, "Task ID should match")
	assert.Equal(t, task1.Name, createdTask.Name, "Task name should match")
	assert.Nil(t, err, "Error should be nil")

	// Test ListTasks
	tasks := taskStorage.ListTasks(ctx)
	assert.Len(t, tasks, 1, "Should have one task in the list")
	assert.Equal(t, task1.ID, tasks[0].ID, "Task ID should match")
	assert.Equal(t, task1.Name, tasks[0].Name, "Task name should match")

	// Test PutTask (Update)
	task1Updated := &model.Task{ID: 1, Name: "Updated Task 1"}
	updatedTask, err := taskStorage.PutTask(ctx, task1Updated)
	assert.NotNil(t, updatedTask, "Updated task should not be nil")
	assert.Equal(t, task1Updated.ID, updatedTask.ID, "Updated task ID should match")
	assert.Equal(t, task1Updated.Name, updatedTask.Name, "Updated task name should match")
	assert.Nil(t, err, "Error should be nil")

	// Test if the task is updated in storage
	tasks = taskStorage.ListTasks(ctx)
	assert.Len(t, tasks, 1, "Should have one task in the list after update")
	assert.Equal(t, task1Updated.Name, tasks[0].Name, "Updated task name should match in the list")

	// Test DeleteTask
	err = taskStorage.DeleteTask(ctx, task1Updated.ID)
	tasks = taskStorage.ListTasks(ctx)
	assert.Len(t, tasks, 0, "Task list should be empty after deletion")
	assert.Nil(t, err, "Error should be nil")

	// Test for non-existent task deletion
	err = taskStorage.DeleteTask(ctx, 999)
	tasks = taskStorage.ListTasks(ctx)
	assert.Len(t, tasks, 0, "Task list should still be empty after deleting non-existent task")
	assert.Nil(t, err, "Error should be nil")
}
