package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01a(t *testing.T) {
	actual := CalculateModuleFuel([]int{12})
	assert.Equal(t, 2, actual)
}

func TestDay01b(t *testing.T) {
	actual := CalculateModuleFuel([]int{14})
	assert.Equal(t, 2, actual)
}
