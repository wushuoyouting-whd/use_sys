package controller

import (
	"user_be_go/internal/config"
	"user_be_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type SysController struct{}

func NewSysController() *SysController {
	return &SysController{}
}

// GetBeType 获取后端类型
// @Summary 获取后端类型
// @Description 获取当前后端系统类型
// @Tags Sys
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Router /api/sys/be_type [get]
func (c *SysController) GetBeType(ctx *gin.Context) {
	response.Success(ctx, map[string]string{
		"beType": config.AppConfig.BeType,
	})
}

