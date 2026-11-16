package binarysearch

func SearchInts(list []int, key int) int {
	left, right := 0, len(list)
	for left < right {
		mid := (left + right) / 2
		if list[mid] == key {
			return mid
		} else if list[mid] < key {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}
