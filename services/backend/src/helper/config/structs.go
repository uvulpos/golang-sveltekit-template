package config

type Configuration struct {
	Webserver Webserver           `mapstructure:"webserver"`
	DB        DatabaseCredentials `mapstructure:"db"`
}

type Webserver struct {
	Port        string `mapstructure:"port"`
	NoFrontend  string `mapstructure:"nofrontend"`
	ShowSwagger string `mapstructure:"showswagger"`
}

type DatabaseCredentials struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
