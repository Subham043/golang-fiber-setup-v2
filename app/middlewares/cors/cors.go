package cors

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
)

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

func CorsMiddleware(config config.CorsConfig) fiber.Handler {
	return cors.New(corsConfig(config))
}
