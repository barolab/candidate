package fetch

import (
	"fmt"
	"net/http"
	"time"

	"github.com/barolab/candidate/url"
)

// Fetcher is an interface helper that is used to execute an http request
type Fetcher interface {
	// Fetch perform an HTTP request and return the response, with error if any
	Fetch(req *http.Request) (*http.Response, error)
}

type nativeFetcher struct {
	client *http.Client
}

func (f *nativeFetcher) Fetch(req *http.Request) (*http.Response, error) {
	return f.client.Do(req)
}

// DefaultFetcher is the default mechanism for performing HTTP request in this module
var DefaultFetcher = &nativeFetcher{
	client: &http.Client{
		Timeout: time.Second * 5,
	},
}

// IsNotFound check of the given URL exists on the internet (no query string accepted)
func IsNotFound(f Fetcher, u string) (result bool, err error) {
	url, err := url.WithoutQuery(u)
	if err != nil {
		return false, fmt.Errorf("Cannot create URL for %s: %s", u, err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, fmt.Errorf("Cannot build request to %s: %s", url, err)
	}

	req.Header.Add("Accept", "text/html")
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "en-US")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:12.0) Gecko/20100101 Firefox/12.0")

	res, err := f.Fetch(req)
	if err != nil {
		return false, fmt.Errorf("Request to %s failed: %s", url, err)
	}

	defer func() {
		if e := res.Body.Close(); e != nil {
			err = fmt.Errorf("Fail to close HTTP response: %s", e)
		}
	}()

	return res.StatusCode == http.StatusNotFound, nil
}
