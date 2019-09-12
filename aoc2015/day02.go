package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Day02 solves the second day
func Day02(scanner *bufio.Scanner) (string, string, error) {
	answer1, answer2 := "", ""
	splitDims := make([]string, 3)
	actualDims := make([]int64, 3)
	totalPaper := int64(0)
	totalRibbon := int64(0)
	// some helper functions
	// it is assumed that dims is of length 3
	// max function for maximum of a slice
	max := func(list []int64) int64 {
		currentMax := list[0]
		for i := range list {
			if list[i] > currentMax {
				currentMax = list[i]
			}
		}
		return currentMax
	}
	// surfaceArea function for total surface area
	surfaceArea := func(dims []int64) int64 {
		return 2 * (dims[0]*dims[1] + dims[1]*dims[2] + dims[2]*dims[0])
	}
	// slack function for the little extra paper
	slack := func(dims []int64) int64 {
		result := int64(1)
		for _, v := range dims {
			result *= v
		}
		result /= max(dims)
		return result
	}
	// ribbon for amount of ribbon to be used
	ribbon := func(dims []int64) int64 {
		result := int64(0)
		for _, v := range dims {
			result += 2 * v
		}
		result -= 2 * max(dims)
		return result
	}
	// bow function for total volume
	bow := func(dims []int64) int64 {
		return dims[0] * dims[1] * dims[2]
	}
	for scanner.Scan() {
		readDims := scanner.Text()
		splitDims = strings.Split(readDims, "x")
		if len(splitDims) != 3 {
			// throw an error
			return "", "", fmt.Errorf("invalid dimensions: %v", readDims)
		}
		// Atoi each dimension
		for ind := range actualDims {
			convt, err := strconv.Atoi(splitDims[ind])
			if err != nil {
				return "", "", fmt.Errorf("not an integer: %v (%v)", readDims, splitDims[ind])
			}
			actualDims[ind] = int64(convt)
		}
		totalPaper += surfaceArea(actualDims) + slack(actualDims)
		totalRibbon += ribbon(actualDims) + bow(actualDims)
	}
	answer1 = strconv.FormatInt(totalPaper, 10)
	answer2 = strconv.FormatInt(totalRibbon, 10)
	return answer1, answer2, nil
}
