//go:build integration
// +build integration

package pointofsale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShoppingCart_SingleItem_ZeroAsStartAmount(t *testing.T) {
	sut, err := NewShoppingCart(0.00)
	if err != nil {
		t.Fatalf("failed to setup shopping_cart: %v", err)
	}
	_, err = sut.Scan("23456")
	if err != nil {
		t.Fatalf("failed to scan: %v", err)
	}
	got := sut.GetTotal()
	assert.Equal(t, "$12.50", got)
}
