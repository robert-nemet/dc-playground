package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort    string `mapstructure:"APP_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbName     string `mapstructure:"DB_NAME"`
	DbType     string `mapstructure:"DB_TYPE"`
}

func LoadConfig() (config AppConfig, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	viper.BindEnv("DB_HOST", "DB_HOST")
	viper.BindEnv("DB_PORT", "DB_PORT")
	viper.BindEnv("DB_USER", "DB_USER")
	viper.BindEnv("DB_PASSWORD", "DB_PASSWORD")
	viper.BindEnv("DB_NAME", "DB_NAME")
	viper.BindEnv("DB_TYPE", "DB_TYPE")
	viper.BindEnv("APP_PORT", "APP_PORT")
	viper.SetDefault("DB_TYPE", "PG")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}

	err = viper.Unmarshal(&config)
	return
}
