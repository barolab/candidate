package url_test

import (
	"fmt"
	"testing"

	"github.com/barolab/candidate/url"
)

type WithoutQueryTestCase struct {
	argument string
	expected string
	err      error
}

func TestWithoutQuery(T *testing.T) {
	cases := []WithoutQueryTestCase{
		{argument: "", expected: "", err: fmt.Errorf("Cannot exclude URL query from an empty URL")},
		{argument: "https://twitter.com/candidate", expected: "https://twitter.com/candidate", err: nil},
		{argument: "https://twitter.com/candidate?this=that", expected: "https://twitter.com/candidate", err: nil},
	}

	for _, c := range cases {
		res, err := url.WithoutQuery(c.argument)
		if res != c.expected {
			T.Errorf("WithoutQuery should have return %s for url %s, got %s", c.expected, c.argument, res)
		}

		if err == nil && c.err != nil {
			T.Errorf("WithoutQuery returned no error but we expected to return %s (for url %s)", c.err, c.argument)
		}

		if err != nil && c.err == nil {
			T.Errorf("WithoutQuery returned an error %s that was not expected (for url %s)", err, c.argument)
		}

		if err != nil && c.err != nil && err.Error() != c.err.Error() {
			T.Errorf("WithoutQuery should have returned an error %s, but we got %s (for url %s)", c.err, err, c.argument)
		}
	}
}
