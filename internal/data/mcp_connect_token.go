package data

import (
	"buffgpt-flow-bridge-mcp/internal/biz"
	"buffgpt-flow-bridge-mcp/internal/data/database"
	"buffgpt-flow-bridge-mcp/internal/data/model"
	"buffgpt-flow-bridge-mcp/pkg/logger"
	"context"
)

type McpConnectToken struct {
	data *database.Data
	log  *logger.Logger
}

func NewMcpConnectToken(data *database.Data, log *logger.Logger) biz.McpConnectTokenRepo {
	return &McpConnectToken{
		data: data,
		log:  log,
	}
}

func (m *McpConnectToken) Create(ctx context.Context, mcpConnectToken *model.McpConnectToken) (err error) {
	err = m.data.Db.WithContext(ctx).Create(mcpConnectToken).Error
	if err != nil {
		m.log.Error("create mcp connect token error: %v", err)
		return
	}
	return
}
