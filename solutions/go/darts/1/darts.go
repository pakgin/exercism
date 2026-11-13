package darts

import "math"

func Score(x, y float64) int {
	dist := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case dist <= 1:
		return 10
	case dist <= 5:
		return 5
	case dist <= 10:
		return 1
	default:
		return 0
	}
}
