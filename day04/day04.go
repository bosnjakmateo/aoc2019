package main

import (
	"strings"
)

func PasswordValid(password string) bool {
	return digitsDecrease(password) && doubleDigits(password)
}

func PasswordValidLargeGroup(password string) bool {
	return digitsDecrease(password) && doubleDigitsNoLargeGroup(password)
}

func doubleDigits(password string) bool {
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			return true
		}
	}

	return false
}

func digitsDecrease(password string) bool {
	for i := 0; i < len(password)-1; i++ {
		if password[i] > password[i+1] {
			return false
		}
	}

	return true
}

func doubleDigitsNoLargeGroup(password string) bool {
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			digitCount := strings.Count(password, string(password[i]))
			if digitCount == 2 {
				return true
			}
		}
	}

	return false
}
