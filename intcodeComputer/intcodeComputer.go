package intcomp

func RunDiagnostics(integers []int, input1, input2 int) int {
	increment := 0
	input := input1

	for i := 0; i < len(integers); i += increment {
		instruction := integers[i]

		if instruction == 99 {
			return 99
		}

		opcode, mode1, mode2, _ := ParseOpcode(instruction)

		switch opcode {
		case 1:
			opcode1(integers, i, mode1, mode2)
			increment = 4
		case 2:
			opcode2(integers, i, mode1, mode2)
			increment = 4
		case 3:
			opcode3(integers, i, input)
			input = input2
			increment = 2
		case 4:
			return opcode4(integers, i, mode1)
		case 5:
			newIndex := opcode5(integers, i, mode1, mode2)
			if newIndex != -1 {
				increment = 0
				i = newIndex
			} else {
				increment = 3
			}
		case 6:
			newIndex := opcode6(integers, i, mode1, mode2)
			if newIndex != -1 {
				increment = 0
				i = newIndex
			} else {
				increment = 3
			}
		case 7:
			opcode7(integers, i, mode1, mode2)
			increment = 4
		case 8:
			opcode8(integers, i, mode1, mode2)
			increment = 4
		}
	}

	return -1
}

func opcode1(integers []int, i, mode1, mode2 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValue(integers, i+3, 1)

	integers[c] = a + b
}

func opcode2(integers []int, i, mode1, mode2 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValue(integers, i+3, 1)

	integers[c] = a * b
}

func opcode3(integers []int, i, input int) {
	integers[integers[i+1]] = input
}

func opcode4(integers []int, i, mode int) int {
	if mode == 0 {
		return integers[integers[i+1]]
	} else {
		return integers[i+1]
	}
}

func opcode5(integers []int, i, mode1, mode2 int) int {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)

	if a != 0 {
		return b
	}

	return -1
}

func opcode6(integers []int, i, mode1, mode2 int) int {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)

	if a == 0 {
		return b
	}

	return -1
}

func opcode7(integers []int, i, mode1, mode2 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValue(integers, i+3, 1)

	if a < b {
		integers[c] = 1
	} else {
		integers[c] = 0
	}
}

func opcode8(integers []int, i, mode1, mode2 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValue(integers, i+3, 1)

	if a == b {
		integers[c] = 1
	} else {
		integers[c] = 0
	}
}

func getValue(integers []int, i, mode int) int {
	if mode == 0 {
		return integers[integers[i]]
	} else {
		return integers[i]
	}
}
