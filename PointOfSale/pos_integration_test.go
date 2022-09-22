//go:build integration
// +build integration

package pointofsale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCart_SingleItem_ZeroAsStartAmount(t *testing.T) {
	shopping_cart, err := NewShoppingCart(0.00)
	if err != nil {
		t.Fatalf("failed to setup shopping_cart: %v", err)
	}
	_, err = shopping_cart.Scan("23456")
	if err != nil {
		t.Fatalf("failed to scan: %v", err)
	}
	got := shopping_cart.GetTotal()
	assert.Equal(t, "$12.50", got)
}
