package encrypt_cookie

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"go.uber.org/fx"
)

type EncryptCookieMiddleware struct {
	Config *config.Config
}

type IEncryptCookieMiddleware interface {
	EncryptCookieMiddleware() fiber.Handler
}

func NewEncryptCookieMiddleware(config *config.Config) *EncryptCookieMiddleware {
	return &EncryptCookieMiddleware{
		Config: config,
	}
}

func encryptCookieConfig(config config.ServerConfig) encryptcookie.Config {
	return encryptcookie.Config{
		Key: config.AppKey,
	}
}

func (e *EncryptCookieMiddleware) EncryptCookieMiddleware() fiber.Handler {
	return encryptcookie.New(encryptCookieConfig(e.Config.Server))
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewEncryptCookieMiddleware),
	)
}
