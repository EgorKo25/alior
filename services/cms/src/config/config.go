package config

import (
	"errors"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

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
	path := fetchConfigPath()

	if path == "" {
		return nil, errors.New("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.New("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, errors.New("Failed to load config: " + err.Error())
	}
	return &cfg, nil
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
