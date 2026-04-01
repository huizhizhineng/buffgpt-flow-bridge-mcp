package biz

import (
	"buffgpt-flow-bridge-mcp/internal/data/model"
	"context"
)

type McpConnectTokenRepo interface {
	Create(ctx context.Context, serverInfo *model.McpConnectToken) (err error)
}
