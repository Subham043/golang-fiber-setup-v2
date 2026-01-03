package health_controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	health_service "github.com/subham043/golang-fiber-setup/app/modules/health/service"
	response "github.com/subham043/golang-fiber-setup/utils/response"
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

	return response.Json(c, response.Response{
		Message: h.healthService.Index(),
		Data:    fiber.Map{"time": time.Now().Format("2006-01-02 15:04:05")},
	})
}
