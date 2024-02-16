package branding

import (
	"github.com/charmbracelet/lipgloss"
)

func getAsciiStartupInformation() (string, lipgloss.Style) {
	var styleInformation = lipgloss.NewStyle().
		Width(100).
		Bold(true).
		Align(lipgloss.Center).
		Underline(true).
		MarginLeft(5).
		MarginRight(5)

	return styleInformation.Render("Example Application by @uvulpos"), styleInformation
}
