package twitter_test

import (
	"testing"

	"github.com/barolab/candidate/twitter"
)

func TestName(T *testing.T) {
	client := twitter.NewTwitter()
	if client.Name() != "Twitter" {
		T.Errorf("Failed to validate client name")
	}
}
