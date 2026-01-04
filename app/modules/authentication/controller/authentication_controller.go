package authentication_controller

import (
	"github.com/gofiber/fiber/v2"
	authentication_dto "github.com/subham043/golang-fiber-setup/app/modules/authentication/dto"
	"github.com/subham043/golang-fiber-setup/utils/response"
)

type AuthenticationController struct{}

type IAuthenticationController interface {
	Login(c *fiber.Ctx) error
}

func NewAuthenticationController() *AuthenticationController {
	return &AuthenticationController{}
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
