package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_ShouldGetErr_WhenPassLenIsLessThanEight(t *testing.T) {
	pass := "134567"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, TOO_SHORT, err.Error())
}

func TestValidate_ShouldGetErr_WhenNotContainsAtLeastTwoNumbers(t *testing.T) {
	pass := "abcdefgh1"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, "the password must contain at least 2 numbers", err.Error())
}
