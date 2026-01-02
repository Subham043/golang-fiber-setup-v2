package encrypt_cookie

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
)

func encryptCookieConfig(config config.ServerConfig) encryptcookie.Config {
	return encryptcookie.Config{
		Key: config.AppKey,
	}
}

func EncryptCookieMiddleware(config config.ServerConfig) fiber.Handler {
	return encryptcookie.New(encryptCookieConfig(config))
}
