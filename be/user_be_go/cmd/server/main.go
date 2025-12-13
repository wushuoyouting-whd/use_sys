package main

import (
	"user_be_go/internal/app"
)

// @title 用户管理系统 API
// @version 1.0
// @description 基于 Go 的用户管理系统后端 API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3013
// @BasePath /api

func main() {
	// 创建并启动应用
	application := app.NewApp()
	application.Run()
}

