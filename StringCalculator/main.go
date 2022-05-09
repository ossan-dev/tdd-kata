package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(numbers string) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	}
	if len(numbers) == 1 {
		numToReturn, _ := strconv.Atoi(numbers)
		return numToReturn, nil
	}

	sum := 0
	delimiter := ","
	for _, line := range strings.Split(numbers, "\n") {
		if line[0] == '/' {
			delimiter = string(line[2:])
			continue
		}
		if _, err := strconv.Atoi(string(line[len(line)-1])); err != nil {
			return 0, fmt.Errorf("the line %q terminates with an unallowed character", line)
		}
		for _, number := range strings.Split(line, delimiter) {
			numberToAdd, _ := strconv.Atoi(number)
			sum += numberToAdd
		}
	}
	return sum, nil
}

func main() {
}
