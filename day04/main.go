package main

import (
	"fmt"
	"strconv"
)

func main() {
	from := 264793
	to := 803935

	validCount := 0
	validLargerGroupCount := 0

	for password := from; password <= to; password++ {
		stringPassword := strconv.Itoa(password)

		if PasswordValid(stringPassword) {
			validCount++
		}

		if PasswordValidLargeGroup(stringPassword) {
			validLargerGroupCount++
		}
	}

	fmt.Printf("Solution part one: %d\n", validCount)
	fmt.Printf("Solution part two: %d\n", validLargerGroupCount)
}
