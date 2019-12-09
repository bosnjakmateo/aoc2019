package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func GenerateCombinations(numbers []int) (combinations [][]int) {
	for i := 0; i <= int(math.Pow10(len(numbers))); i++ {
		combination := fillWithZeros(strconv.Itoa(i))

		if containsNumbers(combination, numbers) {
			combinations = append(combinations, ToIntArray(strings.Split(combination, "")))
		}
	}

	return combinations
}

func containsNumbers(value string, numbers []int) bool {
	for i := range numbers {
		if !strings.Contains(value, strconv.Itoa(numbers[i])) {
			return false
		}
	}

	return true
}

func fillWithZeros(value string) string {
	return fmt.Sprintf("%05s", value)
}
