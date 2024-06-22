package dataaccess

import (
	"context"
	"fmt"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/patrickmn/go-cache"
)

var UserStorage IUserDataAccess

type userInMemory struct {
	c cache.Cache
}

func NewUserStorage(context.Context) *userInMemory {
	c := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	return &userInMemory{
		c: *c,
	}
}

func (u *userInMemory) CreateUser(ctx context.Context, user *model.User) (*model.User, common.ICodeError) {
	if _, found := u.c.Get(fmt.Sprint(user.Username)); found {
		return nil, common.NewCodeError(constant.ErrCodeUserAlreadyExists, constant.ErrUserAlreadyExists)
	}

	u.c.Set(fmt.Sprint(user.Username), user, cache.DefaultExpiration)
	return user, nil
}

func (u *userInMemory) GetUser(ctx context.Context, username string) *model.User {
	if v, found := u.c.Get(username); found {
		user := v.(*model.User)
		return user
	}
	return nil
}
