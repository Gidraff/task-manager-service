package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config exported
type Config struct {
	database struct {
		DBUser string `mapstructure:"dbuser"`
	}
	ServerPort int    `mapstructure:"server_port" env:"SERVER_PORT"`
	DSN        string `mapstructure:"dsn" env:"DSN"`
}

// LoadConfig load configuration from file
func LoadConfig(configpath string) *viper.Viper {
	var conf Config
	viperCfg := viper.New()
	viperCfg.SetConfigName("dev")
	viperCfg.AddConfigPath(configpath)
	viperCfg.AutomaticEnv()
	viperCfg.SetConfigType("yml")
	// configPath := "${path}/config/config.yml"

	if err := viperCfg.ReadInConfig(); err != nil {
		fmt.Printf("Failed to read the config file %s", err)
	}

	if err := viperCfg.Unmarshal(&conf); err != nil {
		fmt.Printf("Failed to read the config file %s", err)
	}

	viperCfg.WatchConfig()

	return viperCfg
}
