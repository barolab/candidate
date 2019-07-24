package url

import (
	"fmt"
	"net/url"
)

// WithoutQuery will parse the given url string and will return a version without the query string
func WithoutQuery(u string) (string, error) {
	if len(u) == 0 {
		return "", fmt.Errorf("Cannot exclude URL query from an empty URL")
	}

	url, err := url.Parse(u)
	if err != nil {
		return "", fmt.Errorf("Failed to parse URL %s with error %s", u, err)
	}

	return fmt.Sprintf("%s://%s%s", url.Scheme, url.Host, url.Path), nil
}
