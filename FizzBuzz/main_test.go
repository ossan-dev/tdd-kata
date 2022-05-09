package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	testSuite := []struct {
		name     string
		inputArg int
		expected string
	}{
		{
			name:     "ShouldReturnOne_WhenOneIsPassed",
			inputArg: 1,
			expected: "1",
		},
		{
			name:     "ShouldReturnTwo_WhenTwoIsPassed",
			inputArg: 2,
			expected: "2",
		},
		{
			name:     "ShouldReturnFizz_WhenSixIsPassed",
			inputArg: 6,
			expected: "Fizz",
		},
		{
			name:     "ShouldReturnBuzz_WhenTenIsPassed",
			inputArg: 10,
			expected: "Buzz",
		},
		{
			name:     "ShouldReturnFizzBuzz_WhenFifteenIsPassed",
			inputArg: 15,
			expected: "FizzBuzz",
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			got := FizzBuzz(tt.inputArg)

			assert.Equal(t, tt.expected, got)
		})
	}
}
