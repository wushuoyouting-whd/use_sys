package routes

import (
	"user_be_go/internal/constants"
	"user_be_go/internal/controller"

	"github.com/gin-gonic/gin"
)

// RegisterSysRoutes 注册系统相关路由
func RegisterSysRoutes(router *gin.RouterGroup, sysController *controller.SysController) {
	sys := router.Group(constants.SysPrefix)
	{
		sys.GET("/be_type", sysController.GetBeType)
	}
}

