package health_controller

import (
	"github.com/gofiber/fiber/v2"
	health_service "github.com/subham043/golang-fiber-setup/app/modules/health/service"
)

type HealthController struct {
	healthService *health_service.HealthService
}

type IHealthController interface {
	Index(c *fiber.Ctx) error
}

func NewHealthController(healthService *health_service.HealthService) *HealthController {
	return &HealthController{
		healthService: healthService,
	}
}

func (h *HealthController) Index(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(map[string]string{
		"message": h.healthService.Index(),
	})
}
