package authentication_controller

import (
	"github.com/gofiber/fiber/v2"
	authentication_dto "github.com/subham043/golang-fiber-setup/app/modules/authentication/dto"
	authentication_service "github.com/subham043/golang-fiber-setup/app/modules/authentication/service"
	"github.com/subham043/golang-fiber-setup/utils/response"
)

type AuthenticationController struct {
	authenticationService *authentication_service.AuthenticationService
}

type IAuthenticationController interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

func NewAuthenticationController(authenticationService *authentication_service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: authenticationService}
}

func (controller *AuthenticationController) Login(c *fiber.Ctx) error {

	payload := authentication_dto.NewSignInPayload()

	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	return response.Json(c, response.Response{
		Message: "Login",
		Data:    fiber.Map{"email": payload.Email},
	})
}

func (controller *AuthenticationController) Register(c *fiber.Ctx) error {

	payload := authentication_dto.NewSignUpPayload()

	if err := c.BodyParser(payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	if err := payload.HashPassword(); err != nil {
		return err
	}
	result, err := controller.authenticationService.Register(c.Context(), payload)
	if err != nil {
		return err
	}
	return response.Json(c, response.Response{
		Message: "Registered successfully",
		Data:    result,
	})
}
