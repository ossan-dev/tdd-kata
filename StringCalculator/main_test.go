package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_ShouldReturnZero_WhenStringEmptyIsPassed(t *testing.T) {
	inputNumbers := ""

	got := Add(inputNumbers)

	assert.Equal(t, 0, got)
}
