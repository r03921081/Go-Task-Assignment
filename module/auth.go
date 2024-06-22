package module

import (
	"context"

	"github.com/andyliao/task-homework/common"
	"github.com/andyliao/task-homework/model"
)

type IAuthModule interface {
	CreateAuthKey(ctx context.Context, user *model.User) (string, common.ICodeError)
	IsAuthorized(ctx context.Context, key string) (*model.User, common.ICodeError)
}

var AuthModule IAuthModule = &authModuleImpl{}

type authModuleImpl struct{}

func (a *authModuleImpl) CreateAuthKey(ctx context.Context, user *model.User) (string, common.ICodeError) {
	return CreateAuthKey(ctx, user)
}

func (a *authModuleImpl) IsAuthorized(ctx context.Context, key string) (*model.User, common.ICodeError) {
	user, err := ValidateAuthKey(ctx, key)
	if err != nil {
		return nil, err
	}
	return user, nil
}
