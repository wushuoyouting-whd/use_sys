# User Management System - Java Backend

基于 Spring Boot + MyBatis 的企业级用户管理系统后端（Java 版本）


## 技术栈

- **Spring Boot 3.2.0**
- **MyBatis 3.0.3**
- **PostgreSQL**
- **SpringDoc OpenAPI (Swagger)**
- **Lombok**
- **Java 21**

## 项目结构

```
src/main/java/com/user/
├── UserApplication.java              # 主启动类
├── config/                           # 配置类
│   ├── MyBatisConfig.java              # MyBatis 配置
│   └── WebConfig.java                  # Web 配置（跨域等）
├── controller/                       # 控制器层
│   ├── UserController.java             # 用户控制器
│   └── SysController.java              # 系统控制器
├── service/                          # 服务层
│   └── UserService.java                # 用户服务
├── mapper/                           # MyBatis Mapper 接口
│   └── UserMapper.java                 # 用户 Mapper
├── domain/                           # 实体类（领域模型）
│   └── User.java                       # 用户实体
├── dto/                             # 数据传输对象
│   ├── UserCreateDTO.java              # 创建用户 DTO
│   ├── UserUpdateDTO.java              # 更新用户 DTO
│   ├── UserQueryDTO.java               # 查询条件 DTO
│   └── PageResponse.java               # 分页响应 DTO
├── common/                          # 公共类
│   ├── exception/                   # 异常处理
│   │   ├── AppException.java           # 自定义异常
│   │   └── GlobalExceptionHandler.java # 全局异常处理器
│   └── response/                    # 响应类
│       └── ApiResponse.java            # 统一响应格式
└── utils/                           # 工具类（预留）

src/main/resources/
├── application.yml                  # 应用配置
└── mapper/                          # MyBatis Mapper XML
    └── UserMapper.xml               # 用户 Mapper XML
```

## 环境要求

请注意JDK 和 Maven之间的版本的兼容性
- JDK 21
- Maven 3.9.11
- PostgreSQL 12+

## 配置说明

### 1. 数据库配置

在 `src/main/resources/application.yml` 中配置数据库连接：

```yaml
# 修改为自己的连接配置
spring:
  datasource:
    url: jdbc:postgresql://localhost:5432/user_sys
    username: owen
    password: 123456
```

### 2. MyBatis 配置

```yaml
spring:
  mybatis:
    mapper-locations: classpath:mapper/*.xml
    type-aliases-package: com.user.domain
    configuration:
      map-underscore-to-camel-case: true
```

### 3. 应用配置

```yaml
server:
  port: 3012      # 服务器端口

app:
  be-type: Java   # 后端类型标识
```

## 运行项目

### 1. 编译项目

```bash
mvn clean compile
```

### 2. 运行项目

```bash
mvn spring-boot:run
```

或

```bash
java -jar target/user-be-java-1.0.0.jar
```

### 3. 访问 Swagger UI

启动后访问：http://localhost:3012/swagger-ui.html

## API 接口

### 用户管理接口

- `GET /api/users`          - 获取用户列表（支持分页和条件查询）
- `GET /api/users/{id}`     - 获取用户详情
- `POST /api/users`         - 创建用户
- `PUT /api/users/{id}`     - 更新用户
- `DELETE /api/users/{id}`  - 删除用户

### 系统接口

- `GET /api/sys/be_type` - 获取后端类型

## 查询参数

用户列表接口支持以下查询参数：

- `page` - 页码（默认：1）
- `limit` - 每页数量（默认：10）
- `name` - 姓名（模糊搜索）
- `email` - 邮箱（模糊搜索）
- `age` - 年龄（精确匹配）
- `beType` - 后端类型（精确匹配）
- `startDate` - 开始日期（格式：YYYY-MM-DD）
- `endDate` - 结束日期（格式：YYYY-MM-DD）

## 响应格式

所有接口统一使用以下响应格式：

```json
{
  "code": 0,
  "message": "success",
  "data": {},
  "traceId": "uuid"
}
```

## 数据库表结构

项目使用已存在的 `user` 表，表结构：

```sql
CREATE TABLE "user" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(400),
    age INTEGER DEFAULT 0,
    email VARCHAR(100) UNIQUE,
    birth_date DATE,
    be_type VARCHAR(100),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

## 项目特点

1. **企业级项目结构**：按照标准分层架构组织代码
2. **MyBatis ORM**：使用 MyBatis 进行数据库操作，SQL 可控制性强
3. **统一响应格式**：所有接口使用统一的响应结构
4. **全局异常处理**：统一的异常处理机制
5. **参数验证**：使用 Jakarta Validation 进行参数校验
6. **Swagger 文档**：自动生成 API 文档
7. **跨域支持**：已配置跨域访问

## 注意事项

1. 确保 PostgreSQL 数据库已创建并存在 `user` 表
2. 确保数据库用户有足够的权限
3. MyBatis Mapper XML 文件位于 `src/main/resources/mapper/` 目录
4. 邮箱字段具有唯一性约束
