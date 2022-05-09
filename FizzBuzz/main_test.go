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
	assert.Equal(t, got, "1")
}
