package dto

// CreateUserRequest example
type CreateUserRequest struct {
	Name     string `json:"name" example:"user name"`
	Password string `json:"password" example:"password"`
}
