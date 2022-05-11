package main

import (
	"errors"
	"regexp"
	"strings"
)

const (
	TOO_SHORT         = "password must be at least 8 characters"
	TOO_FEW_DIGITS    = "the password must contain at least 2 numbers"
	NO_CAPITAL_LETTER = "password must contain at least one capital letter"
	NO_SPECIAL_CHARS  = "password must contain at least one special character"
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

	r, _ = regexp.Compile(`[A-Z]`)
	numOfCapitalLetters := len(r.FindAllString(pass, -1))
	if numOfCapitalLetters == 0 {
		validationErr = append(validationErr, NO_CAPITAL_LETTER)
	}

	r, _ = regexp.Compile(`[^A-Za-z0-9]`)
	numOfSpecialChars := len(r.FindAllString(pass, -1))
	if numOfSpecialChars == 0 {
		validationErr = append(validationErr, NO_SPECIAL_CHARS)
	}

	if len(validationErr) > 0 {
		return false, errors.New(strings.Join(validationErr, "\n"))
	}

	return true, nil
}

func main() {
}
