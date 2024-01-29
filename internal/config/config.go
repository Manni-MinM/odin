package config

import (
	"github.com/spf13/viper"
)

type (
    Config struct {
        Server      Server      `mapstructure:"server"`
        Redis       Redis       `mapstructure:"redis"`
    }

    Server struct {
        Port        int         `mapstructure:"port"`
    }

    Redis struct {
        Host        string      `mapstructure:"host"`
        Port        int         `mapstructure:"port"`
        Password    string      `mapstructure:"password"`
        DBName      string      `mapstructure:"dbname"`
    }
)

func Load() (*Config, error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return &Config{}, err
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		return &Config{}, err
	}

	return &conf, nil
}