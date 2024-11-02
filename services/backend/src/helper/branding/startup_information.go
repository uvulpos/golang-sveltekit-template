package branding

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/uvulpos/golang-sveltekit-template/src/configuration"
)

func getAsciiStartupInformation() (string, lipgloss.Style) {
	var styleInformation = lipgloss.NewStyle().
		Width(100).
		Bold(true).
		Align(lipgloss.Center).
		Underline(true).
		MarginLeft(5).
		MarginRight(5)

	return styleInformation.Render(configuration.CONST_APPLICATION_BRANDING_HEADER), styleInformation
}
