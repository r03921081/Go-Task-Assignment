package constant

const (
	// General errors
	ErrInvalidAuthKey = "invalid auth key"
	ErrIDNotMatch     = "id not match"

	// Auth errors
	ErrWrongPassword = "password is incorrect"

	// User errors
	ErrUserNotFound      = "user not found"
	ErrUserAlreadyExists = "user already exists"

	// Task errors
	ErrTaskNotFound = "task not found"
)

const (
	// General errors
	ErrCodeBadRequest   = 100400
	ErrCodeUnauthorized = 100401

	// Auth errors
	ErrCodeInvalidAuthKey = 101100

	// User errors
	ErrCodeUserNotFound      = 101200
	ErrCodeUserAlreadyExists = 101201
	ErrCodeWrongPassword     = 101202

	// Task errors
	ErrCodeTaskNotFound = 101300
)
