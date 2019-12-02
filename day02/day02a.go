package main

func CalculatePreviousState(integers []int) int {
	for i := 0; i < len(integers); i += 4 {
		opcode := integers[i]

		if opcode == 99 {
			return integers[0]
		}

		firstPosition := integers[i+1]
		secondPosition := integers[i+2]
		output := integers[i+3]

		switch opcode {
		case 1:
			integers[output] = integers[firstPosition] + integers[secondPosition]
		case 2:
			integers[output] = integers[firstPosition] * integers[secondPosition]
		}
	}

	return integers[0]
}
