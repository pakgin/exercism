package wordy

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var OPERATIONS = []string{"plus", "minus", "multiplied", "divided"}

func Answer(question string) (int, bool) {
	isValid := regexp.MustCompile(`^What is.*\?$`).MatchString(question)
	if !isValid {
		return 0, false
	}

	// Remove "What is" and "?" from question
	if len(question) < 9 {
		return 0, false
	}
	tokens := strings.Split(question[8:len(question)-1], " ")

	numbers := []int{}
	operations := []string{}
	lastToken := -1
	for _, t := range tokens {
		intVal, ok := strconv.Atoi(t)
		if ok == nil {
			if lastToken == 0 {
				return 0, false
			}
			lastToken = 0
			numbers = append(numbers, intVal)
		} else {
			if t == "by" {
				continue
			}
			if lastToken == 1 || !slices.Contains(OPERATIONS, t) {
				return 0, false
			}
			lastToken = 1
			operations = append(operations, t)
		}
	}

	if len(numbers) == 0 || len(numbers) <= len(operations) {
		return 0, false
	}

	numbersIndex := 0
	opreationsIndex := 0
	for numbersIndex < len(numbers) && opreationsIndex < len(operations) {
		switch operations[opreationsIndex] {
		case "plus":
			numbers[numbersIndex+1] = numbers[numbersIndex] + numbers[numbersIndex+1]
		case "minus":
			numbers[numbersIndex+1] = numbers[numbersIndex] - numbers[numbersIndex+1]
		case "multiplied":
			numbers[numbersIndex+1] = numbers[numbersIndex] * numbers[numbersIndex+1]
		case "divided":
			numbers[numbersIndex+1] = numbers[numbersIndex] / numbers[numbersIndex+1]
		}
		numbersIndex++
		opreationsIndex++
	}

	return numbers[len(numbers)-1], true
}
