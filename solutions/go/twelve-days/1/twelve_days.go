package twelve

import (
	"fmt"
	"strings"
)

var numbers = []string{"", "first", "second", "third", "fourth", "fifth",
	"sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var phrases = []string{"", "a Partridge in a Pear Tree", "two Turtle Doves",
	"three French Hens", "four Calling Birds", "five Gold Rings",
	"six Geese-a-Laying", "seven Swans-a-Swimming", "eight Maids-a-Milking",
	"nine Ladies Dancing", "ten Lords-a-Leaping", "eleven Pipers Piping",
	"twelve Drummers Drumming"}

func processText(i int) (s string) {
	switch {
	case i == 1:
		s = "a Partridge in a Pear Tree"
	case 1 < i && i <= 12:
		ss := phrases[2 : i+1]
		for i := len(ss) - 1; i >= 0; i-- {
			s += ss[i]
			if i >= 1 {
				s += ", "
			}
		}
		s += ", and " + phrases[1]
	default:
		panic("invalid argument")
	}
	return
}

func Verse(i int) string {
	return fmt.Sprintf(
		"On the %s day of Christmas my true love gave to me: %s.",
		numbers[i], processText(i),
	)
}

func Song() string {
	ss := make([]string, 0)
	for i := 1; i <= 12; i++ {
		ss = append(ss, Verse(i))
	}
	return strings.Join(ss, "\n")
}
