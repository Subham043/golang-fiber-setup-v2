package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/subham043/golang-fiber-setup/app/middlewares/compress"
	"github.com/subham043/golang-fiber-setup/app/middlewares/cors"
	"github.com/subham043/golang-fiber-setup/app/middlewares/encrypt_cookie"
	"github.com/subham043/golang-fiber-setup/app/middlewares/helmet"
	"github.com/subham043/golang-fiber-setup/app/middlewares/limiter"
	"github.com/subham043/golang-fiber-setup/app/middlewares/logger"
	"github.com/subham043/golang-fiber-setup/app/middlewares/recover"
	"github.com/subham043/golang-fiber-setup/app/middlewares/request_id"
	"go.uber.org/fx"
)

type Middleware struct {
	App                 *fiber.App
	Limiter             *limiter.LimiterMiddleware
	GlobalEncryptCookie *encrypt_cookie.EncryptCookieMiddleware
	Cors                *cors.CorsMiddleware
	Recover             *recover.RecoverMiddleware
	Logger              *logger.LoggerMiddleware
	Helmet              *helmet.HelmetMiddleware
	Compress            *compress.CompressMiddleware
	RequestID           *request_id.RequestIDMiddleware
}

func NewMiddleware(app *fiber.App, limiter *limiter.LimiterMiddleware, globalEncryptCookie *encrypt_cookie.EncryptCookieMiddleware, cors *cors.CorsMiddleware, recover *recover.RecoverMiddleware, logger *logger.LoggerMiddleware, helmet *helmet.HelmetMiddleware, compress *compress.CompressMiddleware, requestID *request_id.RequestIDMiddleware) *Middleware {
	return &Middleware{
		App:                 app,
		Limiter:             limiter,
		GlobalEncryptCookie: globalEncryptCookie,
		Cors:                cors,
		Recover:             recover,
		Logger:              logger,
		Helmet:              helmet,
		Compress:            compress,
		RequestID:           requestID,
	}
}

// Add Global Middlewares
func (m *Middleware) Register() {

	m.App.Use(m.Cors.CorsMiddleware())

	m.App.Use(m.Helmet.HelmetMiddleware())

	m.App.Use(m.Compress.CompressMiddleware())

	m.App.Use(m.RequestID.RequestIDMiddleware())

	m.App.Use(m.Recover.RecoverMiddleware())

	m.App.Use(m.Limiter.GlobalLimiterMiddleware())

	m.App.Use(m.GlobalEncryptCookie.EncryptCookieMiddleware())

	m.App.Use(m.Logger.ZapHTTPLoggerMiddleware())

	// m.App.Use(pprof.New(pprof.Config{
	// 	Next: utils.IsEnabled(m.Cfg.Middleware.Pprof.Enable),
	// }))

	// m.App.Use(filesystem.New(filesystem.Config{
	// 	Next:   utils.IsEnabled(m.Cfg.Middleware.Filesystem.Enable),
	// 	Root:   http.Dir(m.Cfg.Middleware.Filesystem.Root),
	// 	Browse: m.Cfg.Middleware.Filesystem.Browse,
	// 	MaxAge: m.Cfg.Middleware.Filesystem.MaxAge,
	// }))

	// m.App.Get(m.Cfg.Middleware.Monitor.Path, monitor.New(monitor.Config{
	// 	Next: utils.IsEnabled(m.Cfg.Middleware.Monitor.Enable),
	// }))
}

// Module returns a fx.Option that configures the middlewares.
func Module() fx.Option {
	return fx.Options(
		limiter.Module(),
		encrypt_cookie.Module(),
		cors.Module(),
		recover.Module(),
		logger.Module(),
		helmet.Module(),
		compress.Module(),
		request_id.Module(),
		fx.Provide(NewMiddleware),
	)
}
