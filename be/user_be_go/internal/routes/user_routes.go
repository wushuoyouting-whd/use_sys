package routes

import (
	"user_be_go/internal/constants"
	"user_be_go/internal/controller"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes 注册用户相关路由
func RegisterUserRoutes(router *gin.RouterGroup, userController *controller.UserController) {
	users := router.Group(constants.UserPrefix)
	{
		users.GET("", userController.List)
		users.GET("/:id", userController.GetByID)
		users.POST("", userController.Create)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}
}

