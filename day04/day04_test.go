package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay04a(t *testing.T) {
	assert.True(t, PasswordValid("111111"))
	assert.False(t, PasswordValid("223450"))
	assert.False(t, PasswordValid("123789"))
}

func TestDay04b(t *testing.T) {
	assert.True(t, PasswordValidLargeGroup("112233"))
	assert.False(t, PasswordValidLargeGroup("123444"))
	assert.True(t, PasswordValidLargeGroup("111122"))
}
