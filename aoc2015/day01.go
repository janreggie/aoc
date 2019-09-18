package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/golang/glog"
)

// Day01 solves the first day puzzle
// "Not Quite Lisp"
func Day01(scanner *bufio.Scanner) (string, string, error) {
	answer1, answer2 := "", ""
	var err error
	currentFloor := int64(0)
	position := int64(1)
	passedBasement := false
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		readFloor := scanner.Text()
		if readFloor != "(" && readFloor != ")" {
			glog.Warningf("Unexpected bit: %v", readFloor)
			continue
		}
		if readFloor == "(" {
			currentFloor++
		} else if readFloor == ")" {
			currentFloor--
		}
		if currentFloor == -1 && !passedBasement {
			answer2 = strconv.FormatInt(position, 10)
			passedBasement = true
		}
		position++
		if err := scanner.Err(); err != nil {
			return "", "", fmt.Errorf("scanner error: %v", err)
		}
	}
	answer1 = strconv.FormatInt(currentFloor, 10)
	return answer1, answer2, err
}
