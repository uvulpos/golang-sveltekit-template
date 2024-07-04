package configuration

var (
	// WEBSERVER
	WEBSERVER_DISPLAYNAME   = getEnvOrDefaultString("WEBSERVER_DISPLAYHOST", "3000")
	WEBSERVER_HOST          = getEnvOrDefaultString("WEBSERVER_HOST", "3000")
	WEBSERVER_PORT          = getEnvOrDefaultInt("WEBSERVER_PORT", 3000)
	WEBSERVER_SHOW_FRONTEND = getEnvOrDefaultBool("WEBSERVER_SHOW_FRONTEND", true)
	WEBSERVER_SHOW_SWAGGER  = getEnvOrDefaultBool("WEBSERVER_SHOW_SWAGGER", false)

	// Database
	DATABASE_ADDR     = getEnvOrDefaultString("DATABASE_ADDR", "postgres")
	DATABASE_PORT     = getEnvOrDefaultInt("DATABASE_PORT", 5432)
	DATABASE_SSL      = getEnvOrDefaultBool("DATABASE_SSL", false)
	DATABASE_USERNAME = getEnvOrDefaultString("DATABASE_USERNAME", "postgres")
	DATABASE_PASSWORD = getEnvOrDefaultString("DATABASE_PASSWORD", "postgres")
	DATABASE_DATABASE = getEnvOrDefaultString("DATABASE_DATABASE", "postgres")

	// Security
	SECURITY_JWT_SECRET = getEnvOrDefaultString("SECURITY_JWT_SECRET", "loafofbread")
)
