package config

import "github.com/spf13/viper"

type ServerConfig struct {
	Host string
	Port uint16
}

const (
	defaultHost = "127.0.0.1"
	defaultPort = 6969
)

func loadServerConfig(v *viper.Viper) ServerConfig {
	cfg := ServerConfig{Host: defaultHost, Port: defaultPort}

	if v == nil {
		return cfg
	}

	if v.IsSet("host") {
		cfg.Host = v.GetString("host")
	}

	if v.IsSet("port") {
		cfg.Port = v.GetUint16("port")
	}

	return cfg
}
