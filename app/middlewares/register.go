package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/storage/redis/v3"
	"github.com/subham043/golang-fiber-setup/app/middlewares/compress"
	"github.com/subham043/golang-fiber-setup/app/middlewares/cors"
	"github.com/subham043/golang-fiber-setup/app/middlewares/encrypt_cookie"
	"github.com/subham043/golang-fiber-setup/app/middlewares/helmet"
	"github.com/subham043/golang-fiber-setup/app/middlewares/limiter"
	"github.com/subham043/golang-fiber-setup/app/middlewares/logger"
	"github.com/subham043/golang-fiber-setup/app/middlewares/recover"
	"github.com/subham043/golang-fiber-setup/app/middlewares/request_id"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Middleware struct {
	App   *fiber.App
	Cfg   *config.Config
	Redis *redis.Storage
	Log   *zap.Logger
}

func NewMiddleware(app *fiber.App, cfg *config.Config, redis *redis.Storage, log *zap.Logger) *Middleware {
	return &Middleware{
		App:   app,
		Cfg:   cfg,
		Redis: redis,
		Log:   log,
	}
}

// Add Global Middlewares
func (m *Middleware) Register() {

	m.App.Use(cors.CorsMiddleware(m.Cfg.Cors))

	m.App.Use(helmet.HelmetMiddleware())

	m.App.Use(compress.CompressMiddleware())

	m.App.Use(request_id.RequestIDMiddleware())

	m.App.Use(recover.RecoverMiddleware(m.Cfg.Server))

	m.App.Use(limiter.GlobalLimiterMiddleware(m.Redis))

	m.App.Use(encrypt_cookie.EncryptCookieMiddleware(m.Cfg.Server))

	m.App.Use(logger.ZapHTTPLoggerMiddleware(m.Log))

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
		fx.Provide(NewMiddleware),
	)
}
