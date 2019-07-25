package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/barolab/candidate"
	_ "github.com/barolab/candidate/github"
	_ "github.com/barolab/candidate/pinterest"
	_ "github.com/barolab/candidate/reddit"
	_ "github.com/barolab/candidate/twitter"
)

// main will parse flags and decide how to call validation process
func main() {
	var (
		start     = time.Now()
		name      = flag.String("name", "", "The name to check")
		parrallel = flag.Bool("parrallel", false, "Enable parrallel processing of the providers")
		providers = candidate.SocialNetworks()
	)

	flag.Parse()
	if *name == "" {
		fmt.Printf("You need to pass a username to validate (go run main.go -name=candidate)")
		os.Exit(1)
	}

	if *parrallel == true {
		async(*name, providers)
	} else {
		procedural(*name, providers)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done in: %v\n", elapsed)
}

// procedural processsing will check for all providers one at a time
func procedural(name string, providers []candidate.NameProvider) {
	for _, provider := range providers {
		fmt.Println(check(name, provider))
	}
}

// async processing will check for providers in parrallel
func async(name string, providers []candidate.NameProvider) {
	ch := make(chan string, len(providers))
	defer close(ch)

	for _, provider := range providers {
		go func(p candidate.NameProvider) {
			ch <- check(name, p)
		}(provider)
	}

	for range providers {
		fmt.Println(<-ch)
	}
}

// check a name against the given provider, the return message can be an error, violations or success
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
