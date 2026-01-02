package redis

import (
	"fmt"

	redis_storage "github.com/gofiber/storage/redis/v3"
	redis_client "github.com/redis/go-redis/v9"
	"github.com/subham043/golang-fiber-setup/bootstrap/config"
	"go.uber.org/fx"
)

type RedisStorage = redis_storage.Storage

func NewRedisStorage(cfg *config.Config) *RedisStorage {

	rdb := redis_client.NewClient(&redis_client.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DBNumber,
	})

	store := redis_storage.NewFromConnection(rdb)

	return store
}

// Module returns a fx.Option that configures the redis storage.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(NewRedisStorage),
	)
}
