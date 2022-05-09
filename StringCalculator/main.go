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
	for _, number := range strings.Split(numbers, ",") {
		numberToAdd, _ := strconv.Atoi(number)
		sum += numberToAdd
	}
	return sum
}

func main() {
}
