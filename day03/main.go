package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	wire1 := utils.FileString("day03/day03w1.txt")
	wire2 := utils.FileString("day03/day03w2.txt")

	fmt.Printf("Solution part one: %d\n", CalculateMinManhattanDistance(wire1, wire2))
	fmt.Printf("Solution part two: %d\n", CalculateMinSteps(wire1, wire2))
}
