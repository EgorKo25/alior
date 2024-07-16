package config

import (
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
	Env       string          `yaml:"env" env-default:"local"`
	Database  DatabaseConfig  `yaml:"db"`
	MsgBroker MsgBrokerConfig `yaml:"msgBroker"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("Config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("Failed to load config: " + err.Error())
	}
	return &cfg
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
