package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"go.uber.org/fx"
)

type RecoverMiddleware struct {
	Config *config.Config
}

type IRecoverMiddleware interface {
	RecoverMiddleware() fiber.Handler
}

func NewRecoverMiddleware(config *config.Config) *RecoverMiddleware {
	return &RecoverMiddleware{
		Config: config,
	}
}

// recoverConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/recover
func recoverConfig(config config.ServerConfig) recover.Config {
	// Define recover settings.

	// Return recover configuration.
	return recover.Config{
		EnableStackTrace: config.Env != "production",
	}
}

func (r *RecoverMiddleware) RecoverMiddleware() fiber.Handler {
	return recover.New(recoverConfig(r.Config.Server))
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewRecoverMiddleware),
	)
}
