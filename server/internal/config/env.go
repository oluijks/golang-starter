package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServerHost        string        `mapstructure:"HTTP_SERVER_HOST"`
	HTTPServerPort        string        `mapstructure:"HTTP_SERVER_PORT"`
	FrondendURL           string        `mapstructure:"FRONDEND_URL"`
	MongoDBURL            string        `mapstructure:"MONGODB_URL"`
	DatabaseName          string        `mapstructure:"DATABASE_NAME"`
	AccountCollectionName string        `mapstructure:"ACCOUNT_COLLECTION_NAME"`
	SymmectricTokenKey    string        `mapstructure:"SYMMETRIC_TOKEN_KEY"`
	AccessTokenDuration   time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func NewConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.ReadInConfig()

	var cfg Config
	viper.Unmarshal(&cfg)
	return &cfg
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
