package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Setup initialize configuration
func Setup(configPath string, logger *zap.Logger) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// read in from env vars
	viper.AutomaticEnv()

	// read in from config file
	if err := viper.ReadInConfig(); err != nil {
		logger.Warn("Service configuration file at " + configPath + " not found. Using environment variables.")
	}

	// set default configs
	viper.SetDefault("server_port", 3000)
	viper.SetDefault("database_host", "localhost")
	viper.SetDefault("database_port", 5432)
	viper.SetDefault("database_driver", "postgres")
	viper.SetDefault("database_max_lifetime", 7200)
	viper.SetDefault("database_open_conns", 150)
	viper.SetDefault("database_idle_conns", 50)
}
