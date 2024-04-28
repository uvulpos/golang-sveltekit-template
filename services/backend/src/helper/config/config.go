package config

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

var (
	config Config
)

/*
*
*	I encountered issues by autodetecting the struct from viper so temporary manual
*   TODO: autoread struct
*
 */
func LoadData() {

	viperConfig := viper.New()
	viperConfig.AddConfigPath(".")
	viperConfig.AddConfigPath("$HOME/.my-application")
	viperConfig.SetConfigName("config")
	viperConfig.SetConfigType("yaml")

	viperConfig.SetEnvPrefix("app")

	replacer := strings.NewReplacer(".", "_")
	viperConfig.SetEnvKeyReplacer(replacer)
	viperConfig.AllowEmptyEnv(true)

	viperConfig.AutomaticEnv()
	bindEnvVariables(reflect.TypeOf(Config{}), "", viperConfig)

	err := viperConfig.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("[INFORMATION]: No config file found, relying on environment variables only.")
		} else {
			log.Fatalf("Fatal error config file: %s", err)
		}
	}

	if err := viperConfig.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode into struct, %s", err)
	}

	fmt.Println("Keys: ", viperConfig.AllKeys())
	fmt.Println("Keys: ", viperConfig.AllSettings())
	fmt.Println("Webserver: ", config.Webserver)
	fmt.Println("Database: ", config.Database)
	fmt.Println("Security: ", config.Security)
}

func bindEnvVariables(t reflect.Type, parent string, viperConfig *viper.Viper) {
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("mapstructure")
			if jsonTag == "" {
				jsonTag = strings.ToLower(field.Name)
			}

			fullPath := jsonTag
			if parent != "" {
				fullPath = parent + "." + jsonTag
			}

			viperConfig.BindEnv(fullPath, strings.ToUpper(viperConfig.GetEnvPrefix()+"_"+strings.ReplaceAll(fullPath, ".", "_")))
			bindEnvVariables(field.Type, fullPath, viperConfig)
		}
	}
}

func GetSqlConnectionString() (string, string) {

	var sslMode string = "disable"
	if config.Database.SslMode {

		sslMode = "require"
	}

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
		sslMode,
	)

	humanString := fmt.Sprintf(
		"%s:%d",
		config.Database.Host,
		config.Database.Port,
	)

	return connString, humanString
}

func GetDatabase() DatabaseConfig {
	return config.Database
}

func GetWebserver() WebserverConfig {
	return config.Webserver
}

func ShowFrontend() bool {
	return config.Webserver.ShowFrontend
}

func GetJwtSecret() string {
	return config.Security.JwtSecret
}
