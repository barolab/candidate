package main

import (
	"fmt"

	"github.com/barolab/candidate/social"
	"github.com/barolab/candidate/version"
)

func main() {
	fmt.Printf("Starting Candidate with version %s\n", version.Version())

	providers := []social.Network{
		social.NewTwitter("twitter.com", "this-is-not-so-secret", "changeme"),
	}

	for _, provider := range providers {
		if err := provider.Validate("JeanMichelSuperRel0u"); err != nil {
			fmt.Printf("Failed to validate username on %s, got %s\n", provider.Name(), err)
		}
	}
}
