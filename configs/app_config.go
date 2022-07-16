package configs

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	Port                       string `mapstructure:"PORT"`
	PostgresqlConnectionString string `mapstructure:"POSTGRESQL_CONNECTION_STRING"`
	RabbitmqConnectionString   string `mapstructure:"RABBITMQ_CONNECTION_STRING"`
}

func LoadConfig() {
	log.Println("Loading config...")

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}

}
