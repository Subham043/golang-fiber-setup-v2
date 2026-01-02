package compress

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

// compressConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/compress
func compressConfig() compress.Config {
	// Define compress settings.

	// Return compress configuration.
	return compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}
}

func CompressMiddleware() fiber.Handler {
	return compress.New(compressConfig())
}
