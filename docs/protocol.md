# BuffGPT Flow Bridge MCP — 接口与协议说明

本文档描述本服务对外暴露的 **HTTP 路由**、**MCP 网关** 行为及与鉴权相关的请求头约定。实现以仓库源码为准（见 `router/router.go`）。

## 1. 基础约定

| 项 | 说明 |
|----|------|
| 默认 HTTP 监听 | 由配置 `server.http.addr` / `server.http.port` 决定，示例配置为 `0.0.0.0:9004` |
| API 版本前缀 | 管理类 REST 接口统一挂在 **`/v1`** 下 |
| 内容类型 | 管理接口多为 JSON Body，具体以各 Handler 实现为准 |
| 追踪 | 中间件会处理 `traceId` / `spanId`（见 `middleware`），便于日志关联 |

## 2. MCP 网关（Streamable HTTP）

MCP 会话由网关转发至内部 MCP 服务管理器处理。

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/gateway/mcp` | 默认 MCP 网关入口 |
| POST | `/gateway/:serverToken/mcp` | 按路径参数 `serverToken` 区分 MCP Server 实例 |

### 2.1 可选请求头（鉴权）

网关会把下列头写入请求上下文，供下游鉴权逻辑使用（常量定义见 `pkg/const/const.go`）：

| 请求头 | 常量键名 | 用途 |
|--------|-----------|------|
| `x-mcp-platform-token` | `PlatformToken` | 平台侧令牌 |
| `x-mcp-service-token` | `ServiceToken` | 服务侧令牌 |

路径参数 `serverToken` 对应上下文键 `serverToken`（`ServerPathToken`）。

## 3. 管理 API（`/v1`）

以下为当前路由表中登记的接口；请求/响应字段以具体 `internal/service` 与 `api` 包中的结构体为准。

### 3.1 OpenAPI 与鉴权

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/v1/openapi/upload` | 上传 OpenAPI 文档，用于生成/关联 MCP Server |
| POST | `/v1/openapi/updateForAuth` | 更新 OpenAPI 相关鉴权配置 |

### 3.2 MCP Server

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/v1/mcpServer/getMcpServerInfoByUUID` | 按 UUID 查询 MCP Server 信息 |
| POST | `/v1/mcpServer/updateByUUID` | 按 UUID 更新 MCP Server |
| POST | `/v1/mcpServer/getMcpConnectTokenByUUID` | 获取 MCP 连接令牌 |
| POST | `/v1/mcpServer/deleteMcpServerByUUID` | 删除 MCP Server |
| POST | `/v1/mcpServer/createByForm` | 表单方式创建 MCP Server |
| POST | `/v1/mcpServer/updateMcpServerByForm` | 表单方式更新 MCP Server |

### 3.3 MCP Tools

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/v1/mcpServer/getMcpServerTools` | 获取某 MCP Server 的工具列表（结构对齐 MCP `tools/list` 语义） |
| POST | `/v1/mcpServer/getMcpServerToolsByUUID` | 按 UUID 获取工具列表 |
| POST | `/v1/mcpServer/createMcpServerTool` | 创建工具 |
| POST | `/v1/mcpServer/updateMcpServerTool` | 更新工具 |
| POST | `/v1/mcpServer/getToolsInfoByUUID` | 按 UUID 获取工具详情 |
| POST | `/v1/mcpServer/testMcpServerTool` | 连通性测试 |

## 4. 其它

- 未匹配路由返回 **404**，JSON：`{"message":"404 not found"}`。
- 协议层行为（如 MCP JSON-RPC 消息格式）遵循项目依赖的 MCP 实现（如 `github.com/ThinkInAIXYZ/go-mcp`）及官方 MCP 规范；本文档仅覆盖本仓库中的 **HTTP 映射与头约定**。

## 5. 许可

软件许可以仓库根目录 `LICENSE`（MIT）为准。
