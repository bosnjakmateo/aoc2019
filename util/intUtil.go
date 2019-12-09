package utils

import "strconv"

func ToIntArray(values []string) (numbers []int) {
	for i := range values {
		n, _ := strconv.Atoi(values[i])
		numbers = append(numbers, n)
	}

	return numbers
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MinMax(a, b int) (min, max int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return min, max
}

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
