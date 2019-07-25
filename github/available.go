package github

import (
	"fmt"
	"net/http"
	"net/url"

	curl "github.com/barolab/candidate/url"
)

// IsAvailable check if the given name is available in Github
func (g *Github) IsAvailable(name string) (bool, error) {
	u, err := curl.WithoutQuery(fmt.Sprintf("https://github.com/%s", url.QueryEscape(name)))
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false, fmt.Errorf("Failed to build HTTP request for URL %s, got error: %s", u, err)
	}

	res, err := g.client.Fetch(req)
	if err != nil {
		return false, fmt.Errorf("Failed to contact %s at %s with error %s", g, u, err)
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusNotFound, nil
}
