package compress

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"go.uber.org/fx"
)

type CompressMiddleware struct {
}

type ICompressMiddleware interface {
	CompressMiddleware() fiber.Handler
}

func NewCompressMiddleware() *CompressMiddleware {
	return &CompressMiddleware{}
}

// compressConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/compress
func compressConfig() compress.Config {
	// Define compress settings.

	// Return compress configuration.
	return compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}
}

func (c *CompressMiddleware) CompressMiddleware() fiber.Handler {
	return compress.New(compressConfig())
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewCompressMiddleware),
	)
}
