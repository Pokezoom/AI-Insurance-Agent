# 保险保单AI分析系统 - 前端

基于 Vue 3 + Element Plus 的现代化前端应用

## 功能特性

- 用户登录/注册
- 保单图片上传与AI分析
- 分析历史记录查看
- 个人信息管理
- 管理员用户管理

## 技术栈

- Vue 3
- Vite
- Vue Router
- Pinia (状态管理)
- Element Plus (UI组件库)
- Axios (HTTP请求)

## 快速开始

### 1. 安装依赖

```bash
cd frontend
npm install
```

### 2. 启动开发服务器

```bash
npm run dev
```

前端将在 `http://localhost:3000` 启动

### 3. 构建生产版本

```bash
npm run build
```

## 项目结构

```
frontend/
├── src/
│   ├── api/              # API 接口
│   │   ├── auth.js       # 认证接口
│   │   ├── policy.js     # 保单接口
│   │   └── admin.js      # 管理员接口
│   ├── components/       # 公共组件
│   ├── router/           # 路由配置
│   ├── store/            # 状态管理
│   │   └── user.js       # 用户状态
│   ├── utils/            # 工具函数
│   │   └── request.js    # Axios 封装
│   ├── views/            # 页面组件
│   │   ├── Login.vue     # 登录/注册
│   │   ├── Layout.vue    # 布局
│   │   ├── Analyze.vue   # 保单分析
│   │   ├── Records.vue   # 历史记录
│   │   ├── Profile.vue   # 个人中心
│   │   └── AdminUsers.vue # 用户管理
│   ├── App.vue
│   └── main.js
├── index.html
├── vite.config.js
└── package.json
```

## 页面说明

### 登录/注册页面
- 支持用户登录和注册
- JWT Token 认证

### 保单分析页面
- 拖拽或点击上传保单图片
- 实时预览上传的图片
- 调用后端 AI 分析接口
- 展示分析结果

### 历史记录页面
- 分页展示所有分析记录
- 查看详细分析结果
- 删除历史记录

### 个人中心
- 查看个人信息
- 修改密码

### 用户管理（管理员）
- 查看所有用户
- 启用/禁用用户

## 配置说明

### API 代理配置

在 `vite.config.js` 中配置了代理，将 `/api` 请求转发到后端服务：

```javascript
server: {
  proxy: {
    '/api': {
      target: 'http://localhost:8080',
      changeOrigin: true
    }
  }
}
```

### 环境变量

可以创建 `.env` 文件配置环境变量：

```
VITE_API_BASE_URL=http://localhost:8080
```

## 默认账号

- 管理员：`admin` / `admin123`
