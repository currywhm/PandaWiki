# PandaWiki Docker 部署指南

本指南介绍如何使用 Docker Compose 部署 PandaWiki。

## 前置要求

- Docker 20.x 或更高版本
- Docker Compose v2.x 或更高版本
- 至少 4GB 可用内存
- 至少 10GB 可用磁盘空间

## 快速开始

### 1. 配置环境变量

复制环境变量示例文件并修改：

```bash
cp .env.example .env
```

编辑 `.env` 文件，修改以下关键配置：

```bash
# 数据库密码（建议修改）
POSTGRES_PASSWORD=your-secure-password

# JWT 密钥（必须修改为随机字符串）
JWT_SECRET=your-random-jwt-secret

# MinIO 密钥（建议修改）
S3_SECRET_KEY=your-minio-secret-key

# 管理员初始密码
ADMIN_PASSWORD=your-admin-password
```

### 2. 启动服务

```bash
# 构建并启动所有服务
docker compose up -d --build

# 查看服务状态
docker compose ps

# 查看日志
docker compose logs -f
```

### 3. 访问服务

- **管理后台**: http://localhost:2443
- **用户端**: http://localhost:3010
- **API 服务**: http://localhost:8000

默认管理员账号：
- 用户名: `admin`
- 密码: 你在 `.env` 中设置的 `ADMIN_PASSWORD`

## 服务说明

| 服务 | 端口 | 说明 |
|------|------|------|
| panda-wiki-admin | 2443 | 管理后台前端 |
| panda-wiki-app | 3010 | 用户端前端 |
| panda-wiki-api | 8000 | 后端 API 服务 |
| panda-wiki-consumer | - | 后端消费者服务 |
| panda-wiki-postgres | 5432 | PostgreSQL 数据库 |
| panda-wiki-redis | 6379 | Redis 缓存 |
| panda-wiki-nats | 4222 | NATS 消息队列 |
| panda-wiki-minio | 9000/9001 | MinIO 对象存储 |

## 常用命令

```bash
# 停止所有服务
docker compose down

# 停止并删除数据卷（警告：会删除所有数据）
docker compose down -v

# 重启单个服务
docker compose restart panda-wiki-api

# 查看单个服务日志
docker compose logs -f panda-wiki-api

# 重新构建并启动
docker compose up -d --build
```

## 配置 AI 模型

PandaWiki 需要配置 AI 模型才能使用 AI 相关功能。登录管理后台后，按照提示配置 AI 模型。

推荐使用 [百智云模型广场](https://baizhi.cloud/) 快速接入 AI 模型。

## 生产环境部署建议

1. **使用反向代理**: 建议在前面部署 Nginx 或 Caddy 作为反向代理，配置 SSL 证书
2. **修改默认密码**: 务必修改所有默认密码
3. **数据备份**: 定期备份 PostgreSQL 数据和 MinIO 存储
4. **资源限制**: 根据实际情况配置容器资源限制

## 故障排除

### 服务启动失败

```bash
# 查看详细日志
docker compose logs panda-wiki-api

# 检查容器状态
docker compose ps -a
```

### 数据库连接失败

确保 PostgreSQL 服务已启动并健康：

```bash
docker compose exec panda-wiki-postgres pg_isready -U panda-wiki
```

### 前端无法访问 API

检查 nginx 配置和网络连接：

```bash
docker compose exec panda-wiki-admin curl http://panda-wiki-api:8000/health
```

## 更多帮助

- [官方文档](https://pandawiki.docs.baizhi.cloud/)
- [GitHub Issues](https://github.com/chaitin/PandaWiki/issues)
