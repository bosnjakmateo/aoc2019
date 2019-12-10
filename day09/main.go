package main

import (
	utils "aoc2019/util"
	"fmt"
)

func main() {
	values := utils.FileInt("day09/day09.txt", ",")

	fmt.Println("Solution part one:")
	code := CalculateBOOSTKeyCode(values, []int{1})
	for i := range code {
		println(code[i])
	}

	fmt.Println("Solution part two:")
	code = CalculateBOOSTKeyCode(values, []int{2})
	for i := range code {
		println(code[i])
	}
}
