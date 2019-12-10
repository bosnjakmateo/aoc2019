package intcodeComputer

var relativeBase = 0

func RunDiagnosticsPrint(integers, inputs []int) (outputCodes []int) {
	increment := 0
	inputIndex := 0

	for i := 0; i < len(integers); i += increment {
		instruction := integers[i]

		if instruction == 99 {
			return
		}

		opcode, mode1, mode2, mode3 := ParseOpcode(instruction)

		switch opcode {
		case 1:
			opcode1(integers, i, mode1, mode2, mode3)
			increment = 4
		case 2:
			opcode2(integers, i, mode1, mode2, mode3)
			increment = 4
		case 3:
			opcode3(integers, i, inputs[inputIndex], mode1)
			inputIndex++
			increment = 2
		case 4:
			output := opcode4(integers, i, mode1)
			outputCodes = append(outputCodes, output)
			increment = 2
		case 5:
			newIndex := opcode5(integers, i, mode1, mode2, mode3)
			if newIndex != -1 {
				increment = 0
				i = newIndex
			} else {
				increment = 3
			}
		case 6:
			newIndex := opcode6(integers, i, mode1, mode2, mode3)
			if newIndex != -1 {
				increment = 0
				i = newIndex
			} else {
				increment = 3
			}
		case 7:
			opcode7(integers, i, mode1, mode2, mode3)
			increment = 4
		case 8:
			opcode8(integers, i, mode1, mode2, mode3)
			increment = 4
		case 9:
			relativeBase = opcode9(integers, i, mode1)
			increment = 2
		}
	}
	return
}

func opcode1(integers []int, i, mode1, mode2, mode3 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValueForWrite(integers, i+3, mode3)

	integers[c] = a + b
}

func opcode2(integers []int, i, mode1, mode2, mode3 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValueForWrite(integers, i+3, mode3)

	integers[c] = a * b
}

func opcode3(integers []int, i, input, mode int) {
	if mode == 2 {
		integers[integers[i+1]+relativeBase] = input
	} else {
		integers[integers[i+1]] = input
	}
}

func opcode4(integers []int, i, mode int) int {
	switch mode {
	case 0:
		return integers[integers[i+1]]
	case 1:
		return integers[i+1]
	case 2:
		return integers[integers[i+1]+relativeBase]
	}

	panic("mode not supported")
	return -1
}

func opcode5(integers []int, i, mode1, mode2, mode3 int) int {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)

	if a != 0 {
		return b
	}

	return -1
}

func opcode6(integers []int, i, mode1, mode2, mode3 int) int {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)

	if a == 0 {
		return b
	}

	return -1
}

func opcode7(integers []int, i, mode1, mode2, mode3 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValueForWrite(integers, i+3, mode3)

	if a < b {
		integers[c] = 1
	} else {
		integers[c] = 0
	}
}

func opcode8(integers []int, i, mode1, mode2, mode3 int) {
	a := getValue(integers, i+1, mode1)
	b := getValue(integers, i+2, mode2)
	c := getValueForWrite(integers, i+3, mode3)

	if a == b {
		integers[c] = 1
	} else {
		integers[c] = 0
	}
}

func opcode9(integers []int, i, mode1 int) int {
	value := getValue(integers, i+1, mode1)

	return relativeBase + value
}

func getValue(integers []int, i, mode int) int {
	switch mode {
	case 0:
		return integers[integers[i]]
	case 1:
		return integers[i]
	case 2:
		return integers[integers[i]+relativeBase]
	}

	panic("unsupported mode")
	return -1
}

func getValueForWrite(integers []int, i, mode int) int {
	switch mode {
	case 0:
		return integers[i]
	case 2:
		return integers[i] + relativeBase
	}

	panic("unsupported mode")
	return -1
}
