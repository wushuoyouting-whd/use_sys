package app

import (
	"log"
	"user_be_go/internal/config"
	"user_be_go/internal/container"
	"user_be_go/internal/middleware"
	"user_be_go/internal/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// App 应用结构
type App struct {
	Engine    *gin.Engine
	Container *container.Container
}

// NewApp 创建应用实例
func NewApp() *App {
	// 创建依赖容器
	ctn := container.NewContainer()

	// 创建 Gin 引擎
	engine := gin.Default()

	// 注册全局中间件
	registerGlobalMiddleware(engine)

	// 注册路由
	registerRoutes(engine, ctn)

	// 注册 Swagger
	registerSwagger(engine)

	return &App{
		Engine:    engine,
		Container: ctn,
	}
}

// registerGlobalMiddleware 注册全局中间件
func registerGlobalMiddleware(engine *gin.Engine) {
	// 错误处理中间件
	engine.Use(middleware.ErrorHandler())
	
	// 可以在这里添加更多全局中间件
	// 例如：日志、CORS、限流等
}

// registerRoutes 注册所有路由
func registerRoutes(engine *gin.Engine, ctn *container.Container) {
	// API 路由组
	api := engine.Group("/api")
	
	// 注册用户路由
	routes.RegisterUserRoutes(api, ctn.UserController)
	
	// 注册系统路由
	routes.RegisterSysRoutes(api, ctn.SysController)
	
	// 后续添加新模块时，只需在这里添加一行即可
	// routes.RegisterXXXRoutes(api, ctn.XXXController)
}

// registerSwagger 注册 Swagger 文档
func registerSwagger(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Run 启动应用
func (a *App) Run() {
	port := ":" + config.AppConfig.ServerPort
	log.Printf("服务器启动在端口 %s", port)
	if err := a.Engine.Run(port); err != nil {
		log.Fatal("服务器启动失败:", err)
	}
}

