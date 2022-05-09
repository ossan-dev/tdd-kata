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

func TestAdd_ShouldReturnOne_WhenOneIsPassed(t *testing.T) {
	inputNumbers := "1"

	got := Add(inputNumbers)

	assert.Equal(t, 1, got)
}

func TestAdd_ShouldReturnThree_WhenOneTwoIsPassed(t *testing.T) {
	inputNumbers := "1,2"

	got := Add(inputNumbers)

	assert.Equal(t, 3, got)
}
