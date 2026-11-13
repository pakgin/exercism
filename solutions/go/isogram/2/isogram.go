package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	for i, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) && strings.ContainsRune(word[i+1:], r) {
			return false
		}
	}
	return true
}
