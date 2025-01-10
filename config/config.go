package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                   string `mapstructure:"PORT"`
	Mongodb_user           string `mapstructure:"MONGODB_USER"`
	Mongodb_password       string `mapstructure:"MONGODB_PASSWORD"`
	Mongodb_host           string `mapstructure:"MONGODB_HOST"`
	Mongodb_name           string `mapstructure:"MONGODB_NAME"`
	Mongodb_connection_uri string `mapstructure:"MONGO_CONNECTION_URI"`
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(".env")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c Config
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return v, nil
}
