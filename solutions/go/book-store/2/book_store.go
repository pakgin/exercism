package bookstore

// groupPrice[n] is the cost of a group of n different books.
var groupPrice = [...]int{0, 800, 2 * 760, 3 * 720, 4 * 640, 5 * 600}

func Cost(books []int) int {
	counts := make(map[int]int)
	for _, book := range books {
		counts[book]++
	}

	// Each pass takes one copy of every title left, so the group size is
	// the number of distinct titles remaining.
	var groups [6]int
	for len(counts) > 0 {
		groups[len(counts)]++
		for book := range counts {
			counts[book]--
			if counts[book] == 0 {
				delete(counts, book)
			}
		}
	}

	// A group of 5 plus a group of 3 (3000+2160) costs more than two groups
	// of 4 (2560+2560), so turn every 5+3 pair into 4+4.
	swaps := min(groups[5], groups[3])
	groups[5] -= swaps
	groups[3] -= swaps
	groups[4] += 2 * swaps

	total := 0
	for size, n := range groups {
		total += n * groupPrice[size]
	}
	return total
}
