package config

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
}

func New(filepath string) Config {
	v := viper.New()

	if filepath != "" {
		v.SetConfigFile(filepath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("toml")
		v.AddConfigPath("/etc/gitjika")
		v.AddConfigPath("$HOME/.gitjika")
	}

	if err := v.ReadInConfig(); err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
		panic(err)
	}

	return Config{
		Server: loadServerConfig(v.Sub("server")),
	}
}
