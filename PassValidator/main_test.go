package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_ShouldGetErr_WhenPassLenIsLessThanEight(t *testing.T) {
	pass := "1A4.567"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, TOO_SHORT, err.Error())
}

func TestValidate_ShouldGetErr_WhenNotContainsAtLeastTwoNumbers(t *testing.T) {
	pass := "aBcd:efgh1"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, TOO_FEW_DIGITS, err.Error())
}

func TestValidate_ShouldGetErr_WhenIsLessThanEightAndNotContainsAtLeastTwoNumbers(t *testing.T) {
	pass := "ab\\2Cd"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, fmt.Sprintf("%v\n%v", TOO_SHORT, TOO_FEW_DIGITS), err.Error())
}

func TestValidate_ShouldGetErr_WhenNoCapitalLetterIsPresent(t *testing.T) {
	pass := "aa3?!ab2cd"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, NO_CAPITAL_LETTER, err.Error())
}

func TestValidate_ShouldGetErr_WhenNoSpecialCharIsPresent(t *testing.T) {
	pass := "aa3ab2cdA"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, NO_SPECIAL_CHARS, err.Error())
}
