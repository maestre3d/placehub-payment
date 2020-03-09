package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	HTTPPort string
}

func NewConfig(logger *zap.SugaredLogger) *Config {
	logger.Infow("configuration loading")
	viper.SetConfigName("payment-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/placehub")
	viper.AddConfigPath("$HOME/.placehub")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			logger.Infow("configuration file not found")
			panic(err)
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}

	config := new(Config)
	viper.SetEnvPrefix("placehub")
	err := viper.BindEnv("id")
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config
}
