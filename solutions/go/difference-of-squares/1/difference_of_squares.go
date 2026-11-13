package diffsquares

import "math"

func SquareOfSum(n int) int {
	counter := 0
	for i := 1; i <= n; i++ {
		counter += i
	}
	return int(math.Pow(float64(counter), 2.0))
}

func SumOfSquares(n int) int {
	counter := 0.0
	for i := 1; i <= n; i++ {
		counter += math.Pow(float64(i), 2.0)
	}
	return int(counter)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
