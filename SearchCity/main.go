package searchcity

import "strings"

var cities = []string{"Paris", "Budapest", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam", "Vienna", "Sydney", "New York City", "London", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul"}

func SearchCity(input string) []string {
	if input == "*" {
		return cities
	}
	if len(input) < 2 {
		return []string{}
	}
	input = strings.ToLower(input)
	var result []string
	for _, v := range cities {
		if strings.Contains(strings.ToLower(v), input) {
			result = append(result, v)
		}
	}
	return result
}
