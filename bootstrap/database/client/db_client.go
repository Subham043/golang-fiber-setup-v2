package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"github.com/subham043/golang-fiber-setup/bootstrap/database/ent"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewDBClient(config *config.Config, log *zap.Logger) *ent.Client {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Error("failed opening connection to mysql: ", zap.Error(err))
		panic("failed opening connection to mysql")
	}
	log.Info("✅ Connected to MySQL")
	return client
}

// Start starts the fiber server.
func Connect(lifecycle fx.Lifecycle, config *config.Config, log *zap.Logger, client *ent.Client) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("📦 Running Ent migrations")
			// return client.Schema.Create(ctx)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("🛑 Closing Ent client")
			return client.Close()
		},
	})
}

func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewDBClient),
		fx.Invoke(Connect),
	)
}
