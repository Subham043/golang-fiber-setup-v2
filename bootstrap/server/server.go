package server

import (
	"context"
	"fmt"
	"time"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/golang-fiber-setup/app/middlewares"
	"github.com/subham043/golang-fiber-setup/app/router"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	error_handler "github.com/subham043/golang-fiber-setup/utils/error"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Module returns a fx.Option that configures the webserver.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewServer),
		fx.Invoke(Start),
	)
}

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func fiberConfig(config config.ServerConfig) fiber.Config {
	// Define server settings.
	// Return Fiber configuration.
	return fiber.Config{
		JSONEncoder:              json.Marshal,
		JSONDecoder:              json.Unmarshal,
		ServerHeader:             "Fiber",
		AppName:                  config.AppName,
		CaseSensitive:            true,
		EnablePrintRoutes:        config.Env != "production",
		EnableSplittingOnParsers: true,
		Prefork:                  true,
		ReadTimeout:              time.Second * time.Duration(config.ReadTimeout),
		ErrorHandler:             error_handler.ServerErrorHandler,
	}
}

// NewServer creates a new fiber server instance.
func NewServer(config *config.Config) *fiber.App {
	app := fiber.New(fiberConfig(config.Server))
	return app
}

// Start starts the fiber server.
func Start(lifecycle fx.Lifecycle, server *fiber.App, config *config.Config, log *zap.Logger, middlewares *middlewares.Middleware, router *router.Router) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Register middlewares & routes
			middlewares.Register()
			router.Register()

			go func() {
				addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
				// addr := fmt.Sprintf(":%d", config.Server.Port)
				if err := server.Listen(addr); err != nil {
					log.Error("server failed to start", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := server.Shutdown(); err != nil {
				log.Error("server failed to stop", zap.Error(err))
			}
			log.Sync()
			return nil
		},
	})
}
