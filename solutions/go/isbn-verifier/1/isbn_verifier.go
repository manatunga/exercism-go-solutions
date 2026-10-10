package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	noDash := strings.ReplaceAll(isbn, "-", "")
	cleaned := strings.ReplaceAll(noDash, " ", "")

    runes := []rune(cleaned)

    if len(runes) != 10 {
        return false
    }

    modifiers := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
    fallbackNum := 10
	sum := 0
    targetIdx := 9
    
	for i, r := range cleaned {
        isDigit := r >= '0' && r <= '9'
        digit := 0

        if i == targetIdx {
            if !isDigit && r == 'X' {
                digit = fallbackNum
            } else if isDigit {
                digit = int(r) - '0'
            } else {
                return false
            }
        } else {
            if !isDigit {
                return false
            } else {
                digit = int(r) - '0'
            }
        }

        newVal := digit * modifiers[i]
        sum += newVal
	}

    return sum%11 == 0
}