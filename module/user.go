package module

import (
	"context"
	"fmt"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/constant"
	"github.com/andyliao/task-homework/model"
	"github.com/andyliao/task-homework/util"
)

type IUserModule interface {
	CreateUser(ctx context.Context, username string, password string) (*model.User, common.ICodeError)
	ValidateUser(ctx context.Context, username string, password string) (*model.User, common.ICodeError)
}

var UserModule IUserModule = &userModuleImpl{}

type userModuleImpl struct{}

func (u *userModuleImpl) CreateUser(ctx context.Context, username string, password string) (*model.User, common.ICodeError) {
	user := model.NewUser(username, password)
	return CreateUser(ctx, user)
}

func (u *userModuleImpl) ValidateUser(ctx context.Context, username string, password string) (*model.User, common.ICodeError) {
	user := GetUser(ctx, username)
	if user == nil {
		return nil, common.NewCodeError(constant.ErrCodeUserNotFound, constant.ErrUserNotFound)
	}
	if user.Password != util.Hash(password) {
		common.Logger.Error(ctx, fmt.Sprintf("wrong password: %s", util.Hash(password)))
		return nil, common.NewCodeError(constant.ErrCodeWrongPassword, constant.ErrWrongPassword)
	}
	return user, nil
}
