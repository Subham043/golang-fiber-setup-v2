package recover

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
)

// recoverConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/recover
func recoverConfig(config config.ServerConfig) recover.Config {
	// Define recover settings.

	// Return recover configuration.
	return recover.Config{
		EnableStackTrace: config.Env != "production",
	}
}

func RecoverMiddleware(config config.ServerConfig) fiber.Handler {
	return recover.New(recoverConfig(config))
}
