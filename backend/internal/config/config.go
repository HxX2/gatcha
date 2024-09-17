package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port string `mapstructure:"PORT"`
	Host string `mapstructure:"HOST"`
}

type DatabaseConfig struct {
	Name     string `mapstructure:"DB_NAME"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASS"`
}

func Load() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var sc ServerConfig
	err := viper.Unmarshal(&sc)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	var dc DatabaseConfig
	err = viper.Unmarshal(&dc)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	var cfg Config
	cfg.Server = sc
	cfg.Database = dc

	return &cfg
}
