package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/golang-fiber-setup/app/modules/authentication"
	"github.com/subham043/golang-fiber-setup/app/modules/health"
	"go.uber.org/fx"
)

type Router struct {
	App          fiber.Router
	HealthRouter *health.HealthRouter
	AuthRouter   *authentication.AuthenticationRouter
}

func NewRouter(fiber *fiber.App, healthRouter *health.HealthRouter, authRouter *authentication.AuthenticationRouter) *Router {
	return &Router{
		App:          fiber,
		HealthRouter: healthRouter,
		AuthRouter:   authRouter,
	}
}

func (r *Router) Register() {
	// Test Routes
	r.App.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong! 👋")
	})

	// Health Routes
	r.HealthRouter.RegisterHealthRoutes()

	// Authentication Routes
	r.AuthRouter.RegisterAuthenticationRoutes()
}

// Module returns a fx.Option that configures the router.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewRouter),
		health.Module(),
		authentication.Module(),
	)
}
