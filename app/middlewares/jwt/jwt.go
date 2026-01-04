package jwt

import (
	"fmt"
	"time"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	error_handler "github.com/subham043/golang-fiber-setup/utils/error"
	"go.uber.org/fx"
)

type JWTMiddleware struct {
	Config *config.Config
}

type JWTUserDTO struct {
	ID    uuid.UUID
	Email string
	Name  string
}

type IJWTMiddleware interface {
	GenerateAccessToken(dto JWTUserDTO) (string, error)
	GenerateRefreshToken(dto JWTUserDTO) (string, error)
	AccessTokenMiddleware() fiber.Handler
	RefreshTokenMiddleware() fiber.Handler
}

func NewJWTMiddleware(config *config.Config) *JWTMiddleware {
	return &JWTMiddleware{
		Config: config,
	}
}

// corsConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/middleware/cors
func jwtConfig(SigningKey jwtMiddleware.SigningKey) jwtMiddleware.Config {
	// Define cors settings.

	// Return Cors configuration.
	return jwtMiddleware.Config{
		SigningKey:   SigningKey,
		ContextKey:   "user", // used in private routes
		ErrorHandler: error_handler.JWTErrorHandler,
	}
}

// JWTProtected func for specify routes group with JWT authentication.
// See: https://github.com/gofiber/contrib/jwt
func (j *JWTMiddleware) AccessTokenMiddleware() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtConfig(jwtMiddleware.SigningKey{Key: []byte(j.Config.JWT.SecretKey)})
	config.TokenLookup = "header:Authorization"

	return jwtMiddleware.New(config)
}

func (j *JWTMiddleware) RefreshTokenMiddleware() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtConfig(jwtMiddleware.SigningKey{Key: []byte(j.Config.JWT.RefreshKey)})
	config.TokenLookup = fmt.Sprintf("header:Authorization,cookie:%s", j.Config.Session.SessionCookie)

	return jwtMiddleware.New(config)
}

func (j *JWTMiddleware) GenerateAccessToken(dto JWTUserDTO) (string, error) {

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    dto.ID,
		"email": dto.Email,
		"name":  dto.Name,
		"exp":   time.Now().Add(time.Minute * time.Duration(j.Config.JWT.SecretKeyExpiration)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.Config.JWT.SecretKey))
}

func (j *JWTMiddleware) GenerateRefreshToken(dto JWTUserDTO) (string, error) {

	// Create the Claims
	claims := jwt.MapClaims{
		"id":    dto.ID,
		"email": dto.Email,
		"name":  dto.Name,
		"exp":   time.Now().Add(time.Minute * time.Duration(j.Config.JWT.RefreshKeyExpiration)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.Config.JWT.RefreshKey))
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewJWTMiddleware),
	)
}
