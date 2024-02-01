package db

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func NewConfig() *Config {
	config := Config{}
	viper.SetConfigFile("configs/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file config.yaml : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	log.Println(config)
	return &config
}
