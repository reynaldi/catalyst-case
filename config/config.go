package config

type Config struct {
	ConnectionString string `env:"CONNECTION_STRING"`
	Dialect          string `env:"DIALECT"`
}
