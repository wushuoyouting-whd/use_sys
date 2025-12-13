# 用户管理系统 - Go 后端

基于 Go 语言开发的企业级用户管理系统，使用 Gin 框架和 GORM ORM。

## 技术栈

- **框架**: Gin 1.9.1
- **ORM**: GORM 1.25.5
- **数据库**: PostgreSQL
- **文档**: Swagger/OpenAPI
- **Go 版本**: 1.21+

## 项目结构

```
user_be_go/
├── cmd/
│   └── server/
│       └── main.go          # 应用入口
├── internal/
│   ├── app/              # 应用初始化
│   │   └── app.go        # 应用结构体和启动逻辑
│   ├── config/           # 配置管理
│   │   ├── config.go     # 配置加载
│   │   └── database.go   # 数据库连接
│   ├── container/        # 依赖注入容器
│   │   └── container.go  # 依赖管理
│   ├── constants/        # 常量定义
│   │   └── constants.go  # 统一常量
│   ├── routes/           # 路由模块
│   │   ├── user_routes.go    # 用户路由
│   │   ├── sys_routes.go     # 系统路由
│   │   └── README.md         # 路由模块说明
│   ├── domain/           # 实体模型
│   │   └── user.go       # 用户实体
│   ├── dto/              # 数据传输对象
│   │   ├── user_dto.go   # 用户 DTO
│   │   └── date.go       # 自定义日期类型
│   ├── repository/       # 数据访问层
│   │   └── user_repository.go
│   ├── service/          # 业务逻辑层
│   │   └── user_service.go
│   ├── controller/       # 控制器层
│   │   ├── user_controller.go
│   │   ├── sys_controller.go
│   │   └── error_handler.go
│   ├── middleware/       # 中间件
│   │   └── error_handler.go
│   └── utils/           # 工具类
│       └── app_error.go
├── pkg/                 # 公共包
│   └── response/        # 响应格式
│       └── api_response.go
├── docs/               # Swagger 文档（自动生成）
├── .env                # 环境变量（不提交）
├── .env.example        # 环境变量示例
├── go.mod              # Go 模块定义
├── Makefile            # 构建脚本
└── README.md           # 项目说明
```

## 快速开始

### 1. 安装依赖

```bash
go mod download
go mod tidy
```

### 2. 安装 Swagger 工具（首次使用）

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 3. 配置环境变量

复制 `.env.example` 为 `.env` 并修改配置：

```bash
cp .env.example .env
```

编辑 `.env` 文件，配置数据库连接信息。

### 4. 生成 Swagger 文档

```bash
# 切换go.mod指定的go版本
make swagger
# 或
swag init -g cmd/server/main.go -o docs
```

### 5. 运行应用

```bash
make run
# 或
go run cmd/server/main.go
```

### 6. 访问 Swagger 文档

打开浏览器访问：http://localhost:3013/swagger/index.html

## API 接口

### 用户管理

- `GET /api/users` - 获取用户列表（支持分页和条件查询）
  - 查询参数：`page`, `limit`, `name`, `email`, `age`, `beType`, `startDate`, `endDate`
- `GET /api/users/:id` - 获取用户详情
- `POST /api/users` - 创建用户
- `PUT /api/users/:id` - 更新用户
- `DELETE /api/users/:id` - 删除用户

### 系统接口

- `GET /api/sys/be_type` - 获取后端类型

## 响应格式

所有接口统一使用以下响应格式：

```json
{
  "code": 0,
  "message": "success",
  "data": {},
  "traceId": "550e8400-e29b-41d4-a716-446655440000"
}
```

## 开发规范

- 遵循 Go 标准项目布局
- 使用分层架构：Controller -> Service -> Repository
- 统一响应格式和错误处理
- 完整的 Swagger API 文档
- 使用 GORM 进行数据库操作
- **模块化设计** - 路由、依赖注入、常量统一管理
- **易于扩展** - 添加新功能只需创建对应的路由、服务、控制器文件

## 模块化设计

### 1. 路由模块化
- 每个功能模块有独立的路由文件（`internal/routes/`）
- 路由注册统一在 `internal/app/app.go` 中管理

### 2. 依赖注入容器
- 所有依赖统一在 `internal/container/container.go` 中管理
- 自动初始化 Repository -> Service -> Controller 的依赖链

### 3. 常量管理
- 所有常量定义在 `internal/constants/constants.go` 中
- 包括 HTTP 状态码、业务错误码、错误消息等

### 4. 应用初始化
- 应用启动逻辑封装在 `internal/app/app.go` 中
- 统一管理中间件、路由、Swagger 的注册

## 添加新功能模块

参考 `internal/routes/README.md` 了解如何添加新的功能模块。

## Makefile 命令

- `make run` - 运行应用
- `make build` - 构建应用
- `make swagger` - 生成 Swagger 文档
- `make deps` - 下载并整理依赖
- `make clean` - 清理构建文件

## 注意事项

1. 确保 PostgreSQL 数据库已创建并运行
2. 数据库表结构应与 Node.js 版本保持一致
3. 首次运行前需要生成 Swagger 文档
4. 环境变量文件 `.env` 不应提交到版本控制
