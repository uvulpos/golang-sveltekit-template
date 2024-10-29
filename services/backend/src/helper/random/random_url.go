package random

import (
	"fmt"
	"net/url"

	"github.com/uvulpos/golang-sveltekit-template/src/helper/customerrors"
)

func GenerateRandomHashURL(baseURL string) (string, customerrors.ErrorInterface) {
	hash, err := generateRandomHash()
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", "Could not generate a simple oauth redirect get parameter hash")
	}

	// URL parsen
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return "", customerrors.NewInternalServerError(err, "", fmt.Sprintf("Couldn't parse URL in oauth callback (func: generateRandomHashURL ; data: %v)", baseURL))
	}

	query := parsedURL.Query()
	query.Set("refresh-hash", hash)
	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}
