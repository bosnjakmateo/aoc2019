package main

import (
	utils "aoc2019/util"
)

func main() {
	instructions := utils.FileInt("day05/day05.txt", ",")

	println("Solution part one:")
	RunDiagnostics(append([]int{}, instructions...), 1)

	print("\nSolution part two: ")
	RunDiagnostics(append([]int{}, instructions...), 5)
}
