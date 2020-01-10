package aoc2015

import (
	"bufio"
	"strconv"

	"github.com/antonholmquist/jason"
	"github.com/pkg/errors"
)

// extractSum returns the sum of all numbers (Int64) in an arbitrary value.
// If the value is an int64 it will return that immediately.
// If the value is an Object it will return the extracted sum of its value
// If the value is an Array it will recursively get the sum of all numbers in that array.
// If the value is anything other than those it will return zero.
func extractSum(value *jason.Value) int64 {
	// check int64
	if result, err := value.Int64(); err == nil {
		return result
	}

	// check Object
	if object, err := value.Object(); err == nil {
		var result int64
		for _, vv := range object.Map() {
			result += extractSum(vv)
		}
		return result
	}

	// check Array
	if array, err := value.Array(); err == nil {
		var result int64
		for _, vv := range array {
			result += extractSum(vv)
		}
		return result
	}

	// if all else fails
	return 0
}

// extractSumButRed is like extractSum but with the following clause:
// If value is an Object and one of its values is the string "red",
// it will return zero.
func extractSumButRed(value *jason.Value) int64 {
	// check int64
	if result, err := value.Int64(); err == nil {
		return result
	}

	// check Object
	if object, err := value.Object(); err == nil {
		var result int64
		for _, vv := range object.Map() {
			// check first if vv is a string
			if str, err := vv.String(); err == nil && str == "red" {
				return 0
			}
			result += extractSumButRed(vv)
		}
		return result
	}

	// check Array
	if array, err := value.Array(); err == nil {
		var result int64
		for _, vv := range array {
			result += extractSumButRed(vv)
		}
		return result
	}

	// if all else fails
	return 0
}

// Day12 solves the twelfth day problem
// "JSAbacusFramework.io"
func Day12(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// let's read everything first...
	result := ""
	if !scanner.Scan() {
		err = errors.Wrap(scanner.Err(), "could not scan text; empty?")
		return
	}
	result = scanner.Text() // just scan one line of text
	data, err := jason.NewObjectFromBytes([]byte(result))

	// but how do we read the data?
	var totalSum int64
	var antiCommunistSum int64 // answer2 (because not red)
	for _, vv := range data.Map() {
		totalSum += extractSum(vv)
		antiCommunistSum += extractSumButRed(vv)
	}
	answer1 = strconv.FormatInt(totalSum, 10)
	answer2 = strconv.FormatInt(antiCommunistSum, 10)

	return
}
