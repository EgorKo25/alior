package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "./config/config.yaml"

func NewAPIGWConfig() (*APIGWConfig, error) {
	cfg := &APIGWConfig{}
	if err := cfg.loadConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}

type APIGWConfig struct {
	*ServerConfig `yaml:"server-config"`
	*BrokerConfig `yaml:"broker-config"`
}

type BrokerConfig struct {
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	Exchange          string `yaml:"exchange"`
	NeedConnect2Queue bool   `yaml:"need_connect_to_queue"`
	// не обязатeльные поля при NeedConnect2Queue false
	ExchangeKind string `yaml:"exchange-kind"`
	RoutingKey   string `yaml:"routing-key"`
	QueueName    string `yaml:"queue-name"`
}

type ServerConfig struct {
	TlsCert string `yaml:"tls-cert"`
	KeyFile string `yaml:"key-file"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
}

func (a *APIGWConfig) loadConfig() error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&a); err != nil {
		return err
	}
	return err
}
