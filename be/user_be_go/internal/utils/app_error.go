package utils

import (
	"fmt"
	"user_be_go/internal/constants"
)

// AppError 应用错误
type AppError struct {
	Code    int
	Message string
	Status  int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// NewAppError 创建应用错误
func NewAppError(code int, message string, status ...int) *AppError {
	statusCode := constants.StatusInternalServerError
	if len(status) > 0 {
		statusCode = status[0]
	}
	return &AppError{
		Code:    code,
		Message: message,
		Status:  statusCode,
	}
}

// NewNotFoundError 创建 404 错误
func NewNotFoundError(message string) *AppError {
	return NewAppError(constants.CodeUserNotFound, message, constants.StatusNotFound)
}

// NewConflictError 创建 409 错误
func NewConflictError(message string) *AppError {
	return NewAppError(constants.CodeUserEmailExists, message, constants.StatusConflict)
}

// NewBadRequestError 创建 400 错误
func NewBadRequestError(message string) *AppError {
	return NewAppError(constants.CodeInvalidParams, message, constants.StatusBadRequest)
}

