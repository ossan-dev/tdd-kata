package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testSuite := []struct {
		name     string
		pass     string
		expected bool
		errorMsg string
	}{
		{
			"ShouldGetErr_WhenPassLenIsLessThanEight",
			"1A4.567",
			false,
			TOO_SHORT,
		},
		{
			"ShouldGetErr_WhenNotContainsAtLeastTwoNumbers",
			"ab\\2Cdaaa",
			false,
			TOO_FEW_DIGITS,
		},
		{
			"ShouldGetErr_WhenIsLessThanEightAndNotContainsAtLeastTwoNumbers",
			"ab\\2Cd",
			false,
			fmt.Sprintf("%v\n%v", TOO_SHORT, TOO_FEW_DIGITS),
		},
		{
			"ShouldGetErr_WhenNoCapitalLetterIsPresent",
			"aa3?!ab2cd",
			false,
			NO_CAPITAL_LETTER,
		},
		{
			"ShouldGetErr_WhenNoSpecialCharIsPresent",
			"aa3ab2cdA",
			false,
			NO_SPECIAL_CHARS,
		},
		{
			"ShouldGetTrue_WhenPassAreValid",
			"a&&a3ab2cdA",
			true,
			"",
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Validate(tt.pass)

			assert.Equal(t, tt.expected, got)
			if tt.errorMsg != "" {
				assert.Equal(t, tt.errorMsg, err.Error())
			}
		})
	}
}
