package updater

import (
	"fmt"
	"net/http"
)

type GitHubReleases []GitHubRelease
type GitHubRelease struct {
	ID              string
	Name            string
	Body            string
	URL             string
	HtmlURL         string
	ZipballURL      string
	TarballURL      string
	TargetCommitish string
	draft           bool
	prerelease      bool
}

func GitHubUpdateApp(username, reponame string) {
	releasesUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", username, reponame)

	_, respErr := http.Get(releasesUrl)
	if respErr != nil {
		panic(respErr)
	}
}
