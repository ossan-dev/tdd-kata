package main

import "errors"

func Validate(pass string) (bool, error) {
	return false, errors.New("password must be at least 8 characters")
}

func main() {
}
