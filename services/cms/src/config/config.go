package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

var ReadConfig = func(path string, cfg interface{}) error {
	return cleanenv.ReadConfig(path, cfg)
}

type DatabaseConfig struct {
	Url string `yaml:"postgresql_url"`
}

type MsgBrokerConfig struct {
	Url string `yaml:"rabbitmq_url"`
}

type Config struct {
	Database  DatabaseConfig  `yaml:"db"`
	MsgBroker MsgBrokerConfig `yaml:"msgBroker"`
}

func Load() (*Config, error) {
	var cfg Config

	path := "./config/config.yaml"

	if err := ReadConfig(path, &cfg); err != nil {
		return nil, errors.New("failed to load config: " + err.Error())
	}
	return &cfg, nil
}
