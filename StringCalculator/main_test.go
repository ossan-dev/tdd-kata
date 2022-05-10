package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testSuite := []struct {
		name          string
		numbers       string
		resExpected   int
		isErrExpected bool
		errMessage    string
	}{
		{
			"ShouldReturnZero_WhenStringEmptyIsPassed",
			"",
			0,
			false,
			"",
		},
		{
			"ShouldReturnOne_WhenOneIsPassed",
			"1",
			1,
			false,
			"",
		},
		{
			"ShouldReturnFour_WhenTwoTwoArePassed",
			"2,2",
			4,
			false,
			"",
		},
		{
			"ShouldReturnTen_WhenOneTwoThreeFourArePassed",
			"1,2,3,4",
			10,
			false,
			"",
		},
		{
			"ShouldReturnSix_WhenOneTwoNewLineThreeArePassed",
			"1,2\n3",
			6,
			false,
			"",
		},
		{
			"WhenTwoCommaNewLineThreeArePassed",
			"2,\n3",
			0,
			true,
			"",
		},
		{
			"ShouldReturnErr_WhenTwoColonNewLineThreeArePassed",
			"2:\n3",
			0,
			true,
			"",
		},
		{
			"ShouldReturnFour_WhenOneThreeWithSemicolonArePassed",
			"//;\n1;3",
			4,
			false,
			"",
		},
		{
			"ShouldReturnSix_WhenOneTwoThreeWithPipeArePassed",
			"//|\n1|2|3",
			6,
			false,
			"",
		},
		{
			"ShouldReturnSeven_WhenTwoFiveWithSepArePassed",
			"//sep\n2sep5",
			7,
			false,
			"",
		},
		{
			"ShouldReturnErr_WhenWrongSeparatorIsPassed",
			"//|\n1|2,3",
			0,
			true,
			"'|' expected but ',' found at position 3",
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.numbers)

			assert.Equal(t, tt.resExpected, got)
			if !tt.isErrExpected {
				assert.Nil(t, err)
				return
			}

			assert.NotNil(t, err)
			if tt.errMessage != "" {
				assert.Equal(t, tt.errMessage, err.Error())
			}
		})
	}
}
