package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay03a(t *testing.T) {
	actual := CalculateMinManhattanDistance(
		[]string{"R8", "U5", "L5", "D3"},
		[]string{"U7", "R6", "D4", "L4"},
	)
	assert.Equal(t, 6, actual)

	actual = CalculateMinManhattanDistance(
		[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
	)
	assert.Equal(t, 159, actual)
}

func TestDay03b(t *testing.T) {
	actual := CalculateMinSteps(
		[]string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
		[]string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
	)
	assert.Equal(t, 610, actual)

	actual = CalculateMinSteps(
		[]string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
		[]string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
	)
	assert.Equal(t, 410, actual)

	actual = CalculateMinSteps(
		[]string{"R8", "U5", "L5", "D3"},
		[]string{"U7", "R6", "D4", "L4"},
	)
	assert.Equal(t, 30, actual)
}
