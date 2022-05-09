package main

import "strconv"

func Add(numbers string) int {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		numToReturn, _ := strconv.Atoi(numbers)
		return numToReturn
	}
	return 0
}

func main() {
}
