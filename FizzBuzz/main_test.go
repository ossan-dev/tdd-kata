package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz_ShouldReturnOne_WhenOneIsPassed(t *testing.T) {
	// arrange
	intArg := 1

	// act
	got := FizzBuzz(intArg)

	// assert
	assert.Equal(t, "1", got)
}

func TestFizzBuzz_ShouldReturnTwo_WhenTwoIsPassed(t *testing.T) {
	// arrange
	intArg := 2

	// act
	got := FizzBuzz(intArg)

	// assert
	assert.Equal(t, "2", got)
}

func TestFizzBuzz_ShouldReturnFizz_WhenSixIsPassed(t *testing.T) {
	// arrange
	intArg := 6

	// act
	got := FizzBuzz(intArg)

	// assert
	assert.Equal(t, "Fizz", got)
}

func TestFizzBuzz_ShouldReturnBuzz_WhenTenIsPassed(t *testing.T) {
	// arrange
	intArg := 10

	// act
	got := FizzBuzz(intArg)

	// assert
	assert.Equal(t, "Buzz", got)
}
