package dataaccess

import (
	"context"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

var AuthStorage IAuthDataAccess

type authInMemory struct {
	c cache.Cache
}

func NewAuthStorage(ctx context.Context) *authInMemory {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	return &authInMemory{
		c: *c,
	}
}

func (a *authInMemory) CreateAuthKey(ctx context.Context, user *model.User) (string, common.ICodeError) {
	key := uuid.New().String()
	a.c.Set(key, user, constant.DefaultAuthDuration)
	return key, nil
}

func (a *authInMemory) ValidateAuthKey(ctx context.Context, key string) (*model.User, common.ICodeError) {
	if v, found := a.c.Get(key); found {
		return v.(*model.User), nil
	}
	return nil, common.NewCodeError(constant.ErrCodeInvalidAuthKey, constant.ErrInvalidAuthKey)
}
