package pangram

import "strings"

func IsPangram(input string) bool {
	seen := make(map[rune]bool)
	for i := range 26 {
		seen[rune('a'+i)] = false
	}

	input = strings.ToLower(input)
	for _, char := range input {
		if _, ok := seen[char]; ok {
			seen[char] = true
		}
	}

	for _, v := range seen {
		if !v {
			return false
		}
	}
	return true
}
