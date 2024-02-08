package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadData() *Configuration {
	fmt.Println("---- LOAD CONFIG")

	viperInstance := viper.New()

	viperInstance.SetDefault("db-host", "postgres")
	viperInstance.SetDefault("db-port", "5432")
	viperInstance.SetDefault("db-username", "postgres")
	viperInstance.SetDefault("db-password", "postgres")
	viperInstance.SetDefault("db-database", "postgres")
	viperInstance.SetDefault("db-sslmode", false)

	viperInstance.AutomaticEnv()

	return &Configuration{
		Webserver: Webserver{
			Port:        3000,
			NoFrontend:  false,
			ShowSwagger: false,
		},
		DB: DatabaseCredentials{
			Host:     "postgres",
			Port:     5432,
			Username: "postgres",
			Password: "postgres",
			Database: "postgres",
			SslMode:  false,
		},
	}
}
