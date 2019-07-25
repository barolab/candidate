package fetch_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/barolab/candidate/fetch"
)

type IsAvailableTestCase struct {
	argument string
	expected bool
	err      error
	fetcher  *FakeFetcher
}

type FakeFetcher struct {
	res *http.Response
	err error
}

func (f *FakeFetcher) Fetch(req *http.Request) (*http.Response, error) {
	if f.err != nil {
		return nil, f.err
	}

	return f.res, nil
}

var (
	OkFetcher = &FakeFetcher{
		err: nil,
		res: &http.Response{
			Status:     "200 Ok",
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Testing")),
		},
	}

	NotFoundFetcher = &FakeFetcher{
		err: nil,
		res: &http.Response{
			Status:     "404 Not Found",
			StatusCode: http.StatusNotFound,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Testing")),
		},
	}

	ErrorFetcher = &FakeFetcher{
		err: fmt.Errorf("Test"),
		res: nil,
	}
)

func TestIsNotFound(T *testing.T) {
	cases := []IsAvailableTestCase{
		{argument: "https://provider.com/candidate", expected: false, err: nil, fetcher: OkFetcher},
		{argument: "https://provider.com/candidate", expected: true, err: nil, fetcher: NotFoundFetcher},
		{argument: "https://provider.com/candidate", expected: false, err: fmt.Errorf("Request to https://provider.com/candidate failed: %s", ErrorFetcher.err), fetcher: ErrorFetcher},
	}

	for _, c := range cases {
		ok, err := fetch.IsNotFound(c.fetcher, c.argument)

		if ok != c.expected {
			T.Errorf("IsNotFound should have return %v, got %v for username %s", c.expected, ok, c.argument)
		}

		if err != nil && c.err == nil {
			T.Errorf("IsNotFound returned unexpected error %s for username %s", err, c.argument)
		}

		if err == nil && c.err != nil {
			T.Errorf("IsNotFound should have return error %s, got no error instead for username %s", c.err, c.argument)
		}

		if err != nil && c.err != nil && err.Error() != c.err.Error() {
			T.Errorf("IsNotFound should have return error %s, got %s for username %s", c.err, err, c.argument)
		}
	}
}
