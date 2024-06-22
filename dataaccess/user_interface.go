package dataaccess

import (
	"context"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/model"
)

type IUserDataAccess interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, common.ICodeError)
	GetUser(ctx context.Context, username string) *model.User
}
