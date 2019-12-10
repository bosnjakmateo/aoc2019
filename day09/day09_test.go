package main

import (
	utils "aoc2019/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateBOOSTKeyCode(t *testing.T) {
	var codes []int

	codes = CalculateBOOSTKeyCode([]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, []int{})
	assert.Equal(t, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, codes)

	codes = CalculateBOOSTKeyCode([]int{104, 1125899906842624, 99}, []int{})
	assert.Equal(t, []int{1125899906842624}, codes)

	codes = CalculateBOOSTKeyCode([]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, []int{})
	assert.Equal(t, 16, utils.Len(codes[0]))
}
