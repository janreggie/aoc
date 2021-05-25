package aoc

import (
	"strconv"

	"github.com/pkg/errors"
)

// StringsToInts turns a slice of strings to ints.
// Returns an error if one of the strings cannot be parsed.
func StringsToInts(strs []string) ([]int, error) {

	result := make([]int, len(strs))
	for ii := range result {
		var err error
		result[ii], err = strconv.Atoi(strs[ii])
		if err != nil {
			return nil, errors.Errorf("could not interpret line `%v` as int", strs[ii])
		}
	}
	return result, nil
}
