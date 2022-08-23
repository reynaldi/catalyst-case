package config

type Config struct {
	ConnectionString string `env:"CONNECTION_STRING"`
	Dialect          string `env:"DIALECT"`
	AppPort          string `env:"APP_PORT"`
	AppHost          string `env:"APP_HOST"`
}
