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
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	RedisPort      string `mapstructure:"REDIS_PORT"`
	ServerPort     string `mapstructure:"PORT"`
	JWT_KEY        string `mapstructure:"JWTSECRET_KEY"`
	Mode           string `mapstructure:"MODE"`
	ClientOrigin   string `mapstructure:"CLIENT_ORIGIN"`
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
