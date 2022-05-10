package main

import (
	"errors"
	"unicode"
)

const TOO_SHORT = "password must be at least 8 characters"

func Validate(pass string) (bool, error) {
	if len(pass) < 8 {
		return false, errors.New(TOO_SHORT)
	}
	numbersInPass := 0
	for _, v := range pass {
		if unicode.IsDigit(v) {
			numbersInPass++
		}
	}
	if numbersInPass < 2 {
		return false, errors.New("the password must contain at least 2 numbers")
	}
	return true, nil
}

func main() {
}
