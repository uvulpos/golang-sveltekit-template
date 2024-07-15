package configuration

import (
	configurationHelper "github.com/uvulpos/go-svelte/basic-utils/configuration"
)

var (
	// WEBSERVER
	WEBSERVER_DISPLAYNAME   = configurationHelper.GetEnvOrDefaultString("WEBSERVER_DISPLAYHOST", "http://web.localhost/")
	WEBSERVER_HOST          = configurationHelper.GetEnvOrDefaultString("WEBSERVER_HOST", "127.0.0.1")
	WEBSERVER_PORT          = configurationHelper.GetEnvOrDefaultInt("WEBSERVER_PORT", 443)
	WEBSERVER_SHOW_FRONTEND = configurationHelper.GetEnvOrDefaultBool("WEBSERVER_SHOW_FRONTEND", true)
	WEBSERVER_SHOW_SWAGGER  = configurationHelper.GetEnvOrDefaultBool("WEBSERVER_SHOW_SWAGGER", true)

	// Database
	DATABASE_ADDR = configurationHelper.GetEnvOrDefaultString("DATABASE_ADDR", "127.0.0.1")
	// DATABASE_ADDR     = configurationHelper.GetEnvOrDefaultString("DATABASE_ADDR", "postgres")
	DATABASE_PORT     = configurationHelper.GetEnvOrDefaultInt("DATABASE_PORT", 5432)
	DATABASE_SSL      = configurationHelper.GetEnvOrDefaultBool("DATABASE_SSL", false)
	DATABASE_USERNAME = configurationHelper.GetEnvOrDefaultString("DATABASE_USERNAME", "postgres")
	DATABASE_PASSWORD = configurationHelper.GetEnvOrDefaultString("DATABASE_PASSWORD", "postgres")
	DATABASE_DATABASE = configurationHelper.GetEnvOrDefaultString("DATABASE_DATABASE", "postgres")

	// Oauth
	AUTHORIZATION_OAUTH_KEY          = configurationHelper.GetEnvOrDefaultString("AUTHORIZATION_OAUTH_KEY", "key123")
	AUTHORIZATION_OAUTH_SECRET       = configurationHelper.GetEnvOrDefaultString("AUTHORIZATION_OAUTH_SECRET", "oh-oh-ganz-geheim")
	AUTHORIZATION_OAUTH_CALLBACK_URL = configurationHelper.GetEnvOrDefaultString("AUTHORIZATION_OAUTH_CALLBACK_URL", "http://web.localhost/api/v1/callback")

	// Certificates
	CERTIFICATE_IDENTITY_NAME                        = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_NAME", "uVulpos - My Application")
	CERTIFICATE_IDENTITY_ORGANIZATION                = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_ORGANIZATION", "The Innovators")
	CERTIFICATE_IDENTITY_ORGANIZATION_UNIT           = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_ORGANIZATION_UNIT", "Technical Innovation")
	CERTIFICATE_IDENTITY_ORGANIZATION_ADRESS         = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_ORGANIZATION_ADRESS", "Example Street 12a")
	CERTIFICATE_IDENTITY_ORGANIZATION_CITY           = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_ORGANIZATION_CITY", "Munich")
	CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_CODE    = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_CODE", "12345")
	CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_STATE   = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_STATE", "Bavaria")
	CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_COUNTRY = configurationHelper.GetEnvOrDefaultString("CERTIFICATE_IDENTITY_ORGANIZATION_POSTAL_COUNTRY", "Germany")
)
