package main

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

const (
	TOO_SHORT      = "password must be at least 8 characters"
	TOO_FEW_DIGITS = "the password must contain at least 2 numbers"
)

func Validate(pass string) (bool, error) {
	validationErr := []string{}

	if len(pass) < 8 {
		validationErr = append(validationErr, TOO_SHORT)
	}

	r, _ := regexp.Compile(`[\d]`)
	numOfDigits := len(r.FindAllString(pass, -1))
	if numOfDigits < 2 {
		validationErr = append(validationErr, TOO_FEW_DIGITS)
	}

	if len(validationErr) > 0 {
		return false, errors.New(strings.Join(validationErr, "\n"))
	}

	numOfCapitalLetters := 0
	for _, v := range pass {
		if unicode.IsUpper(v) {
			numOfCapitalLetters++
		}
	}
	if numOfCapitalLetters == 0 {
		return false, errors.New("password must contain at least one capital letter")
	}

	return true, nil
}

func main() {
}
