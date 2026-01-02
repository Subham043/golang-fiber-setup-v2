package bootstrap

import (
	"github.com/subham043/golang-fiber-setup/app/middlewares"
	"github.com/subham043/golang-fiber-setup/app/modules/health"
	"github.com/subham043/golang-fiber-setup/app/router"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"github.com/subham043/golang-fiber-setup/bootstrap/logger"
	"github.com/subham043/golang-fiber-setup/bootstrap/server"
	"github.com/subham043/golang-fiber-setup/bootstrap/storage/redis"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		config.Module(),
		logger.Module(),
		redis.Module(),
		middlewares.Module(),
		health.Module(),
		router.Module(),
		server.Module(),
	)
}
