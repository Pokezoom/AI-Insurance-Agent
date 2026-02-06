# AI Insurance Agent

保险保单 AI 分析系统

## 功能特性

- 用户注册/登录（JWT认证）
- 保单图片上传与AI分析
- 分析历史记录存储与查询
- 用户管理（管理员功能）

## 快速开始

### 1. 数据库初始化

```bash
mysql -u root -p < init.sql
```

默认管理员账号：
- 用户名: `admin`
- 密码: `admin123`

### 2. 修改配置

编辑 `config/app.yaml`，修改数据库连接信息：

```yaml
database:
  dsn: "root:你的密码@tcp(127.0.0.1:3306)/insurance_agent?charset=utf8mb4&parseTime=True&loc=Local"
```

### 3. 启动服务

```bash
go run cmd/server/main.go
```

服务将在 `http://localhost:8080` 启动

## API 接口

### 认证接口

#### 注册
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"agent001","password":"password123","email":"agent@example.com"}'
```

#### 登录
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"agent001","password":"password123"}'
```

### 保单分析接口

#### 分析保单
```bash
curl -X POST http://localhost:8080/api/policy/analyze \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"image_base64":"base64编码","image_type":"image/png"}'
```

#### 查询历史记录
```bash
curl -X GET "http://localhost:8080/api/policy/records?page=1&page_size=20" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 获取单条记录
```bash
curl -X GET http://localhost:8080/api/policy/records/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 删除记录
```bash
curl -X DELETE http://localhost:8080/api/policy/records/1 \
  -H "Authorization: Bearer YOUR_TOKEN"
```

### 用户接口

#### 获取个人信息
```bash
curl -X GET http://localhost:8080/api/user/profile \
  -H "Authorization: Bearer YOUR_TOKEN"
```

#### 修改密码
```bash
curl -X PUT http://localhost:8080/api/user/password \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"old_password":"password123","new_password":"newpassword456"}'
```

### 管理员接口

#### 获取用户列表
```bash
curl -X GET "http://localhost:8080/api/admin/users?page=1&page_size=20" \
  -H "Authorization: Bearer ADMIN_TOKEN"
```

#### 禁用/启用用户
```bash
curl -X PUT http://localhost:8080/api/admin/users/2/status \
  -H "Authorization: Bearer ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"status":"inactive"}'
```

## 项目结构

```
AI-Insurance-Agent/
├── cmd/server/          # 主程序入口
├── config/              # 配置文件
├── internal/
│   ├── client/          # GLM API 客户端
│   ├── handler/         # HTTP 处理器
│   ├── middleware/      # 中间件（JWT认证）
│   ├── model/           # 数据模型
│   ├── repository/      # 数据访问层
│   └── service/         # 业务逻辑层
├── pkg/utils/           # 工具函数
├── prompt/              # Prompt 配置
└── init.sql             # 数据库初始化脚本
```

## 响应码说明

- `0` - 成功
- `1001` - 参数错误
- `1002` - 用户不存在
- `1003` - 密码错误
- `1004` - token无效或过期
- `1005` - 权限不足
- `2001` - 分析失败
- `5000` - 服务器内部错误
