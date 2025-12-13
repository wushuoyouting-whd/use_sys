package container

import (
	"user_be_go/internal/config"
	"user_be_go/internal/controller"
	"user_be_go/internal/repository"
	"user_be_go/internal/service"

	"gorm.io/gorm"
)

// Container 依赖注入容器
type Container struct {
	DB *gorm.DB

	// Repository
	UserRepo *repository.UserRepository

	// Service
	UserService *service.UserService

	// Controller
	UserController *controller.UserController
	SysController  *controller.SysController
}

// NewContainer 创建并初始化容器
func NewContainer() *Container {
	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	config.InitDB()

	container := &Container{
		DB: config.DB,
	}

	// 初始化 Repository
	container.initRepositories()

	// 初始化 Service
	container.initServices()

	// 初始化 Controller
	container.initControllers()

	return container
}

// initRepositories 初始化所有 Repository
func (c *Container) initRepositories() {
	c.UserRepo = repository.NewUserRepository(c.DB)
}

// initServices 初始化所有 Service
func (c *Container) initServices() {
	c.UserService = service.NewUserService(c.UserRepo)
}

// initControllers 初始化所有 Controller
func (c *Container) initControllers() {
	c.UserController = controller.NewUserController(c.UserService)
	c.SysController = controller.NewSysController()
}

