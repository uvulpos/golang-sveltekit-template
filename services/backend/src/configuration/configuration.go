package configuration

import (
	"fmt"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

var (
	Configuration ConfigurationModel
)

type ConfigurationModel struct {
	Webserver struct {
		Host         string `env:"WEBSERVER_HOST" default:"localhost" usage:"webserver host"`
		Port         int    `env:"WEBSERVER_PORT" default:"3000" usage:"webserver port"`
		ShowFrontend bool   `env:"SHOW_FRONTEND" default:"true" usage:"deliver static frontend by webserver"`
		ShowSwagger  bool   `env:"SHOW_SWAGGER" default:"true" usage:"expose or hide swagger ep + openapi definition"`
	}
	Session struct {
		JwtSecret string `env:"JWT_SECRET" default:"loafofbread" usage:"jwt signing key"`
	}
	Database struct {
		Addr string `env:"DB_ADDRESS" default:"postgres" usage:"postgres database address"`
		Port int    `env:"DB_PORT" default:"5432" usage:"postgres database port"`
		SSL  bool   `env:"DB_SSL" default:"false" usage:"postgres database ssl mode"`

		Username string `env:"DB_USERNAME" default:"postgres" usage:"postgres database username"`
		Password string `env:"DB_PASSWORD" default:"postgres" usage:"postgres database password"`
		Database string `env:"DB_DATABASENAME" default:"postgres" usage:"postgres database database name"`
	}
}

func LoadConfiguation() {
	var cfg ConfigurationModel
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{

		// SkipDefaults: true,
		SkipFiles: true,
		// SkipEnv:   true,
		SkipFlags: true,

		AllowUnknownEnvs: true,

		EnvPrefix:  "APP",
		FlagPrefix: "app",
		Files:      []string{"/var/opt/" + CONST_APPLICATION_PATH_NAME + "/config.yaml", "configuration.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
			".yml":  aconfigyaml.New(),
			".env":  aconfigdotenv.New(),
		},
	})

	if err := loader.Load(); err != nil {
		fmt.Println(">>", err.Error(), "<<")
		panic(err.Error())
	}

	Configuration = cfg
}
