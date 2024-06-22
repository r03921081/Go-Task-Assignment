package dataaccess

import (
	"context"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/model"
)

type IAuthDataAccess interface {
	CreateAuthKey(ctx context.Context, user *model.User) (string, common.ICodeError)
	ValidateAuthKey(ctx context.Context, key string) (*model.User, common.ICodeError)
}
