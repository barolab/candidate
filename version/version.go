package version

import "fmt"

// VERSION is the app-global version string, which should be substituted with a real value during build.
var VERSION = "UNKNOWN"

// GOVERSION is the Golang version used to generate the binary.
var GOVERSION = "UNKNOWN"

// BUILDTIME is the timestamp at which the binary was created.
var BUILDTIME = "UNKNOWN"

// COMMITHASH is the git commit hash that was used to generate the binary.
var COMMITHASH = "UNKNOWN"

// Version string containing the version number, the commit hash, the golang version & the build time
func Version() string {
	return fmt.Sprintf("%s (commit %s built with go %s the %s)", VERSION, COMMITHASH, GOVERSION, BUILDTIME)
}
