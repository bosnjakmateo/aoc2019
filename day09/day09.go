package main

import (
	"aoc2019/intcodeComputer"
)

func CalculateBOOSTKeyCode(values, inputs []int) []int {
	for i := 0; i < 2000; i++ {
		values = append(values, 0)
	}

	outputCodes := intcodeComputer.RunDiagnosticsPrint(values, inputs)

	return outputCodes
}
