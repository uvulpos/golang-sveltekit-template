package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadData() {
	fmt.Println("---- LOAD CONFIG")

	viperInstance := viper.New()

	viperInstance.SetDefault("host", "postgres")
	viperInstance.SetDefault("port", "5432")
	viperInstance.SetDefault("username", "postgres")
	viperInstance.SetDefault("password", "postgres")
	viperInstance.SetDefault("database", "postgres")

	viperInstance.AutomaticEnv()
}
