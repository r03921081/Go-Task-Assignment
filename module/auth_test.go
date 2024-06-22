package module

import (
	"context"
	"testing"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/stretchr/testify/assert"
)

func TestAuthModule_CreateAuth(t *testing.T) {
	mockCreateAuthKey := func(ctx context.Context, user *model.User) (string, common.ICodeError) {
		return "test_auth_key", nil
	}
	CreateAuthKey = mockCreateAuthKey

	authModule := &authModuleImpl{}

	user := &model.User{Username: "test_user"}

	ctx := context.Background()
	result, err := authModule.CreateAuthKey(ctx, user)

	assert.Nil(t, err)
	assert.Equal(t, "test_auth_key", result)
}

func TestAuthModule_IsAuthorized(t *testing.T) {
	mockValidateAuthKey := func(ctx context.Context, key string) (*model.User, common.ICodeError) {
		return &model.User{Username: "test_user"}, nil
	}
	ValidateAuthKey = mockValidateAuthKey

	authModule := &authModuleImpl{}

	authKey := "test_auth_key"

	ctx := context.Background()
	user, err := authModule.IsAuthorized(ctx, authKey)

	expectedUser := &model.User{Username: "test_user"}
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestAuthModule_IsAuthorized_InvalidKey(t *testing.T) {
	mockValidateAuthKey := func(ctx context.Context, key string) (*model.User, common.ICodeError) {
		return nil, common.NewCodeError(constant.ErrCodeInvalidAuthKey, constant.ErrInvalidAuthKey)
	}
	ValidateAuthKey = mockValidateAuthKey

	authModule := &authModuleImpl{}

	authKey := "invalid_auth_key"

	ctx := context.Background()
	_, err := authModule.IsAuthorized(ctx, authKey)

	assert.NotNil(t, err)
	assert.Equal(t, constant.ErrCodeInvalidAuthKey, err.ErrorCode())
	assert.Equal(t, constant.ErrInvalidAuthKey, err.ErrorMessage())
}
