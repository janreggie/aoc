package aoc2015

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

// lookAndSay looks at the string and returns a string
// containing how many times a rune has repeated itself following said rune.
// For instance:
// lookAndSay("1") returns "11" (one '1')
// lookAndSay("111221") returns "312211" (three '1', two '2', one '1')
// If string is empty it will return an empty string.
func lookAndSay(input string) string {
	if input == "" {
		return ""
	}

	count := 0
	last := '\x00'
	hasStarted := false
	var total strings.Builder

	for _, current := range input {
		if !hasStarted {
			count = 1
			last = current
			hasStarted = true
			continue
		}
		if current != last {
			total.WriteString(strconv.Itoa(count))
			total.WriteRune(last)
			count = 1
			last = current
			continue
		}
		count++
	}
	total.WriteString(strconv.Itoa(count))
	total.WriteRune(last)

	return total.String()
}

// Day10 solves the tenth day puzzle
// "Elves Look, Elves Say"
func Day10(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// the entire scanner is read.
	input := ""
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		err = errors.New("first line of file is empty")
		return
	}

	// iterate forty times
	for ii := 0; ii < 40; ii++ {
		then := time.Now()
		input = lookAndSay(input)
		now := time.Now()
		glog.Infof("%v took %v nanoseconds", ii+1, now.Sub(then).Nanoseconds())
	}
	answer1 = strconv.Itoa(len(input))

	// now continue this 10 more times
	for ii := 0; ii < 10; ii++ {
		then := time.Now()
		input = lookAndSay(input)
		now := time.Now()
		glog.Infof("%v took %v nanoseconds", ii+41, now.Sub(then).Nanoseconds())
	}
	answer2 = strconv.Itoa(len(input))

	return
}