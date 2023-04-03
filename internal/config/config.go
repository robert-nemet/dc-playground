package config

import (
	"github.com/spf13/viper"
)

// TODO add topic to config and to docker-compose

type AppConfig struct {
	AppPort     string `mapstructure:"APP_PORT"`
	DbHost      string `mapstructure:"DB_HOST"`
	DbPort      string `mapstructure:"DB_PORT"`
	DbUser      string `mapstructure:"DB_USER"`
	DbPassword  string `mapstructure:"DB_PASSWORD"`
	DbName      string `mapstructure:"DB_NAME"`
	DbType      string `mapstructure:"DB_TYPE"`
	Target      string `mapstructure:"TARGET"`
	ErrorRate   int    `mapstructure:"ERROR_RATE"`
	KafkaBroker string `mapstructure:"KAFKA_BROKER"`
	KafkaTopic  string `mapstructure:"KAFKA_TOPIC"`
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
	viper.BindEnv("TARGET", "TARGET")
	viper.BindEnv("ERROR_RATE", "ERROR_RATE")
	viper.BindEnv("KAFKA_BROKER", "KAFKA_BROKER")
	viper.SetDefault("ERROR_RATE", 10)
	viper.SetDefault("DB_TYPE", "PG")
	viper.SetDefault("KAFKA_BROKER", "none")
	viper.SetDefault("KAFKA_TOPIC", "test-topic")
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
