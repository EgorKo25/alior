package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Database struct {
	Engine       string `yaml:"engine"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	SSLMode      string `yaml:"sslmode"`
	DatabaseName string `yaml:"dbName"`
	User         string `yaml:"dbUser"`
	UserPassword string `yaml:"dbPassword"`
}

type Config struct {
	Databases struct {
		SMSDatabase Database `yaml:"smsDatabase"`
	} `yaml:"databases"`
}

func LoadConfig() (*Config, error) {
	path := "./config/config.yaml" // TODO: Брать путь из окружения или типо того

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Config file does not exist: " + path)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("Failed to load config: " + err.Error())
	}

	var config Config
	err = yaml.Unmarshal(data, &config)

	if err != nil {
		return nil, errors.New("Failed to unmarshal config data: " + err.Error())
	}

	return &config, nil
}

// TODO: Переделать как-нибудь (мб в качестве аргумента брать конфиг)
// А еще лучше наверное это вынести в database.go, куда передовать конфиг - хз.
func BuildConnString(db *Database) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.UserPassword, db.DatabaseName, db.SSLMode)
}
