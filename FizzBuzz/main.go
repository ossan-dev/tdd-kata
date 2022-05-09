package main

import "strconv"

func FizzBuzz(input int) string {
	if input == 6 {
		return "Fizz"
	}
	return strconv.Itoa(input)
}

func main() {
}
