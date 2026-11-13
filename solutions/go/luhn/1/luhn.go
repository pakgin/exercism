package luhn

import "unicode"

func Valid(id string) bool {
	sum := 0
	var new_id []int
	for _, v := range id {
		if !(unicode.IsNumber(v) || unicode.IsSpace(v)) {
			return false
		}
		if unicode.IsNumber(v) {
			new_id = append(new_id, int(v)-'0')
		}
	}
	if len(new_id) == 1 && new_id[0] == 0 {
		return false
	}
	for i := len(new_id) - 2; i >= 0; i -= 2 {
		new_id[i] *= 2
		if new_id[i] > 9 {
			new_id[i] -= 9
		}
	}

	for _, v := range new_id {
		sum += v
	}

	return sum%10 == 0
}
