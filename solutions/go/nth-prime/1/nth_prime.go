package nthprime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Prime number can't be computed")
	}

	i := 0
	for num := 2; ; num++ {
		if isPrime(num) {
			i++
			if i == n {
				return num, nil
			}
		}
	}
}

func isPrime(num int) bool {
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			return false
		}
	}
	return true
}
