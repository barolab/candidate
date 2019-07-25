package main

import (
	"fmt"
	"os"

	"github.com/barolab/candidate"
	_ "github.com/barolab/candidate/github"
	_ "github.com/barolab/candidate/reddit"
	_ "github.com/barolab/candidate/twitter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("You need to pass a username to validate (go run main.go candidate)")
		os.Exit(1)
	}

	name := os.Args[1]
	providers := candidate.SocialNetworks()

	for _, p := range providers {
		fmt.Println(check(name, p))
	}
}

func check(name string, provider candidate.NameProvider) string {
	violations := provider.Validate(name)
	if !violations.IsNil() {
		return fmt.Sprintf("\"%s\" is not valid for %s:\n%s", name, provider, violations)
	}

	ok, err := provider.IsAvailable(name)
	if err != nil {
		return fmt.Sprintf("\"%s\" is valid for %s, but we got an error checking if it already exist : %s", name, provider, err)
	}

	if !ok {
		return fmt.Sprintf("\"%s\" is valid for %s, but it's already taken ¯\\_(ツ)_/¯", name, provider)
	}

	return fmt.Sprintf("\"%s\" is valid and available on %s", name, provider)
}
