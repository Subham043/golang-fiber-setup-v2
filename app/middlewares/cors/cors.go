package cors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"go.uber.org/fx"
)

type CorsMiddleware struct {
	Config *config.Config
}

type ICorsMiddleware interface {
	CorsMiddleware() fiber.Handler
}

func NewCorsMiddleware(config *config.Config) *CorsMiddleware {
	return &CorsMiddleware{
		Config: config,
	}
}

// corsConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/cors
func corsConfig(config config.CorsConfig) cors.Config {
	// Define cors settings.

	// Return Cors configuration.
	return cors.Config{
		AllowOrigins:     config.AllowedOrigins,
		AllowMethods:     config.AllowedMethods,
		AllowHeaders:     config.AllowedHeaders,
		AllowCredentials: config.AllowedCredentials,
		ExposeHeaders:    config.ExposeHeaders,
		MaxAge:           config.MaxAge,
	}
}

func (c *CorsMiddleware) CorsMiddleware() fiber.Handler {
	return cors.New(corsConfig(c.Config.Cors))
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewCorsMiddleware),
	)
}
