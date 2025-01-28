package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config struct to map configuration variables
type Config struct {
	AppPort              string `mapstructure:"APP_PORT"`
	AppEnv               string `mapstructure:"APP_ENV"`
	MongodbUser          string `mapstructure:"MONGODB_USER"`
	MongodbPassword      string `mapstructure:"MONGODB_PASSWORD"`
	MongodbHost          string `mapstructure:"MONGODB_HOST"`
	MongodbName          string `mapstructure:"MONGODB_NAME"`
	MongodbConnectionURI string `mapstructure:"MONGODB_CONNECTION_URI"`
}

// Global configuration instance
var AppConfig Config

// InitConfig initializes the configuration loader
func InitConfig(configPath, configFile, configType string) error {
	v := viper.New()

	// Set configuration file name, type, and path
	v.SetConfigName(configFile)
	v.SetConfigType(configType)
	v.AddConfigPath(configPath) // Add custom path
	v.AddConfigPath(".")        // Add current directory as a fallback

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	// Unmarshal configuration into the struct
	if err := v.Unmarshal(&AppConfig); err != nil {
		return err
	}

	log.Printf("Configuration loaded from: %s", filepath.Join(configPath, configFile))
	return nil
}

// GetConfig provides access to the loaded configuration
func GetConfig() Config {
	return AppConfig
}
