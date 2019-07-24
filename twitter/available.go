package twitter

import (
	"fmt"
	"net/http"
	"net/url"

	curl "github.com/barolab/candidate/url"
)

// IsAvailable check if the given name is available in Twitter
func (t *Twitter) IsAvailable(name string) (bool, error) {
	u, err := curl.WithoutQuery(fmt.Sprintf("https://twitter.com/%s", url.QueryEscape(name)))
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false, fmt.Errorf("Failed to build HTTP request for URL %s, got error: %s", u, err)
	}

	res, err := t.client.Fetch(req)
	if err != nil {
		return false, fmt.Errorf("Failed to contact Twitter at %s with error %s", u, err)
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusNotFound, nil
}
