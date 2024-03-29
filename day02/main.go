package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileInt("day02/day02.txt", "\r\n")
	initialValues := append([]int{}, values...)

	values[1] = 12
	values[2] = 2

	fmt.Printf("Solution part one: %d\n", CalculatePreviousState(values))
	fmt.Printf("Solution part two: %d\n", FindNounAndVerb(initialValues))
}
