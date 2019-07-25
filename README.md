# CANDIDATE

[![Build Status](https://drone.raimon.eu/api/badges/romain/candidate/status.svg)](https://drone.raimon.eu/romain/candidate)

Web API for checking if a name is available in multiple services.

## Getting Started

After getting the package using `go get github.com/barolab/candidate` you can get all supported social network like so :

```go
package main

import (
	"fmt"

	"github.com/barolab/candidate"
	_ "github.com/barolab/candidate/github"
	_ "github.com/barolab/candidate/instagram"
	_ "github.com/barolab/candidate/pinterest"
	_ "github.com/barolab/candidate/reddit"
	_ "github.com/barolab/candidate/twitter"
)

func main() {
	providers := candidate.SocialNetworks()

    for _, provider := range providers {
        violations := provider.Validate("Candidate")
        if !violations.isNil() {
            fmt.Printf("Candidate is NOT valid on %s, got %s\n", provider, violations)
        }
    }

    fmt.Println("Candidate is valid on all providers")
}
```

## Documentation

- [Changelog](/doc/CHANGELOG)
- [Code of Conduct](/doc/CODE_OF_CONDUCT.md)
- [Contributing](/doc/CONTRIBUTING.md)
