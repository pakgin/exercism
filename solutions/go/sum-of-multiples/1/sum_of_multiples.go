package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	seen := make(map[int]bool)
	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}
		for i := 1; divisor*i < limit; i++ {
			if flag, ok := seen[divisor*i]; !ok || !flag {
				sum += divisor * i
				seen[divisor*i] = true
			}
		}
	}
	return sum
}
