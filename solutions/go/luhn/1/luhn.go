package luhn

import "unicode"

func Valid(id string) bool {
	var digits []int

	for _, r := range id {
		if r == ' ' {
			continue 
		}
		if !unicode.IsDigit(r) {
			return false 
		}
		digits = append(digits, int(r-'0'))
	}

	if len(digits) <= 1 {
		return false
	}

	sum := 0
	isSecond := false

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if isSecond {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		isSecond = !isSecond
	}

	return sum%10 == 0
}
