package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct to hold all config sections
type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Port int
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

// LoadConfig loads the configuration from file and environment variables
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yaml")   // Required if config file has no extension
	viper.AddConfigPath(".")      // Look for config in the current directory

	// Allow environment variable overrides
	viper.AutomaticEnv()

	// Optional: map env var names to config fields
	viper.SetEnvPrefix("APP") // e.g. APP_PORT overrides app.port
	viper.BindEnv("app.port", "APP_PORT")
	viper.BindEnv("database.host", "PGHOST")
	viper.BindEnv("database.port", "PGPORT")
	viper.BindEnv("database.user", "PGUSER")
	viper.BindEnv("database.password", "PGPASSWORD")
	viper.BindEnv("database.name", "PGDATABASE")

	// Load the config file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found, using environment variables only.")
	}

	var config Config
	// Unmarshal the config into the Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &config, nil
}
