package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/golang-fiber-setup/app/middlewares/limiter"
	authentication_controller "github.com/subham043/golang-fiber-setup/app/modules/authentication/controller"
	"go.uber.org/fx"
)

type AuthenticationRouter struct {
	App        fiber.Router
	Controller *authentication_controller.AuthenticationController
	Middleware *limiter.LimiterMiddleware
}

// NewAuthenticationRouter creates a new authentication router.
func NewAuthenticationRouter(fiber *fiber.App, controller *authentication_controller.AuthenticationController, middleware *limiter.LimiterMiddleware) *AuthenticationRouter {
	return &AuthenticationRouter{
		App:        fiber,
		Controller: controller,
		Middleware: middleware,
	}
}

// Registers health routes
func (r *AuthenticationRouter) RegisterAuthenticationRoutes() {
	// Create a group
	group := r.App.Group("/api/v1/auth")

	// Apply middleware to the group
	// group.Use(r.Middleware.AuthLimiterMiddleware())

	// Define routes
	group.Post("/login", r.Controller.Login)
}

// Module returns a fx.Option that configures the health module.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(authentication_controller.NewAuthenticationController),
		fx.Provide(NewAuthenticationRouter),
	)
}
