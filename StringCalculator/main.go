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
	num1Str := strings.Split(numbers, ",")[0]
	num2Str := strings.Split(numbers, ",")[1]
	num1, _ := strconv.Atoi(num1Str)
	num2, _ := strconv.Atoi(num2Str)
	return num1 + num2
}

func main() {
}
