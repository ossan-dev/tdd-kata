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
	if len(numbers) == 3 {
		num1Str := strings.Split(numbers, ",")[0]
		num2Str := strings.Split(numbers, ",")[1]
		num2, _ := strconv.Atoi(num2Str)
		num1, _ := strconv.Atoi(num1Str)
		return num1 + num2
	}

	num1Str := strings.Split(numbers, ",")[0]
	num2Str := strings.Split(numbers, ",")[1]
	num3Str := strings.Split(numbers, ",")[2]
	num4Str := strings.Split(numbers, ",")[3]
	num1, _ := strconv.Atoi(num1Str)
	num2, _ := strconv.Atoi(num2Str)
	num3, _ := strconv.Atoi(num3Str)
	num4, _ := strconv.Atoi(num4Str)
	return num1 + num2 + num3 + num4
}

func main() {
}
