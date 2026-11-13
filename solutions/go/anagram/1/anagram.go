package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	loweredSubject := strings.ToLower(subject)
	sortedSubject := sortString(loweredSubject)

	ans := []string{}
	for i, candidate := range candidates {
		loweredCandidate := strings.ToLower(candidate)
		sortedCandidate := sortString(loweredCandidate)
		if sortedCandidate == sortedSubject && loweredSubject != loweredCandidate {
			ans = append(ans, candidates[i])
		}
	}

	return ans
}

func sortString(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
