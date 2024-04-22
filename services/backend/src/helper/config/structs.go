package config

type Config struct {
	Database  DatabaseConfig
	Webserver WebserverConfig
	Security  SecurityConfig
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	SslMode  bool   `mapstructure:"sslmode"`
}

type WebserverConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ShowFrontend bool   `mapstructure:"show_frontend"`
	ShowSwagger  bool   `mapstructure:"show_swagger"`
}

type SecurityConfig struct {
	JwtSecret   string
	SecretGreet string
}
