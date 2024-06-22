package module

import (
	"context"
	"testing"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/andyliao/task-homework/util"
	"github.com/stretchr/testify/assert"
)

func TestUserModule_CreateUser(t *testing.T) {
	mockCreateUser := func(ctx context.Context, user *model.User) (*model.User, common.ICodeError) {
		user.ID = 1
		return user, nil
	}
	CreateUser = mockCreateUser

	userModule := &userModuleImpl{}

	ctx := context.Background()
	user, err := userModule.CreateUser(ctx, "new_user", "password")

	expectedUser := &model.User{ID: 1, Username: "new_user", Password: util.Hash("password")}
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserModule_ValidateUser(t *testing.T) {
	mockGetUser := func(ctx context.Context, username string) *model.User {
		return &model.User{Username: "test_user", Password: util.Hash("password")}
	}
	GetUser = mockGetUser

	userModule := &userModuleImpl{}

	ctx := context.Background()
	user, err := userModule.ValidateUser(ctx, "test_user", "password")

	expectedUser := &model.User{Username: "test_user", Password: util.Hash("password")}
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserModule_ValidateUser_UserNotFound(t *testing.T) {
	mockGetUser := func(ctx context.Context, username string) *model.User {
		return nil
	}
	GetUser = mockGetUser

	userModule := &userModuleImpl{}

	ctx := context.Background()
	user, err := userModule.ValidateUser(ctx, "nonexistent_user", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrCodeUserNotFound, err.ErrorCode())
	assert.Equal(t, constant.ErrUserNotFound, err.ErrorMessage())
}

func TestUserModule_ValidateUser_WrongPassword(t *testing.T) {
	// Mock the GetUser function to return a user with a different password
	mockGetUser := func(ctx context.Context, username string) *model.User {
		return &model.User{Username: "test_user", Password: util.Hash("wrong_password")}
	}
	GetUser = mockGetUser

	userModule := &userModuleImpl{}

	ctx := context.Background()
	user, err := userModule.ValidateUser(ctx, "test_user", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrCodeWrongPassword, err.ErrorCode())
	assert.Equal(t, constant.ErrWrongPassword, err.ErrorMessage())
}
