package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

// ReadConfig variable to store ReadConfig result
var ReadConfig = func(path string, cfg interface{}) error {
	return cleanenv.ReadConfig(path, cfg)
}

// DatabaseConfig is a structure to store db config
type DatabaseConfig struct {
	URL string `yaml:"postgresql_url"`
}

// MsgBrokerConfig is a structure to store broker config
type MsgBrokerConfig struct {
	URL string `yaml:"rabbitmq_url"`
}

// Config is a structure to store db and broker configs
type Config struct {
	Database  DatabaseConfig  `yaml:"db"`
	MsgBroker MsgBrokerConfig `yaml:"msgBroker"`
}

// Load is a function to load config from file
func Load() (*Config, error) {
	var cfg Config

	path := "./config/config.yaml"

	if err := ReadConfig(path, &cfg); err != nil {
		return nil, errors.New("failed to load config: " + err.Error())
	}
	return &cfg, nil
}
