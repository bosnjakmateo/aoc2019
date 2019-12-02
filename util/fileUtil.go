package utils

import (
	"bufio"
	"os"
	"strconv"
)

func File(filePath string) (values []int) {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		values = append(values, i)
	}

	return values
}
