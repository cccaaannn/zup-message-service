package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	Port             string `mapstructure:"PORT"`
	ConnectionString string `mapstructure:"POSTGRESQL_CONNECTION_STRING"`
}

func LoadConfig() {
	log.Println("Loading Server Configurations...")

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(AppConfig.ConnectionString)
}
