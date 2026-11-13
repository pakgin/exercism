package luhn

func Valid(id string) bool {
	sum, digit := 0, 0

	for i := len(id) - 1; i >= 0; i-- {
		switch c := id[i]; {
		case c == ' ':
			continue
		case '0' <= c && c <= '9':
			digit++
			num := int(c - '0')
			if digit%2 == 0 {
				num *= 2
			}
			if num > 9 {
				num -= 9
			}
			sum += num
		default:
			return false
		}
	}

	return sum%10 == 0 && digit > 1
}
