package model

import (
	"github.com/andyliao/task-homework/util"
)

var userID autoIncr

// User example
type User struct {
	ID       int    `json:"id,omitempty" example:"1"`
	Username string `json:"username,omitempty" example:"user name"`
	Password string `json:"password,omitempty" example:"password"`
}

// NewUser example
func NewUser(username, password string) *User {
	return &User{
		ID:       userID.GetID(),
		Username: username,
		Password: util.Hash(password),
	}
}
