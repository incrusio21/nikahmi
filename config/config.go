package config

import (
	"bytes"
	_ "embed"

	"github.com/spf13/viper"
)

//go:embed .yaml
var defaultConfiguration []byte

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Config struct {
	DB *Database
}

var viper_config = viper.New()

func init() {
	// Configuration file
	viper_config.SetConfigType("yaml")

	// Read configuration
	if err := viper_config.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		panic(err)
	}

}

func Read() (*Config, error) {
	var config Config

	// Unmarshal the configuration
	if err := viper_config.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
