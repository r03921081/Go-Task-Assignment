package dataaccess

import (
	"context"
	"testing"
	"time"

	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthInMemory(t *testing.T) {
	// Initialize the auth storage
	ctx := context.Background()
	authStorage := NewAuthStorage(ctx)

	// Test CreateAuthKey
	user := model.NewUser("test", "test")
	key, err := authStorage.CreateAuthKey(ctx, user)
	assert.NotEmpty(t, key, "Auth key should not be empty")
	assert.Nil(t, err, "Error should be nil")

	// Check if the user is stored correctly
	storedUser, err := authStorage.ValidateAuthKey(ctx, key)
	assert.NotNil(t, storedUser, "Stored user should not be nil")
	assert.Equal(t, user.ID, storedUser.ID, "Stored user ID should match")
	assert.Equal(t, user.Username, storedUser.Username, "Stored user name should match")
	assert.Nil(t, err, "Error should be nil")

	// Test ValidateAuthKey with an invalid key
	invalidKey := uuid.New().String()
	invalidUser, err := authStorage.ValidateAuthKey(ctx, invalidKey)
	assert.Nil(t, invalidUser, "User should be nil for an invalid key")
	assert.NotNil(t, err, "Error should not be nil")
	assert.Equal(t, err.ErrorCode(), constant.ErrCodeInvalidAuthKey, constant.ErrInvalidAuthKey)

	// Test expiration
	authStorage.c.Set(key, user, 1*time.Second)
	time.Sleep(2 * time.Second)
	expiredUser, err := authStorage.ValidateAuthKey(ctx, key)
	assert.Nil(t, expiredUser, "User should be nil after expiration")
	assert.NotNil(t, err, "Error should not be nil")
	assert.Equal(t, err.ErrorCode(), constant.ErrCodeInvalidAuthKey, constant.ErrCodeInvalidAuthKey)
}
