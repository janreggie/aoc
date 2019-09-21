package aoc2015

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
)

// Day04 solves the fourth day puzzle
// "The Ideal Stocking Stuffer"
func Day04(scanner *bufio.Scanner) (string, string, error) {
	answer1, answer2 := "", ""
	// the entire scanner is read.
	input := ""
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		return "", "", errors.New("first line of file is empty")
	}
	// now what do we do with the input
	var index int64 = 1
	foundFive := false
	foundSix := false
	for {
		builtString := []byte(fmt.Sprintf("%v%v", input, index))
		if !foundFive {
			firstFive := fmt.Sprintf("%x", md5.Sum(builtString))[:5]
			if firstFive == "00000" {
				foundFive = true
				answer1 = strconv.FormatInt(index, 10)
			}
		}
		if !foundSix {
			firstSix := fmt.Sprintf("%x", md5.Sum(builtString))[:6]
			if firstSix == "000000" {
				foundSix = true
				answer2 = strconv.FormatInt(index, 10)
			}

		}
		if foundFive && foundSix == true {
			break
		}
		index++
	}
	return answer1, answer2, nil
}
