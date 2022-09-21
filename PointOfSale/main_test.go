package pointofsale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan_12345(t *testing.T) {
	got := Scan("12345")
	assert.Equal(t, "$7.25", got)
}

func TestScan_23456(t *testing.T) {
	got := Scan("23456")
	assert.Equal(t, "$12.50", got)
}
