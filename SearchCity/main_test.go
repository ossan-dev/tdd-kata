package searchcity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCity(t *testing.T) {
	test_suite := []struct {
		name                   string
		input                  string
		cities_to_be_in_result []string
	}{
		{
			name:                   "MustBeEmpty_WhenLessThanTwoCharsArePassed",
			input:                  "a",
			cities_to_be_in_result: []string{},
		},
		{
			name:                   "MustReturnCities_WhenInputIsLongerThanTwoChars",
			input:                  "Va",
			cities_to_be_in_result: []string{"Valencia", "Vancouver"},
		},
		{
			name:                   "MustReturnCities_WhenInputCharsAreNotInTheCorrectCase",
			input:                  "VA",
			cities_to_be_in_result: []string{"Valencia", "Vancouver"},
		},
		{
			name:                   "MustReturnCities_WhenInputCharsAreInTheMiddle",
			input:                  "aPe",
			cities_to_be_in_result: []string{"Budapest"},
		},
	}
	for _, tt := range test_suite {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchCity(tt.input)
			for _, v := range tt.cities_to_be_in_result {
				if !assert.Contains(t, got, v) {
					t.Errorf("expected %q to be present in result which contains %q", v, got)
				}
			}
		})
	}
}
