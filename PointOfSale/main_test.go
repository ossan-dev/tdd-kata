package pointofsale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: requirement #3
func TestScan_12345(t *testing.T) {
	got, _ := Scan("12345")
	assert.Equal(t, "$7.25", got)
}

func TestScan_23456(t *testing.T) {
	got, _ := Scan("23456")
	assert.Equal(t, "$12.50", got)
}

func TestScan_99999(t *testing.T) {
	got, err := Scan("99999")
	assert.Equal(t, "$0.00", got)
	assert.EqualError(t, err, "Error: barcode not found")
}
