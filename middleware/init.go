package middleware

import "github.com/andyliao/task-homework/module"

var (
	IsAuthorized = module.AuthModule.IsAuthorized
)
