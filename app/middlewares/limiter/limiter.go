package limiter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/storage/redis/v3"
)

// GlobalLimiterConfig func for configuration Fiber app.
// See: "github.com/gofiber/fiber/v2/middleware/limiter"
func globalLimiterConfig(redis *redis.Storage) limiter.Config {
	// Return Limiter configuration.
	return limiter.Config{
		Max:        100,
		Expiration: 60 * time.Second,
		Storage:    redis,
		// LimitReached: utils.LimitReachedHandler,
	}
}

// AuthLimiterConfig func for configuration Fiber app.
// See: "github.com/gofiber/fiber/v2/middleware/limiter"
func authLimiterConfig(redis *redis.Storage) limiter.Config {
	// Return Limiter configuration.
	return limiter.Config{
		Max:        3,
		Expiration: 60 * time.Second,
		Storage:    redis,
		// LimitReached: utils.LimitReachedHandler,
	}
}

func GlobalLimiterMiddleware(redis *redis.Storage) fiber.Handler {
	return limiter.New(globalLimiterConfig(redis))
}

func AuthLimiterMiddleware(redis *redis.Storage) fiber.Handler {
	return limiter.New(authLimiterConfig(redis))
}
