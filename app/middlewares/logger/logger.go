package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func ZapHTTPLoggerMiddleware(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process request
		err := c.Next()

		latency := time.Since(start)

		status := c.Response().StatusCode()

		fields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Duration("latency", latency),
			zap.String("ip", c.IP()),
			zap.String("user_agent", string(c.Request().Header.UserAgent())),
		}

		switch {
		case status >= 500:
			log.Error("http_request", fields...)
		case status >= 400:
			log.Warn("http_request", fields...)
		default:
			log.Info("http_request", fields...)
		}

		return err
	}
}
