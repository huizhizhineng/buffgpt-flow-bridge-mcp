//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"buffgpt-flow-bridge-mcp/internal/biz"
	"buffgpt-flow-bridge-mcp/internal/conf"
	"buffgpt-flow-bridge-mcp/internal/data"
	"buffgpt-flow-bridge-mcp/internal/data/database"
	"buffgpt-flow-bridge-mcp/internal/mcp/proxy"
	"buffgpt-flow-bridge-mcp/internal/mcp/server"
	"buffgpt-flow-bridge-mcp/internal/mcp/transformer/openapi"
	"buffgpt-flow-bridge-mcp/internal/pkg/cache"
	"buffgpt-flow-bridge-mcp/internal/pkg/startup"
	"buffgpt-flow-bridge-mcp/internal/service"
	"buffgpt-flow-bridge-mcp/middleware"
	"buffgpt-flow-bridge-mcp/pkg/logger"
	"buffgpt-flow-bridge-mcp/router"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// initApp init gin application.
func initApp(config *conf.Conf) (*gin.Engine, func(), error) {
	panic(wire.Build(
		middleware.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		router.ProviderSet,
		logger.ProviderSet,
		openapi.ProviderSet,
		database.ProviderSet,
		server.ProviderSet,
		cache.ProviderSet,
		startup.ProviderSet,
		proxy.ProviderSet,
	))
}
