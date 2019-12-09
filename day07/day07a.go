package main

import (
	intcomp "aoc2019/intcodeComputer"
	utils "aoc2019/util"
)

func CalculateMaxThrusterSignal(values []int) (maxThrusterSignal int) {
	combinations := utils.GenerateCombinations([]int{4, 3, 2, 1, 0})

	for i := range combinations {
		thrusterSignal := runAmplifiers(append([]int{}, values...), combinations[i])

		if thrusterSignal > maxThrusterSignal {
			maxThrusterSignal = thrusterSignal
		}
	}

	return maxThrusterSignal
}

func runAmplifiers(values, phases []int) (thrusterSignal int) {
	input := 0

	for i := range phases {
		input = intcomp.RunDiagnostics(append([]int{}, values...), phases[i], input)
	}

	return input
}
