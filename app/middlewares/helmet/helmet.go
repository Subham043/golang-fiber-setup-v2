package helmet

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func HelmetMiddleware() fiber.Handler {
	return helmet.New()
}
