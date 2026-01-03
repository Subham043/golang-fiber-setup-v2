package limiter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis/v3"
	error_handler "github.com/subham043/golang-fiber-setup/utils/error"
)

// CommonLimiterConfig func for configuration Fiber app.
// See: "github.com/gofiber/fiber/v2/middleware/limiter"
func commonLimiterConfig(redis *redis.Storage) limiter.Config {
	return limiter.Config{
		Storage:      redis,
		LimitReached: error_handler.LimititerErrorHandler,
	}
}

func GlobalLimiterMiddleware(redis *redis.Storage) fiber.Handler {
	cfg := commonLimiterConfig(redis)
	cfg.Max = 100
	cfg.Expiration = 60 * time.Second
	return limiter.New(cfg)
}

func AuthLimiterMiddleware(redis *redis.Storage) fiber.Handler {
	cfg := commonLimiterConfig(redis)
	cfg.Max = 3
	cfg.Expiration = 60 * time.Second
	return limiter.New(cfg)
}
