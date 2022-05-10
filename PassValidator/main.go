package main

import "errors"

const TOO_SHORT = "password must be at least 8 characters"

func Validate(pass string) (bool, error) {
	if len(pass) < 8 {
		return false, errors.New(TOO_SHORT)
	}
	return true, nil
}

func main() {
}
