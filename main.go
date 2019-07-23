package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

var (
	illegalPatternRegexp = regexp.MustCompile("(?i)twitter")
	legalRunesRegexp     = regexp.MustCompile("^[0-9A-Za-z_]+$")
)

func isShortEnough(username string) bool {
	return utf8.RuneCountInString(username) < 15
}

func isLongEnough(username string) bool {
	return utf8.RuneCountInString(username) > 1
}

func containsNoIllegalPattern(username string) bool {
	return illegalPatternRegexp.MatchString(username)
}

func onlyContainsLegalRunes(username string) bool {
	return legalRunesRegexp.MatchString(username)
}

func main() {
	names := []string{
		"JeanMichelSuperRelou",
		"",
		"tWItter",
		"Bad-Candidate",
		"Candidate",
	}

	for _, name := range names {
		if !isShortEnough(name) {
			fmt.Printf("Username \"%s\" is too long\n", name)
		}

		if !isLongEnough(name) {
			fmt.Printf("Username \"%s\" is too short\n", name)
		}

		if containsNoIllegalPattern(name) {
			fmt.Printf("Username \"%s\" contains illegal pattern\n", name)
		}

		if !onlyContainsLegalRunes(name) {
			fmt.Printf("Username \"%s\" contains illegal characters\n", name)
		}
	}
}
