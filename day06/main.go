package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileString("day06/day06.txt", "\r\n")

	fmt.Printf("Soulution part one: %d\n", CountNumberOfOrbits(values))
	fmt.Printf("Soulution part two: %d\n", CountOrbitsToSanta(values))
}
