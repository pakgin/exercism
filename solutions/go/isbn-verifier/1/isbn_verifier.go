package isbnverifier

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}

	num := 0
	for i, ch := range isbn {
		if !(unicode.IsDigit(ch) || (ch == 'X' && i == 9)) {
			return false
		}

		switch ch {
		case 'X':
			num += 10 * (10 - i)
		default:
			num += int(ch-'0') * (10 - i)
		}
	}
	return num%11 == 0
}
