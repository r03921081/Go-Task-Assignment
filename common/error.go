package common

type ICodeError interface {
	ErrorCode() int
	ErrorMessage() string
}

type codeErrorImpl struct {
	Code    int
	Message string
}

func NewCodeError(code int, message string) ICodeError {
	return &codeErrorImpl{
		Code:    code,
		Message: message,
	}
}

func (c *codeErrorImpl) ErrorCode() int {
	return c.Code
}

func (c *codeErrorImpl) ErrorMessage() string {
	return c.Message
}
