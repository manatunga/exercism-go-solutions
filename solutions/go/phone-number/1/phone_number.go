package phonenumber

import (
    "errors"
    "strings"
    "unicode"
)

func Number(phoneNumber string) (string, error) {
    if strings.HasPrefix(phoneNumber, "+1") {
        phoneNumber = phoneNumber[len("+1"):]
    } else if strings.HasPrefix(phoneNumber, "1") {
        phoneNumber = phoneNumber[len("1"):]
    }
    if strings.Contains(phoneNumber, "(") {
        phoneNumber = strings.ReplaceAll(phoneNumber, "(", "")
        phoneNumber = strings.ReplaceAll(phoneNumber, ")", "")
    }
    if strings.Contains(phoneNumber, "-") {
        phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
    }
    if strings.Contains(phoneNumber, ".") {
        phoneNumber = strings.ReplaceAll(phoneNumber, ".", "")
    }
    if strings.Contains(phoneNumber, " ") {
        phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
    }
    if len(phoneNumber) != 10 {
        return "", errors.New("Invalid phone number entered.")
    }
    for _, r := range phoneNumber {
        if !unicode.IsDigit(r) {
            return "", errors.New("Invalid phone number entered.")
        }
    }
    if phoneNumber[0] < '2' {
        return "", errors.New("Invalid phone number entered.")
    }
    if phoneNumber[3] < '2' {
        return "", errors.New("Invalid phone number entered.")
    }
    return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return cleanNumber[:3], nil
}

func Format(phoneNumber string) (string, error) {
	cleanNumber, err := Number(phoneNumber)
    if err != nil {
        return "", err
    }
    return "(" + cleanNumber[:3] + ") " + cleanNumber[3:6] + "-" + cleanNumber[6:], nil
}
