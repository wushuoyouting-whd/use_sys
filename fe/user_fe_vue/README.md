# 用户管理系统 - Vue 前端

基于 Vue 3 + Element Plus 的用户管理系统前端项目。

## 技术栈

- **Vue 3** - 渐进式 JavaScript 框架
- **Vite** - 下一代前端构建工具
- **Element Plus** - Vue 3 组件库
- **Axios** - HTTP 客户端

## 功能特性

- ✅ 用户列表展示（分页）
- ✅ 用户搜索筛选（姓名、邮箱、年龄）
- ✅ 新增用户
- ✅ 编辑用户
- ✅ 删除用户
- ✅ 清新淡雅的界面风格

## 安装运行

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 构建生产版本
npm run build
```

## 项目结构

```
 src/
    ├── api/           # API 接口
    │   └── user.js    # 用户相关接口
    ├── components/    # 组件
    │   ├── UserList.vue   # 用户列表
    │   └── UserForm.vue   # 用户表单
    ├── styles/            # 样式
    ├── App.vue        # 根组件
    └── main.js        # 入口文件
 ├──── package.json    # 项目依赖配置
 ├──── vite.config.js  # Vite 配置（包含 API 代理）
 ├──── README.md       # 项目说明
 ├──── index.html      # HTML 模板
```

## 接口说明

所有接口通过 `/api` 代理到后端服务器（默认 `http://localhost:3000`）

- `GET /api/users` - 获取用户列表
- `GET /api/users/:id` - 获取用户详情
- `POST /api/users` - 创建用户
- `PUT /api/users/:id` - 更新用户
- `DELETE /api/users/:id` - 删除用户

## 界面风格

采用清新淡雅的配色方案：
- 主色调：淡蓝色 (#6b9bd1)
- 背景：浅灰色渐变
- 卡片：白色背景，圆角设计
- 整体风格：简洁、现代、舒适



