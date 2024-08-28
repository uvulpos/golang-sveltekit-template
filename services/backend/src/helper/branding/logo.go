package branding

import (
	"image"

	"github.com/charmbracelet/lipgloss"
	"github.com/qeesung/image2ascii/convert"
	"github.com/uvulpos/golang-sveltekit-template/src/assets"
)

func getAsciiLogo() (string, error) {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 80
	convertOptions.FixedHeight = 35

	// Create the image converter
	converter := convert.NewImageConverter()

	// Open the image file
	file, err := assets.TerminalFS.Open("terminal-assets/gopher.png")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Decode the image file
	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	// Convert the image to an ASCII matrix and print it
	asciiImage := converter.Image2ASCIIString(img, &convertOptions)

	var styleImage = lipgloss.NewStyle().
		Width(100).
		Align(lipgloss.Center).
		MarginLeft(5).
		MarginRight(5)

	return styleImage.Render(asciiImage), nil
}
