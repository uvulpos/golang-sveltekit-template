package config

type Configuration struct {
	Webserver Webserver           `mapstructure:"webserver"`
	DB        DatabaseCredentials `mapstructure:"db"`
}

type Webserver struct {
	Port        int  `mapstructure:"port"`
	NoFrontend  bool `mapstructure:"nofrontend"`
	ShowSwagger bool `mapstructure:"showswagger"`
}

type DatabaseCredentials struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SslMode  bool   `mapstructure:"sslmode"`
}

func createDefaultConfiguration() *Configuration {
	return &Configuration{
		Webserver: Webserver{
			Port:        3000,
			NoFrontend:  false,
			ShowSwagger: true,
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
