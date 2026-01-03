package config

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type ServerConfig struct {
	Env         string
	Host        string
	Port        int
	ReadTimeout int
	AppKey      string
	AppName     string
}

type JWTConfig struct {
	SecretKey            string
	SecretKeyExpiration  int
	RefreshKey           string
	RefreshKeyExpiration int
}

type DBConfig struct {
	Type                  string
	Host                  string
	Port                  int
	User                  string
	Password              string
	Name                  string
	MaxConnection         int
	MaxIdleConnection     int
	MaxLifeTimeConnection int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DBNumber int
}

type CorsConfig struct {
	AllowedOrigins     string
	AllowedMethods     string
	AllowedHeaders     string
	AllowedCredentials bool
	ExposeHeaders      string
	MaxAge             int
}

type MinioConfig struct {
	MinioRootUser     string
	MinioRootPassword string
	MinioPort         int
	MinioEndpoint     string
}

type MailConfig struct {
	MailHost        string
	MailPort        int
	MailUsername    string
	MailPassword    string
	MailFromAddress string
	MailFromName    string
}

type RabbitMQConfig struct {
	RabbitMQHost      string
	RabbitMQPort      int
	RabbitMQUsername  string
	RabbitMQPassword  string
	RabbitMQQueueName string
}

type SessionConfig struct {
	SessionCookie   string
	SessionPath     string
	SessionDomain   string
	SessionLifetime int
	SessionSecure   bool
	SessionHttpOnly bool
}

type Config struct {
	Server   ServerConfig
	JWT      JWTConfig
	DB       DBConfig
	Redis    RedisConfig
	Cors     CorsConfig
	Minio    MinioConfig
	Mail     MailConfig
	RabbitMQ RabbitMQConfig
	Session  SessionConfig
}

func LoadConfig() *Config {
	viper.SetConfigFile("../docker/.env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config := &Config{
		Server: ServerConfig{
			Env:         viper.GetString("NODE_ENV"),
			Host:        viper.GetString("SERVER_HOST"),
			Port:        viper.GetInt("SERVER_PORT"),
			ReadTimeout: viper.GetInt("SERVER_READ_TIMEOUT"),
			AppKey:      viper.GetString("SERVER_APP_KEY"),
			AppName:     viper.GetString("SERVER_APP_NAME"),
		},
		DB: DBConfig{
			Type:                  viper.GetString("DB_TYPE"),
			Host:                  viper.GetString("DB_HOST"),
			Port:                  viper.GetInt("DB_PORT"),
			User:                  viper.GetString("DB_USER"),
			Password:              viper.GetString("DB_PASSWORD"),
			Name:                  viper.GetString("DB_NAME"),
			MaxConnection:         viper.GetInt("DB_MAX_CONNECTION"),
			MaxIdleConnection:     viper.GetInt("DB_MAX_IDLE_CONNECTION"),
			MaxLifeTimeConnection: viper.GetInt("DB_MAX_LIFE_TIME_CONNECTION"),
		},
		JWT: JWTConfig{
			SecretKey:            viper.GetString("JWT_SECRET_KEY"),
			SecretKeyExpiration:  viper.GetInt("JWT_SECRET_KEY_EXPIRATION"),
			RefreshKey:           viper.GetString("JWT_REFRESH_KEY"),
			RefreshKeyExpiration: viper.GetInt("JWT_REFRESH_KEY_EXPIRATION"),
		},
		Redis: RedisConfig{
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetInt("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DBNumber: viper.GetInt("REDIS_DB_NUMBER"),
		},
		Cors: CorsConfig{
			AllowedOrigins:     viper.GetString("CORS_ALLOWED_ORIGINS"),
			AllowedMethods:     viper.GetString("CORS_ALLOWED_METHODS"),
			AllowedHeaders:     viper.GetString("CORS_ALLOWED_HEADERS"),
			AllowedCredentials: viper.GetBool("CORS_ALLOWED_CREDENTIALS"),
			ExposeHeaders:      viper.GetString("CORS_EXPOSE_HEADERS"),
			MaxAge:             viper.GetInt("CORS_MAX_AGE"),
		},
		Minio: MinioConfig{
			MinioRootUser:     viper.GetString("MINIO_ROOT_USER"),
			MinioRootPassword: viper.GetString("MINIO_ROOT_PASSWORD"),
			MinioPort:         viper.GetInt("MINIO_PORT"),
			MinioEndpoint:     viper.GetString("MINIO_ENDPOINT"),
		},
		Mail: MailConfig{
			MailHost:        viper.GetString("MAIL_HOST"),
			MailPort:        viper.GetInt("MAIL_PORT"),
			MailUsername:    viper.GetString("MAIL_USERNAME"),
			MailPassword:    viper.GetString("MAIL_PASSWORD"),
			MailFromAddress: viper.GetString("MAIL_FROM_ADDRESS"),
			MailFromName:    viper.GetString("MAIL_FROM_NAME"),
		},
		RabbitMQ: RabbitMQConfig{
			RabbitMQHost:      viper.GetString("RABBITMQ_HOST"),
			RabbitMQPort:      viper.GetInt("RABBITMQ_PORT"),
			RabbitMQUsername:  viper.GetString("RABBITMQ_USERNAME"),
			RabbitMQPassword:  viper.GetString("RABBITMQ_PASSWORD"),
			RabbitMQQueueName: viper.GetString("RABBITMQ_QUEUE_NAME"),
		},
		Session: SessionConfig{
			SessionCookie:   viper.GetString("SESSION_COOKIE"),
			SessionPath:     viper.GetString("SESSION_PATH"),
			SessionDomain:   viper.GetString("SESSION_DOMAIN"),
			SessionLifetime: viper.GetInt("SESSION_LIFETIME"),
			SessionSecure:   viper.GetBool("SESSION_SECURE"),
			SessionHttpOnly: viper.GetBool("SESSION_HTTP_ONLY"),
		},
	}

	return config
}

// Module returns a fx.Option that configures the config.
func Module() fx.Option {
	return fx.Options(
		fx.Provide(LoadConfig),
	)
}
