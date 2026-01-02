package logger

import (
	"os"

	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(config *config.Config) (*zap.Logger, error) {
	if config.Server.Env == "production" {
		return productionLogger()
	}
	return developmentLogger()
}

func developmentLogger() (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()

	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 🎨 COLORS
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"

	return cfg.Build()
}

func productionLogger() (*zap.Logger, error) {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:      "timestamp",
		LevelKey:     "level",
		MessageKey:   "msg",
		CallerKey:    "caller",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zapcore.InfoLevel,
	)

	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel)), nil
}

// Module returns a fx.Option that configures the config.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewLogger),

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log.Named("fx").WithOptions(zap.IncreaseLevel(zap.ErrorLevel))}
		}),
	)
}
