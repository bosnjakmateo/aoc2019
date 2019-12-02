package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay02a(t *testing.T) {
	actual := CalculatePreviousState([]int{1, 1, 1, 4, 99, 5, 6, 0, 99})
	assert.Equal(t, 30, actual)
}
