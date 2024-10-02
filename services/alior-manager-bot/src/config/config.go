package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

// ReadConfig variable to store ReadConfig result
var ReadConfig = func(path string, cfg interface{}) error {
	return cleanenv.ReadConfig(path, cfg)
}

// BotConfig is a structure to store db config
type BotConfig struct {
	BotToken    string `yaml:"token"`
	BotPolingTO int    `yaml:"poling_to"`
}

// Config is a structure to store bot settings
type Config struct {
	Bot BotConfig `yaml:"bot"`
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
