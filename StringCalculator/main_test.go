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

func TestAdd_ShouldReturnFour_WhenTwoTwoIsPassed(t *testing.T) {
	inputNumbers := "2,2"

	got := Add(inputNumbers)

	assert.Equal(t, 4, got)
}
