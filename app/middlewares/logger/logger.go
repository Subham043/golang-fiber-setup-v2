package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	Log *zap.Logger
}

type ILoggerMiddleware interface {
	ZapHTTPLoggerMiddleware() fiber.Handler
}

func NewLoggerMiddleware(log *zap.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		Log: log,
	}
}

func (l *LoggerMiddleware) ZapHTTPLoggerMiddleware() fiber.Handler {
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
			l.Log.Error("http_request", fields...)
		case status >= 400:
			l.Log.Warn("http_request", fields...)
		default:
			l.Log.Info("http_request", fields...)
		}

		return err
	}
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewLoggerMiddleware),
	)
}
