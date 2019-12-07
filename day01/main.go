package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileInt("day01/day01.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateModuleFuel(values))
	fmt.Printf("Solution part two: %d\n", CalculateAllNonNegativeFuel(values))
}
