package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
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

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, errors.New("failed to load config: " + err.Error())
	}

	envBotToken := os.Getenv("TOKEN")
	if envBotToken != "" {
		cfg.Bot.BotToken = envBotToken
	}

	if cfg.Bot.BotToken == "" {
		return nil, errors.New("bot token is not provided")
	}

	return &cfg, nil
}
