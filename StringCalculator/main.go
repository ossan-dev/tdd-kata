package main

import (
	"fmt"
	"regexp"
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
	delimiterPassed := false
	for _, line := range strings.Split(numbers, "\n") {
		if line[0] == '/' {
			delimiter = string(line[2:])
			delimiterPassed = true
			continue
		}
		if _, err := strconv.Atoi(string(line[len(line)-1])); err != nil {
			return 0, fmt.Errorf("the line %q terminates with an unallowed character", line)
		}
		for _, number := range strings.Split(line, delimiter) {
			numberToAdd, err := strconv.Atoi(number)
			if err != nil {
				r, _ := regexp.Compile(`[\D]`)
				charNotAllowed := r.FindString(number)
				indexOfCharNotAllowed := -1
				if delimiterPassed {
					stringToSearch := numbers[4:]
					indexOfCharNotAllowed = strings.Index(stringToSearch, charNotAllowed)
					return 0, fmt.Errorf("'%v' expected but '%v' found at position %d.", delimiter, charNotAllowed, indexOfCharNotAllowed)
				}
				indexOfCharNotAllowed = strings.Index(numbers, charNotAllowed)
				return 0, fmt.Errorf("'%v' expected but '%v' found at position %d.", delimiter, charNotAllowed, indexOfCharNotAllowed)
			}
			sum += numberToAdd
		}
	}
	return sum, nil
}

func main() {
}
