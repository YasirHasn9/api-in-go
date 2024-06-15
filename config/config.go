// the goal of this package is to manage the configuration settings
// im using `viber` package to manage the configuration.
// This file will load the necessary settings.

// In Go language, config is a package
package config

import (
	"log"

	"github.com/spf13/viper"
)

// I think this like an interface in ts or asp.net
// oh no
// it is data type that groups together variables under one name
// structs are rigidly typed, they provide a fixed blueprint for data
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// *Config means the LoadConfig function returns a pointer to a Config
func LoadConfig() *Config {

	// tell viber to look for the file .env
	viper.SetConfigFile(".env")

	// if there is an error loading the file, then logs the error
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading the config file: %v", err)
	}

	// Note
	// In go the := is a shortcut for declaring and initializing a var
	// and then Go uses the value on the right to determine the type
	config := &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
	}

	return config
}
