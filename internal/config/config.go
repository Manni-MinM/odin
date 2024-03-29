package config

import (
	"github.com/spf13/viper"
)

type (
    Config struct {
        API             API             `mapstructure:"api"`
        HealthCheck     HealthCheck     `mapstructure:"healthcheck"`
    }

    API struct {
        Server          Server          `mapstructure:"server"`
        Redis           Redis           `mapstructure:"redis"`
    }

    Server struct {
        Port            int             `mapstructure:"port"`
    }

    Redis struct {
        Host            string          `mapstructure:"host"`
        Port            int             `mapstructure:"port"`
        DBName          int             `mapstructure:"dbname"`
        Password        string
    }

    HealthCheck struct {
        Redis           Redis           `mapstructure:"redis"`
        Cron            Cron            `mapstructure:"cron"`

    }

    Cron struct {
        Timeout         int             `mapstructure:"timeout"`
        Pattern         string          `mapstructure:"pattern"`
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

    viper.AutomaticEnv()

    conf.API.Redis.Password = viper.GetString("REDIS_PASSWORD")
    conf.HealthCheck.Redis.Password = viper.GetString("REDIS_PASSWORD")

	return &conf, nil
}
