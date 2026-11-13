package raindrops

import "fmt"

func Convert(number int) (s string) {
	if number%3 == 0 {
		s += "Pling"
	}
	if number%5 == 0 {
		s += "Plang"
	}
	if number%7 == 0 {
		s += "Plong"
	}
	if s != "" {
		return
	}
	return fmt.Sprint(number)
}
