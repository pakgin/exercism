package bookstore

import (
	"container/heap"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func Cost(books []int) int {
	h := &maxHeap{}
	m := make(map[int]int)
	for _, book := range books {
		m[book]++
	}
	for _, v := range m {
		heap.Push(h, v)
	}

	total := 0
	fives, threes := 0, 0
	for h.Len() > 0 {
		if h.Len() >= 5 {
			total += group(h, 5)
			fives++
		} else if h.Len() >= 4 {
			total += group(h, 4)
		} else if h.Len() >= 3 {
			total += group(h, 3)
			threes++
		} else if h.Len() >= 2 {
			total += group(h, 2)
		} else {
			total += group(h, 1)
		}
	}

	// A group of 5 plus a group of 3 (3000+2160) costs more than two groups
	// of 4 (2560+2560), so turn every 5+3 pair into 4+4.
	total -= min(fives, threes) * (800*.75*5 + 800*.9*3 - 2*800*.8*4)

	return total
}

func group(h *maxHeap, size int) int {
	buf := make([]int, 0, 5)
	switch size {
	case 5:
		{
			for range 5 {
				n := h.Pop().(int)
				if n-1 > 0 {
					buf = append(buf, n-1)
				}
			}

			for _, n := range buf {
				heap.Push(h, n)
			}
			return 800 * .75 * 5
		}
	case 4:
		{
			for range 4 {
				n := h.Pop().(int)
				if n-1 > 0 {
					buf = append(buf, n-1)
				}
			}

			for _, n := range buf {
				heap.Push(h, n)
			}
			return 800 * 0.8 * 4
		}
	case 3:
		{
			for range 3 {
				n := h.Pop().(int)
				if n-1 > 0 {
					buf = append(buf, n-1)
				}
			}

			for _, n := range buf {
				heap.Push(h, n)
			}
			return 800 * 0.9 * 3
		}
	case 2:
		{
			for range 2 {
				n := h.Pop().(int)
				if n-1 > 0 {
					buf = append(buf, n-1)
				}
			}

			for _, n := range buf {
				heap.Push(h, n)
			}
			return 800 * 0.95 * 2
		}
	default:
		{
			n := h.Pop().(int)
			if n-1 > 0 {
				buf = append(buf, n-1)
			}

			for _, n := range buf {
				heap.Push(h, n)
			}
			return 800
		}
	}
}
