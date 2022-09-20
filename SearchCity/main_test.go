package searchcity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCity_Correct(t *testing.T) {
	got := SearchCity("st")
	assert.Contains(t, got, "Budapest")
}

func TestSearchCity_LessThanTwoCharsPassedAsInput(t *testing.T) {
	got := SearchCity("s")
	assert.Empty(t, got, "s")
}
