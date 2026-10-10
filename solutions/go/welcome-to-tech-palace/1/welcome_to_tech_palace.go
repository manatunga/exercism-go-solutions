package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	lines := strings.Split(oldMsg, "\n")
    if len(lines) >= 2 {
        targetLine := lines[1]
        cleanedLine := strings.Replace(targetLine, "*", "", 2)
        return strings.TrimSpace(cleanedLine)
    } else {
        return ""
    }
}
