package main

import "fmt"

// VERSION is the app-global version string, which should be substituted with a real value during build.
var VERSION = "UNKNOWN"

// GOVERSION is the Golang version used to generate the binary.
var GOVERSION = "UNKNOWN"

// BUILDTIME is the timestamp at which the binary was created.
var BUILDTIME = "UNKNOWN"

// COMMITHASH is the git commit hash that was used to generate the binary.
var COMMITHASH = "UNKNOWN"

func main() {
	fmt.Printf("Samples version %s (commit %s built with go %s the %s)\n", VERSION, COMMITHASH, GOVERSION, BUILDTIME)
}
