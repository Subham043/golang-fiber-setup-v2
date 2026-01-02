package health

import (
	"github.com/gofiber/fiber/v2"
	health_controller "github.com/subham043/golang-fiber-setup/app/modules/health/controller"
	health_service "github.com/subham043/golang-fiber-setup/app/modules/health/service"
	"go.uber.org/fx"
)

type HealthRouter struct {
	App        fiber.Router
	Controller *health_controller.HealthController
}

// NewHealthRouter creates a new health router.
func NewHealthRouter(fiber *fiber.App, controller *health_controller.HealthController) *HealthRouter {
	return &HealthRouter{
		App:        fiber,
		Controller: controller,
	}
}

// Registers health routes
func (r *HealthRouter) RegisterHealthRoutes() {

	// Define routes
	r.App.Route("/health", func(router fiber.Router) {
		router.Get("/", r.Controller.Index)
	})
}

// Module returns a fx.Option that configures the health module.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(health_controller.NewHealthController),
		fx.Provide(health_service.NewHealthService),
		fx.Provide(NewHealthRouter),
	)
}
