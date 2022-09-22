package pointofsale

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	test_suite := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "MustReturn$7.25_WhenInputIs12345",
			input:  "12345",
			output: "$7.25",
			err:    nil,
		},
		{
			name:   "MustReturn$12.50_WhenInputIs23456",
			input:  "23456",
			output: "$12.50",
			err:    nil,
		},
		{
			name:   "MustReturnBarcodeNotFound_WhenInputNotExists",
			input:  "99999",
			output: "$0.00",
			err:    fmt.Errorf("error: %v", ErrBarcodeNotFound),
		},
		{
			name:   "MustReturnEmptyBarcode_WhenInputIsBlank",
			input:  "",
			output: "$0.00",
			err:    fmt.Errorf("error: %v", ErrBarcodeEmpty),
		},
	}
	for _, tt := range test_suite {
		t.Run(tt.name, func(t *testing.T) {
			shopping_cart := NewShoppingCart(0.00)
			got, err := shopping_cart.Scan(tt.input)
			if tt.err != nil {
				assert.Equal(t, "$0.00", got)
				assert.EqualError(t, err, tt.err.Error())
				return
			}
			assert.Equal(t, tt.output, got)
			assert.Nil(t, err)
		})
	}
}

func TestGetTotal(t *testing.T) {
	shopping_cart := NewShoppingCart(0.0)
	shopping_cart.Scan("12345")
	got := shopping_cart.GetTotal()
	assert.Equal(t, "$7.25", got)
}

func TestNewShoppingCart_NegativeStartValue(t *testing.T) {
	shopping_cart, err := NewShoppingCart(-12.50)
	assert.Nil(t, shopping_cart)
	assert.EqualError(t, err, "error: total of shopping cart cannot be less than zero")
}
