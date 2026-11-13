package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	counter := make(map[rune]int)
	for _, r := range strings.ToLower(word) {
		if unicode.IsLetter(r) {
			counter[r]++
			if counter[r] > 1 {
				return false
			}
		}
	}
	return true
}
