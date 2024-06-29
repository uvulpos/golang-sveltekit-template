package branding

import (
	"fmt"

	"github.com/uvulpos/go-svelte/src/configuration"
	"github.com/uvulpos/go-svelte/src/helper/branding/models"
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
		Host:         configuration.Configuration.Webserver.Host,
		Port:         fmt.Sprint(configuration.Configuration.Webserver.Port),
		ShowFrontend: configuration.Configuration.Webserver.ShowFrontend,
		ShowSwagger:  configuration.Configuration.Webserver.ShowSwagger,
	}, defaultConfConfiguraton)

	fmt.Println(logo)
	fmt.Println(startupInformation)
	fmt.Println(startupPresetTable)
	fmt.Println("")
}
