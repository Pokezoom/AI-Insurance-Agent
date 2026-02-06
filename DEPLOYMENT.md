# 保险保单AI分析系统 - 完整部署指南

## 项目概述

这是一个完整的保险保单AI分析系统，包含后端（Go）和前端（Vue 3）。

## 系统架构

```
AI-Insurance-Agent/
├── backend (Go)
│   ├── cmd/server/          # 主程序
│   ├── config/              # 配置
│   ├── internal/            # 内部模块
│   └── pkg/                 # 工具包
└── frontend (Vue 3)
    ├── src/                 # 源代码
    └── public/              # 静态资源
```

## 快速启动

### 1. 数据库准备

```bash
# 登录 MySQL
mysql -u root -p

# 执行初始化脚本
source init.sql
```

默认会创建：
- 数据库：`insurance_agent`
- 管理员账号：`admin` / `admin123`

### 2. 后端启动

```bash
# 安装依赖
go mod tidy

# 启动服务
go run cmd/server/main.go
```

后端将在 `http://localhost:8080` 启动

### 3. 前端启动

```bash
# 进入前端目录
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

前端将在 `http://localhost:3000` 启动

### 4. 访问系统

打开浏览器访问：`http://localhost:3000`

使用默认管理员账号登录：
- 用户名：`admin`
- 密码：`admin123`

## 功能说明

### 用户功能
1. **登录/注册**：支持新用户注册和登录
2. **保单分析**：上传保单图片，AI自动分析
3. **历史记录**：查看所有分析历史，支持分页
4. **个人中心**：查看个人信息，修改密码

### 管理员功能
1. **用户管理**：查看所有用户，启用/禁用用户

## 配置说明

### 后端配置

编辑 `config/app.yaml`：

```yaml
server:
  port: 8080

database:
  dsn: "root:@tcp(127.0.0.1:3306)/insurance_agent?charset=utf8mb4&parseTime=True&loc=Local"

jwt:
  secret: "your-secret-key-change-in-production"
  expire_hours: 168

glm:
  api_key: "你的GLM API Key"
  url: "https://open.bigmodel.cn/api/paas/v4/chat/completions"
```

### 前端配置

编辑 `frontend/vite.config.js` 中的代理配置：

```javascript
server: {
  port: 3000,
  proxy: {
    '/api': {
      target: 'http://localhost:8080',  // 后端地址
      changeOrigin: true
    }
  }
}
```

## 生产部署

### 后端部署

```bash
# 编译
go build -o insurance-agent cmd/server/main.go

# 运行
./insurance-agent
```

### 前端部署

```bash
cd frontend

# 构建
npm run build

# 将 dist 目录部署到 Nginx 或其他 Web 服务器
```

Nginx 配置示例：

```nginx
server {
    listen 80;
    server_name your-domain.com;

    root /path/to/frontend/dist;
    index index.html;

    location / {
        try_files $uri $uri/ /index.html;
    }

    location /api {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

## API 接口文档

详见 `README.md` 中的接口说明。

## 技术栈

### 后端
- Go 1.25
- Gin (Web框架)
- GORM (ORM)
- MySQL 8.0
- JWT 认证
- GLM-4V (AI模型)

### 前端
- Vue 3
- Vite
- Element Plus
- Vue Router
- Pinia
- Axios

## 常见问题

### 1. 数据库连接失败
检查 `config/app.yaml` 中的数据库配置是否正确。

### 2. 前端无法访问后端
检查前端 `vite.config.js` 中的代理配置，确保后端服务已启动。

### 3. 图片上传失败
检查图片大小是否超过限制（建议小于 10MB）。

### 4. AI 分析失败
检查 GLM API Key 是否正确配置。

## 开发建议

1. 开发时使用 `npm run dev` 启动前端，支持热更新
2. 修改后端代码后需要重启服务
3. 生产环境建议使用 HTTPS
4. 定期备份数据库
5. 修改默认的 JWT Secret

## 联系方式

如有问题，请提交 Issue。
