package branding

import (
	"fmt"

	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
	"github.com/uvulpos/golang-sveltekit-template/src/helper/branding/models"
)

func PrintBranding() {
	logo, logoErr := getAsciiLogo()
	if logoErr != nil {
		panic(logoErr)
	}

	startupInformation, _ := getAsciiStartupInformation()

	fmt.Println(logo)
	fmt.Println(startupInformation)
	fmt.Println("")
}

func PrintBrandingWithConfig() {
	logo, logoErr := getAsciiLogo()
	if logoErr != nil {
		panic(logoErr)
	}

	startupInformation, defaultConfConfiguraton := getAsciiStartupInformation()

	startupPresetTable := getAsciiConfigurationTable(models.ConfigurationTable{
		DisplayHost:  configuration.WEBSERVER_DISPLAYNAME,
		Host:         configuration.WEBSERVER_HOST,
		Port:         fmt.Sprint(configuration.WEBSERVER_PORT),
		ShowFrontend: configuration.WEBSERVER_SHOW_FRONTEND,
		ShowSwagger:  configuration.WEBSERVER_SHOW_SWAGGER,
	}, defaultConfConfiguraton)

	fmt.Println(logo)
	fmt.Println(startupInformation)
	fmt.Println(startupPresetTable)
	fmt.Println("")
}
