package main

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	TOO_SHORT      = "password must be at least 8 characters"
	TOO_FEW_DIGITS = "the password must contain at least 2 numbers"
)

func Validate(pass string) (bool, error) {
	r, _ := regexp.Compile(`[\d]`)
	numOfDigits := len(r.FindAllString(pass, -1))
	if len(pass) < 8 && numOfDigits < 2 {
		return false, fmt.Errorf("%v\n%v", TOO_SHORT, TOO_FEW_DIGITS)
	}
	if numOfDigits < 2 {
		return false, errors.New(TOO_FEW_DIGITS)
	}
	if len(pass) < 8 {
		return false, errors.New(TOO_SHORT)
	}
	return true, nil
}

func main() {
}
