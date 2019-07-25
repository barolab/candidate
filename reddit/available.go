package reddit

import (
	"fmt"
	"net/http"
	"net/url"

	curl "github.com/barolab/candidate/url"
)

// IsAvailable check if the given name is available in Twitter
func (r *Reddit) IsAvailable(name string) (bool, error) {
	u, err := curl.WithoutQuery(fmt.Sprintf("https://www.reddit.com/user/%s", url.QueryEscape(name)))
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return false, fmt.Errorf("Failed to build HTTP request for URL %s, got error: %s", u, err)
	}

	res, err := r.client.Fetch(req)
	if err != nil {
		return false, fmt.Errorf("Failed to contact %s at %s with error %s", r, u, err)
	}

	defer res.Body.Close()
	return res.StatusCode == http.StatusNotFound, nil
}
