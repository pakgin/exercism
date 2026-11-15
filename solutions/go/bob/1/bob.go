// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	isQuestion := strings.HasSuffix(trimmed, "?")
	isAllCap := isAllCap(trimmed)
	isSilence := trimmed == ""

	switch {
	case isAllCap && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isQuestion:
		return "Sure."
	case isAllCap:
		return "Whoa, chill out!"
	case isSilence:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}

func isAllCap(s string) bool {
	if len(s) == 0 {
		return false
	}

	flag := false
	for _, r := range s {
		if !flag && unicode.IsLetter(r) {
			flag = true
		}
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}
	}
	return flag
}
