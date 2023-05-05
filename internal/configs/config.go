package configs

import (
	"sync"

	"github.com/spf13/viper"
)

const ConfigPrefix = "LICENSE_MANAGEMENT"

var (
	configure Configure
	mu        sync.RWMutex
)

type Configure struct {
	POSTGRES_HOST     string `mapstructure:"LICENSE_MANAGEMENT_POSTGRES_HOST"`
	POSTGRES_USER     string `mapstructure:"LICENSE_MANAGEMENT_POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"LICENSE_MANAGEMENT_POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"LICENSE_MANAGEMENT_POSTGRES_DB"`
	POSTGRES_PORT     string `mapstructure:"LICENSE_MANAGEMENT_POSTGRES_PORT"`
	REDIS_PORT        string `mapstructure:"LICENSE_MANAGEMENT_REDIS_PORT"`
	PORT              string `mapstructure:"LICENSE_MANAGEMENT_PORT"`
	CLIENT_ORIGIN     string `mapstructure:"LICENSE_MANAGEMENT_CLIENT_ORIGIN"`
	JWT_SECRET_KEY    string `mapstructure:"LICENSE_MANAGEMENT_JWT_SECRET_KEY"`
	Mode              string `mapstructure:"LICENSE_MANAGEMENT_MODE"`
}

func LoadConfigure(path string) (config Configure, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func GetConfig() (*Configure, error) {
	config, err := LoadConfigure("./conf")
	return &config, err
}
