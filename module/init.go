package module

import "github.com/andyliao/task-homework/dataaccess"

var (
	CreateAuthKey   = dataaccess.AuthStorage.CreateAuthKey
	ValidateAuthKey = dataaccess.AuthStorage.ValidateAuthKey

	ListTasks  = dataaccess.TaskStorage.ListTasks
	CreateTask = dataaccess.TaskStorage.CreateTask
	PutTask    = dataaccess.TaskStorage.PutTask
	DeleteTask = dataaccess.TaskStorage.DeleteTask

	CreateUser = dataaccess.UserStorage.CreateUser
	GetUser    = dataaccess.UserStorage.GetUser
)
