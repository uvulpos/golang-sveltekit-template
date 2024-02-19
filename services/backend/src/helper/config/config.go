package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadData() *Configuration {

	config := createDefaultConfiguration()
	fmt.Println("---- LOAD CONFIG")

	viperInstance := viper.New()

	viperInstance.SetDefault("db-host", "postgres")
	viperInstance.SetDefault("db-port", "5432")
	viperInstance.SetDefault("db-username", "postgres")
	viperInstance.SetDefault("db-password", "postgres")
	viperInstance.SetDefault("db-database", "postgres")
	viperInstance.SetDefault("db-sslmode", false)

	viperInstance.AutomaticEnv()

	return config
}
