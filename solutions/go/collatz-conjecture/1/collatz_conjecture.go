package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	loop := 0

	if n <= 0 {
		return 0, errors.New("only positive numbers are allowed")
	}

	for n > 1 {
		switch {
		case n%2 == 0:
			n /= 2
		default:
			n = n*3 + 1
		}
		loop++
	}

	return loop, nil
}
