package main

import (
	"aoc2019/intcodeComputer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCode(t *testing.T) {
	opcode, mode1, mode2, mode3 := intcodeComputer.ParseOpcode(1002)

	assert.Equal(t, 2, opcode)
	assert.Equal(t, 0, mode1)
	assert.Equal(t, 1, mode2)
	assert.Equal(t, 0, mode3)

	opcode, mode1, mode2, mode3 = intcodeComputer.ParseOpcode(10002)

	assert.Equal(t, 2, opcode)
	assert.Equal(t, 0, mode1)
	assert.Equal(t, 0, mode2)
	assert.Equal(t, 1, mode3)
}
