package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	runeToRemove := "+().- "
	phoneNumber = strings.Map(func(r rune) rune {
		if strings.ContainsRune(runeToRemove, r) {
			return -1
		}
		return r
	}, phoneNumber)

	// Handle 11-digit numbers starting with 1
	if len(phoneNumber) == 11 && phoneNumber[0] == '1' {
		phoneNumber = phoneNumber[1:]
	} else if len(phoneNumber) != 10 {
		return "", errors.New("Invalid phone number")
	}

	for i, r := range phoneNumber {
		if !unicode.IsDigit(r) {
			return "", errors.New("Invalid phone number")
		}
		if digit := int(r - '0'); (i == 0 || i == 3) && (digit < 2) {
			return "", errors.New("Invalid phone number")
		}
	}

	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return number, err
	}

	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return number, err
	}

	areaCode, err := AreaCode(number)
	if err != nil {
		return number, err
	}

	return fmt.Sprintf("(%s) %s-%s", areaCode, number[3:6], number[6:]), nil
}
