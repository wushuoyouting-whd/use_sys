# 路由模块说明

## 模块化路由设计

路由模块采用模块化设计，每个功能模块都有独立的路由文件。

## 添加新模块路由

### 1. 创建路由文件

在 `internal/routes/` 目录下创建新的路由文件，例如 `product_routes.go`:

```go
package routes

import (
    "user_be_go/internal/controller"
    "user_be_go/internal/constants"
    
    "github.com/gin-gonic/gin"
)

// RegisterProductRoutes 注册产品相关路由
func RegisterProductRoutes(router *gin.RouterGroup, productController *controller.ProductController) {
    products := router.Group("/products")
    {
        products.GET("", productController.List)
        products.GET("/:id", productController.GetByID)
        products.POST("", productController.Create)
        products.PUT("/:id", productController.Update)
        products.DELETE("/:id", productController.Delete)
    }
}
```

### 2. 在 Container 中添加依赖

在 `internal/container/container.go` 中添加新的依赖：

```go
type Container struct {
    // ... 现有字段
    
    // 新产品模块
    ProductRepo      *repository.ProductRepository
    ProductService   *service.ProductService
    ProductController *controller.ProductController
}

// 在 initRepositories 中添加
func (c *Container) initRepositories() {
    // ... 现有代码
    c.ProductRepo = repository.NewProductRepository(c.DB)
}

// 在 initServices 中添加
func (c *Container) initServices() {
    // ... 现有代码
    c.ProductService = service.NewProductService(c.ProductRepo)
}

// 在 initControllers 中添加
func (c *Container) initControllers() {
    // ... 现有代码
    c.ProductController = controller.NewProductController(c.ProductService)
}
```

### 3. 在 App 中注册路由

在 `internal/app/app.go` 的 `registerRoutes` 函数中添加：

```go
func registerRoutes(engine *gin.Engine, ctn *container.Container) {
    api := engine.Group("/api")
    
    // 现有路由
    routes.RegisterUserRoutes(api, ctn.UserController)
    routes.RegisterSysRoutes(api, ctn.SysController)
    
    // 新添加的路由
    routes.RegisterProductRoutes(api, ctn.ProductController)
}
```

## 路由组织原则

1. **一个模块一个文件** - 每个功能模块有独立的路由文件
2. **统一前缀** - 所有 API 路由使用 `/api` 前缀
3. **RESTful 设计** - 遵循 RESTful API 设计规范
4. **清晰命名** - 路由函数使用 `RegisterXXXRoutes` 命名

