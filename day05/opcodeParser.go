package main

import "strconv"

func ParseOpcode(value int) (opcode, mode1, mode2, mode3 int) {
	stringValue := strconv.Itoa(value)
	opcodeIndex := len(stringValue) - 1

	opcode, _ = strconv.Atoi(stringValue[opcodeIndex:])

	if opcodeIndex == 0 {
		return opcode, 0, 0, 0
	}

	modes := stringValue[:opcodeIndex-1]

	modes, mode1 = getMode(modes)
	modes, mode2 = getMode(modes)
	modes, mode3 = getMode(modes)

	return opcode, mode1, mode2, mode3
}

func getMode(modes string) (string, int) {
	if len(modes) == 0 {
		return modes, 0
	}

	mode, _ := strconv.Atoi(modes[len(modes)-1:])
	modes = modes[:len(modes)-1]

	return modes, mode
}
