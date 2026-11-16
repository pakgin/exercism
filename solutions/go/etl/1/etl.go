package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	output := map[string]int{}
	for score, letters := range in {
		for _, letter := range letters {
			output[strings.ToLower(letter)] = score
		}
	}
	return output
}
