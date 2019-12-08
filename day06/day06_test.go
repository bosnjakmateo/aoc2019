package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountNumberOfOrbits(t *testing.T) {
	numberOfOrbits := CountNumberOfOrbits([]string{"A)B", "COM)A"})
	assert.Equal(t, 3, numberOfOrbits)

	numberOfOrbits = CountNumberOfOrbits([]string{"G)H", "D)I", "E)J", "J)K", "COM)B", "E)F", "B)C", "C)D", "D)E", "B)G", "K)L"})
	assert.Equal(t, 42, numberOfOrbits)
}

func TestCountOrbitsToSanta(t *testing.T) {
	numberOfOrbits := CountOrbitsToSanta([]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"})
	assert.Equal(t, 4, numberOfOrbits)
}
