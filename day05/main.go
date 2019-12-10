package main

import (
	"aoc2019/intcodeComputer"
	utils "aoc2019/util"
	"fmt"
)

func main() {
	instructions := utils.FileInt("day05/day05.txt", ",")

	result := intcodeComputer.RunDiagnosticsPrint(append([]int{}, instructions...), []int{1})
	fmt.Printf("Solution part one: %v\n", result)

	result = intcodeComputer.RunDiagnosticsPrint(append([]int{}, instructions...), []int{5})
	fmt.Printf("Solution part two: %d\n", result[0])
}
