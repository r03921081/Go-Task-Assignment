package dto

import (
	"context"
)

// Response example
type Response struct {
	Code    int         `json:"code,omitempty" example:"0"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty" example:"success"`
}

// NewResponse example
func NewResponse(ctx context.Context, result interface{}) Response {
	return Response{
		Result: result,
	}
}

// NewErrorResponse example
func NewErrorResponse(ctx context.Context, code int, message string) Response {
	return Response{
		Code:    code,
		Message: message,
	}
}
