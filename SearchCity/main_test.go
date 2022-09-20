package searchcity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCity(t *testing.T) {
	got := SearchCity("st")
	assert.Contains(t, got, "st")
}
