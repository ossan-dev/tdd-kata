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

func TestAdd_ShouldReturnFour_WhenOneThreeWithSemicolonArePassed(t *testing.T) {
	inputNumbers := "//;\n1;3"

	got, err := Add(inputNumbers)

	assert.Equal(t, 4, got)
	assert.Nil(t, err)
}

func TestAdd_ShouldReturnSix_WhenOneTwoThreeWithPipeArePassed(t *testing.T) {
	inputNumbers := "//|\n1|2|3"

	got, err := Add(inputNumbers)

	assert.Equal(t, 6, got)
	assert.Nil(t, err)
}

func TestAdd_ShouldReturnSeven_WhenTwoFiveWithSepArePassed(t *testing.T) {
	inputNumbers := "//sep\n2sep5"

	got, err := Add(inputNumbers)

	assert.Equal(t, 7, got)
	assert.Nil(t, err)
}

func TestAdd_ShouldReturnErr_WhenWrongSeparatorIsPassed(t *testing.T) {
	inputNumbers := "//|\n1|2,3"

	_, err := Add(inputNumbers)

	assert.NotNil(t, err)
	assert.Equal(t, "'|' expected but ',' found at position 3.", err.Error())
}
