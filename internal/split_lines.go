package aoc

import "strings"

// SplitLines splits an input to its lines except the last few ones if they are empty
func SplitLines(input string) []string {
	result := strings.Split(input, "\n")
	for ii := len(result) - 1; ii >= 0; ii-- {
		if result[ii] == "" {
			result = result[:ii]
		}
	}
	return result
}
