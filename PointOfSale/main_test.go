package pointofsale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	got := Scan("12345")
	assert.Equal(t, "$7.25", got)
}
