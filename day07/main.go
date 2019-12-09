package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileInt("day07/day07.txt", ",")

	fmt.Printf("Solution part one: %d\n", CalculateMaxThrusterSignal(values))
}
