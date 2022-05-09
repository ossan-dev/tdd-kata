package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_ShouldReturnZero_WhenStringEmptyIsPassed(t *testing.T) {
	inputNumbers := ""

	got, _ := Add(inputNumbers)

	assert.Equal(t, 0, got)
}

func TestAdd_ShouldReturnOne_WhenOneIsPassed(t *testing.T) {
	inputNumbers := "1"

	got, _ := Add(inputNumbers)

	assert.Equal(t, 1, got)
}

func TestAdd_ShouldReturnFour_WhenTwoTwoArePassed(t *testing.T) {
	inputNumbers := "2,2"

	got, _ := Add(inputNumbers)

	assert.Equal(t, 4, got)
}

func TestAdd_ShouldReturnTen_WhenOneTwoThreeFourArePassed(t *testing.T) {
	inputNumbers := "1,2,3,4"

	got, _ := Add(inputNumbers)

	assert.Equal(t, 10, got)
}

func TestAdd_ShouldReturnSix_WhenOneTwoNewLineThreeArePassed(t *testing.T) {
	inputNumbers := "1,2\n3"

	got, _ := Add(inputNumbers)

	assert.Equal(t, 6, got)
}

func TestAdd_ShouldReturnErr_WhenTwoCommaNewLineThreeArePassed(t *testing.T) {
	inputNumbers := "2,\n3"

	_, err := Add(inputNumbers)

	assert.NotNil(t, err)
}

func TestAdd_ShouldReturnErr_WhenTwoColonNewLineThreeArePassed(t *testing.T) {
	inputNumbers := "2:\n3"

	_, err := Add(inputNumbers)

	assert.NotNil(t, err)
}
