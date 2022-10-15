package configs

import (
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	Port                       string `mapstructure:"SERVER_PORT"`
	PostgresqlConnectionString string `mapstructure:"POSTGRESQL_CONNECTION_STRING"`
	RabbitmqConnectionString   string `mapstructure:"RABBITMQ_CONNECTION_STRING"`

	ApiPathPrefix string `mapstructure:"API_PATH_PREFIX"`

	UserServiceUrl           string `mapstructure:"USER_SERVICE_URL"`
	UserServiceApiPathPrefix string `mapstructure:"USER_SERVICE_API_PATH_PREFIX"`
}

func LoadConfig() {
	log.Println("Loading system env variables...")

	/*
	 * Since I couldn't find any better lib alternative, before reading env values from file read them from sys env to override file values
	 */
	var appConfigTemp Config
	appConfigTempFields := reflect.ValueOf(appConfigTemp)
	for i := 0; i < appConfigTempFields.Type().NumField(); i++ {
		envName := appConfigTempFields.Type().Field(i).Tag.Get("mapstructure")
		envValue := os.Getenv(envName)
		log.Printf("%s=%s\n", envName, envValue)
		reflect.ValueOf(&appConfigTemp).Elem().Field(i).SetString(envValue)
	}

	log.Printf("Loading file env variables...\n\n")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error while reading config file %s\n", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Println(err)
	}

	/*
	* Replace file env variables with system.
	 */
	log.Println("Final env variables. (System overrides)")
	for i := 0; i < appConfigTempFields.Type().NumField(); i++ {
		envName := appConfigTempFields.Type().Field(i).Tag.Get("mapstructure")
		envValue := reflect.ValueOf(appConfigTemp).Field(i).String()
		if envValue != "" {
			reflect.ValueOf(AppConfig).Elem().Field(i).SetString(envValue)
		} else {
			envValue = reflect.ValueOf(*AppConfig).Field(i).String()
		}

		if envValue == "" {
			log.Fatalf("Environment value %s is not defined.", envName)
		}

		log.Printf("%s=%s\n", envName, envValue)
	}

}
