package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

// ReadConfig variable to store ReadConfig result
var ReadConfig = func(path string, cfg interface{}) error {
	return cleanenv.ReadConfig(path, cfg)
}

type PublisherConfig struct {
	Name       string `yaml:"name"`
	RoutingKey string `yaml:"routing_key"`
}

type ConsumerConfig struct {
	Name  string `yaml:"name"`
	Queue string `yaml:"queue"`
}

type ExchangeConfig struct {
	Name string `yaml:"name"`
	Kind string `yaml:"kind"`
}

type BrokerConfig struct {
	URL       string          `yaml:"url"`
	Publisher PublisherConfig `yaml:"publisher"`
	Consumer  ConsumerConfig  `yaml:"consumer"`
	Exchange  ExchangeConfig  `yaml:"exchange"`
}

// BotConfig is a structure to store db config
type BotConfig struct {
	BotToken    string `yaml:"token"`
	BotPolingTO int    `yaml:"poling_to"`
}

// Config is a structure to store bot settings
type Config struct {
	Bot    BotConfig    `yaml:"bot"`
	Broker BrokerConfig `yaml:"broker"`
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
