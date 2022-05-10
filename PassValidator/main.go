package main

import (
	"errors"
	"regexp"
)

const (
	TOO_SHORT      = "password must be at least 8 characters"
	TOO_FEW_DIGITS = "the password must contain at least 2 numbers"
)

func Validate(pass string) (bool, error) {
	if len(pass) < 8 {
		return false, errors.New(TOO_SHORT)
	}

	r, _ := regexp.Compile(`[\d]`)
	numOfDigits := len(r.FindAllString(pass, -1))
	if numOfDigits < 2 {
		return false, errors.New(TOO_FEW_DIGITS)
	}
	return true, nil
}

func main() {
}
