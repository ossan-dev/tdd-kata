package pointofsale

import "fmt"

const ErrBarcodeNotFound = "barcode not found"

var products = map[string]float64{
	"12345": 7.25,
	"23456": 12.5,
}

func Scan(barcode string) (price string, err error) {
	if price, found := products[barcode]; found {
		return fmt.Sprintf("$%.2f", price), nil
	}
	return fmt.Sprintf("$%.2f", 0.0), fmt.Errorf("error: %v", ErrBarcodeNotFound)
}
