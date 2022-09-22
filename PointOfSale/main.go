package pointofsale

import "fmt"

const (
	ErrBarcodeNotFound = "barcode not found"
	ErrBarcodeEmpty    = "empty barcode"
	ErrNegativeAmount  = "amount cannot be less than zero"
)

var products = map[string]float64{
	"12345": 7.25,
	"23456": 12.5,
}

type ShoppingCart struct {
	total float64
}

func NewShoppingCart(starting_total float64) (*ShoppingCart, error) {
	if starting_total < 0.00 {
		return nil, fmt.Errorf("error: %v", ErrNegativeAmount)
	}
	return &ShoppingCart{total: starting_total}, nil
}

func (s *ShoppingCart) GetTotal() string {
	return fmt.Sprintf("$%.2f", s.total)
}

func (s *ShoppingCart) Scan(barcode string) (price string, err error) {
	if barcode == "" {
		return fmt.Sprintf("$%.2f", 0.0), fmt.Errorf("error: %v", ErrBarcodeEmpty)
	}
	if price, found := products[barcode]; found {
		s.total += price
		return fmt.Sprintf("$%.2f", price), nil
	}
	return fmt.Sprintf("$%.2f", 0.0), fmt.Errorf("error: %v", ErrBarcodeNotFound)
}
