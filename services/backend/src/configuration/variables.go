package configuration

import (
	configurationHelper "github.com/uvulpos/go-svelte/basic-utils/configuration"
)

var (
	// WEBSERVER
	WEBSERVER_DISPLAYNAME   = configurationHelper.GetEnvOrDefaultString("WEBSERVER_DISPLAYHOST", "http://web.localhost/")
	WEBSERVER_HOST          = configurationHelper.GetEnvOrDefaultString("WEBSERVER_HOST", "127.0.0.1")
	WEBSERVER_PORT          = configurationHelper.GetEnvOrDefaultInt("WEBSERVER_PORT", 3000)
	WEBSERVER_SHOW_FRONTEND = configurationHelper.GetEnvOrDefaultBool("WEBSERVER_SHOW_FRONTEND", true)
	WEBSERVER_SHOW_SWAGGER  = configurationHelper.GetEnvOrDefaultBool("WEBSERVER_SHOW_SWAGGER", true)

	// Database
	DATABASE_ADDR     = configurationHelper.GetEnvOrDefaultString("DATABASE_ADDR", "postgres")
	DATABASE_PORT     = configurationHelper.GetEnvOrDefaultInt("DATABASE_PORT", 5432)
	DATABASE_SSL      = configurationHelper.GetEnvOrDefaultBool("DATABASE_SSL", false)
	DATABASE_USERNAME = configurationHelper.GetEnvOrDefaultString("DATABASE_USERNAME", "postgres")
	DATABASE_PASSWORD = configurationHelper.GetEnvOrDefaultString("DATABASE_PASSWORD", "postgres")
	DATABASE_DATABASE = configurationHelper.GetEnvOrDefaultString("DATABASE_DATABASE", "postgres")

	// Security
	SECURITY_JWT_SECRET = configurationHelper.GetEnvOrDefaultString("SECURITY_JWT_SECRET", "loafofbread")
)
