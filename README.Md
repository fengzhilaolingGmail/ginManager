> 一套 Go + Gin + GORM 的通用后台管理系统，支持 MySQL / PostgreSQL 双数据库、JWT 鉴权、按钮级权限、操作日志、菜单管理等基础功能；分层清晰，代码简洁，开箱即用。

---

# 🚀 核心特性

| 功能                                       | 状态         |
| ------------------------------------------ | ------------ |
| 用户/用户组/角色/菜单/权限                 | ✅           |
| JWT 登录 & 刷新令牌                        | ✅           |
| 按钮级权限控制（通配符）                   | ✅           |
| 操作日志（自动记录+导出）                  | ✅           |
| 登录日志（可扩展）                         | ✅           |
| Layui 前端适配（树形菜单+分页）            | ✅           |
| MySQL / PostgreSQL 双数据库                | ✅           |
| 分层架构（repository → service → handler） | ✅           |
| Docker 一键部署                            | ❗（见文末） |

---

# 📦 技术栈

- Web 框架：Gin v1.9+
- ORM：GORM v1.25+
- 数据库：MySQL 8.x ｜ PostgreSQL 16+
- 缓存：Redis（可选）
- 鉴权：JWT（HS256）
- 日志：Zap + Lumberjack
- 配置：Viper（yaml）
- 前端：Layui 2.8（模板随包提供）

---

# 📁 项目结构

```
.
├── config/               // 配置解析
├── logger/               // Zap 日志
├── middleware/           // JWT、操作日志、CORS...
├── model/
│   └── entity/          // 数据库实体
├── repository/          // 数据映射层（零业务）
├── service/             // 业务编排
├── handler/             // 接口层（Layui 格式）
├── router/              // 路由注册
├── utils/               // 工具（JWT、密码、分页）
├── docker/              // Dockerfile & compose
└── run.go              // 启动入口
```

---
