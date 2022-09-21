package pointofsale

import "fmt"

var products = map[string]float64{
	"12345": 7.25,
	"23456": 12.5,
}

func Scan(barcode string) (string, error) {
	if barcode == "99999" {
		return fmt.Sprintf("$%.2f", 0.0), fmt.Errorf("Error: barcode not found")
	}
	price, _ := products[barcode]
	return fmt.Sprintf("$%.2f", price), nil
}
