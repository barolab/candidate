package twitter_test

import (
	"testing"

	"github.com/barolab/candidate/twitter"
)

func TestConfig(T *testing.T) {
	client := twitter.NewTwitter(
		twitter.WithURL("url"),
		twitter.WithAPIKey("key"),
		twitter.WithSecretKey("secret"),
	)

	if client.URL() != "url" {
		T.Errorf("Failed to validate that WithURL correctly change the URL")
	}

	if client.APIKey() != "key" {
		T.Errorf("Failed to validate that WithAPIKey correctly change the API Key")
	}

	if client.APISecret() != "secret" {
		T.Errorf("Failed to validate that WithSecretKey correctly change the API Secret")
	}
}
