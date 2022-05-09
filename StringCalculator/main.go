package main

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		numToReturn, _ := strconv.Atoi(numbers)
		return numToReturn
	}

	sum := 0
	for _, line := range strings.Split(numbers, "\n") {
		for _, number := range strings.Split(line, ",") {
			numberToAdd, _ := strconv.Atoi(number)
			sum += numberToAdd
		}
	}
	return sum
}

func main() {
}
