package helmet

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"go.uber.org/fx"
)

type HelmetMiddleware struct {
}

type IHelmetMiddleware interface {
	HelmetMiddleware() fiber.Handler
}

func NewHelmetMiddleware() *HelmetMiddleware {
	return &HelmetMiddleware{}
}

func (h *HelmetMiddleware) HelmetMiddleware() fiber.Handler {
	return helmet.New()
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewHelmetMiddleware),
	)
}
