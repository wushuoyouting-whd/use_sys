package controller

import (
	"log"
	"user_be_go/internal/utils"
	"user_be_go/pkg/response"

	"github.com/gin-gonic/gin"
)

// handleError 统一错误处理
func handleError(ctx *gin.Context, err error) {
	log.Printf("错误: %v", err)

	if appErr, ok := err.(*utils.AppError); ok {
		response.Error(ctx, appErr.Status, appErr.Message)
		return
	}

	// 默认错误
	response.Error(ctx, 500, "服务器内部错误: "+err.Error())
}

