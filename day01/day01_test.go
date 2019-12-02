package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01a(t *testing.T) {
	got := CalculateModuleFuel([]int{12})
	assert.Equal(t, 2, got)
}

func TestDay01b(t *testing.T) {
	got := CalculateModuleFuel([]int{14})
	assert.Equal(t, 2, got)
}
