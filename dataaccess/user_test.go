package dataaccess

import (
	"context"
	"testing"

	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/stretchr/testify/assert"
)

func TestUserInMemory(t *testing.T) {
	// Initialize the user storage
	ctx := context.Background()
	userStorage := NewUserStorage(ctx)

	// Test CreateUser
	user1 := &model.User{Username: "john_doe", Password: "John Doe"}
	createdUser, err := userStorage.CreateUser(ctx, user1)
	assert.NotNil(t, createdUser, "Created user should not be nil")
	assert.Equal(t, user1.Username, createdUser.Username, "User username should match")
	assert.Equal(t, user1.Password, createdUser.Password, "User name should match")
	assert.Nil(t, err, "Error should be nil")

	// Test CreateUser with existing username
	user2 := &model.User{Username: "john_doe", Password: "John Doe"}
	_, err = userStorage.CreateUser(ctx, user2)
	assert.NotNil(t, err, "Error should not be nil")
	assert.Equal(t, err.ErrorCode(), constant.ErrCodeUserAlreadyExists, constant.ErrUserAlreadyExists)

	// Test GetUser
	retrievedUser := userStorage.GetUser(ctx, user1.Username)
	assert.NotNil(t, retrievedUser, "Retrieved user should not be nil")
	assert.Equal(t, user1.Username, retrievedUser.Username, "Retrieved user username should match")
	assert.Equal(t, user1.Password, retrievedUser.Password, "Retrieved user name should match")

	// Test GetUser with non-existent username
	nonExistentUser := userStorage.GetUser(ctx, "non_existent")
	assert.Nil(t, nonExistentUser, "User should be nil for non-existent username")
}
