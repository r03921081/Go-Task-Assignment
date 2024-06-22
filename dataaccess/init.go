package dataaccess

import "context"

func init() {
	UserStorage = NewUserStorage(context.Background())
	AuthStorage = NewAuthStorage(context.Background())
	TaskStorage = NewTaskStorage(context.Background())
}
