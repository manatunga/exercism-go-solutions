// Package leap determines whether a given year
// is a leap year or not.
package leap

// IsLeapYear calculates if the given year is a leap year.
func IsLeapYear(year int) bool {
	if (year % 4 == 0 && year % 100 != 0) {
        return true
    } else if year % 400 == 0 {
        return true	
    } else {
        return false
    }
}
