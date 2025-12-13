package response

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ApiResponse 统一响应格式
type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"traceId"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}, message ...string) {
	msg := "success"
	if len(message) > 0 && message[0] != "" {
		msg = message[0]
	}
	c.JSON(200, ApiResponse{
		Code:    0,
		Message: msg,
		Data:    data,
		TraceID: uuid.New().String(),
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(code, ApiResponse{
		Code:    code,
		Message: message,
		Data:    nil,
		TraceID: uuid.New().String(),
	})
}

// Fail 失败响应（用于构造响应对象，不直接发送）
func Fail(code int, message string) ApiResponse {
	return ApiResponse{
		Code:    code,
		Message: message,
		Data:    nil,
		TraceID: uuid.New().String(),
	}
}

