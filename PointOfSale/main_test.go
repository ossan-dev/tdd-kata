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
			shopping_cart, _ := NewShoppingCart(0.00)
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

func TestGetTotal_ReturnStartValueWithoutAnyScans(t *testing.T) {
	start_amount := 8.50
	shopping_cart, _ := NewShoppingCart(start_amount)
	got := shopping_cart.GetTotal()
	assert.Equal(t, fmt.Sprintf("$%.2f", start_amount), got)
}

func TestNewShoppingCart_NegativeStartValue(t *testing.T) {
	shopping_cart, err := NewShoppingCart(-12.50)
	assert.Nil(t, shopping_cart)
	assert.EqualError(t, err, fmt.Sprintf("error: %v", ErrNegativeAmount))
}
