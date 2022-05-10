package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_ShouldGetErr_WhenPassLenIsLessThanEight(t *testing.T) {
	pass := "134567"

	got, err := Validate(pass)

	assert.Equal(t, false, got)
	assert.Equal(t, "password must be at least 8 characters", err.Error())
}
