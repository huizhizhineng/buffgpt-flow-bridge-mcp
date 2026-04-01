#### 存量api转mcp
#### 基于gin的api脚手架

---

## BuffGPT Flow Bridge MCP

在 **Gin** 上搭建的网关与管理服务：将 **存量 OpenAPI** 描述的能力接入 **MCP（Model Context Protocol）**，提供 MCP Streamable HTTP 入口及一套用于上传 OpenAPI、管理 MCP Server / Tools 的 REST API。

- **仓库**：<https://github.com/huizhizhineng/buffgpt-flow-bridge-mcp>
- **许可**：MIT，见 [LICENSE](./LICENSE)
- **接口与协议说明**：[docs/protocol.md](./docs/protocol.md)

## 功能概览

- 上传 OpenAPI 文档，构建或更新 MCP Server 与工具定义  
- MCP 网关：`POST /gateway/mcp`、`POST /gateway/:serverToken/mcp`  
- 基于 UUID / 表单的 MCP Server、Tools 的增删改查与连通性测试  
- 支持 PostgreSQL / MySQL、Redis、可选 Nacos 注册（见配置）

## 环境要求

- Go **1.24+**（以 `go.mod` 为准）  
- 数据库、Redis 等按 `configs/config.yaml` 或环境变量准备  

## 快速开始

```bash
# 安装依赖
go mod download

# 本地运行（默认读取 ./configs）
go run ./cmd/main.go ./cmd/wire_gen.go -conf ./configs
```

或使用 Makefile / Docker（见项目内 `Makefile`、`Dockerfile`、`docker-compose.yml`）。

## 配置说明

- 配置文件目录：启动参数 `-conf`，例如 `-conf ./configs`  
- 主要项：`server.http`（端口、超时）、`data.database`、`data.redis`、`registry.nacos`  
- **生产环境**请通过环境变量覆盖敏感项，勿将真实口令提交到仓库  

详细 HTTP 路径与 MCP 网关请求头约定见 **[docs/protocol.md](./docs/protocol.md)**。

## 项目结构（简要）

| 路径 | 说明 |
|------|------|
| `cmd/` | 程序入口、Wire 注入 |
| `router/` | Gin 路由与中间件挂载 |
| `internal/service/` | HTTP 处理器 |
| `internal/biz/`、`internal/data/` | 业务与数据访问 |
| `internal/mcp/` | MCP 服务、代理与 OpenAPI 转换 |
| `configs/` | 默认配置示例 |

## 文档

- [docs/protocol.md](./docs/protocol.md) — HTTP / MCP 网关协议与路由说明  
- [docs/README.md](./docs/README.md) — 文档目录说明  

欢迎通过 Issue / PR 反馈问题与改进建议。
