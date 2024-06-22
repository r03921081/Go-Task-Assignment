package v1

import "github.com/andyliao/task-homework/module"

var (
	ListTasks  = module.TaskModule.ListTasks
	CreateTask = module.TaskModule.CreateTask
	PutTask    = module.TaskModule.PutTask
	DeleteTask = module.TaskModule.DeleteTask
)

var (
	CreateUser   = module.UserModule.CreateUser
	ValidateUser = module.UserModule.ValidateUser
)

var (
	CreateAuthKey = module.AuthModule.CreateAuthKey
	IsAuthorized  = module.AuthModule.IsAuthorized
)
