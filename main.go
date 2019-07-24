package main

import (
	"fmt"

	"github.com/barolab/candidate/social"
	"github.com/barolab/candidate/twitter"
	"github.com/barolab/candidate/version"
)

func main() {
	fmt.Printf("Starting Candidate with version %s\n", version.Version())

	names := []string{
		"JeanMichelSuperRel0u",
		"this-is-great",
		"this-is-not-TwITTer",
		"Zoideberg",
	}

	providers := []social.Network{
		twitter.NewTwitter("twitter.com", "this-is-not-so-secret", "changeme"),
	}

	for _, name := range names {
		validate(name, providers)
	}
}

func validate(name string, providers []social.Network) {
	for _, provider := range providers {
		if violations := provider.Validate(name); !violations.IsNil() {
			fmt.Printf("Failed to validate \"%s\" on %s:\n%s\n", name, provider.Name(), violations)
		}
	}
}
