package aoc2015

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// checkHash checks the first n nibbles (half-octets) of the MD5 of Augend+string(Addend)
// and returns true if they are all zero.
func checkHash(n int, augend string, addend int) bool {
	hash := md5.Sum([]byte(augend + strconv.Itoa(addend)))

	// This is where counting gets weird.
	// Note that hash is a byte array of length 16,
	// and a nibble is half a byte, thus having 32 of them...
	// Let us consider the ii'th nibble as the LSBs and MSBs of hash[ii/2].
	for ii := 0; ii < n; ii++ {
		concern := hash[ii/2]
		if ii%2 == 0 { // MSBs
			if concern&0xF0 != 0 {
				return false
			}
		} else { // LSBs
			if concern&0x0F != 0 {
				return false
			}
		}
	}

	// once we know first n bits is zero...
	return true
}

// Day04 solves the fourth day puzzle "The Ideal Stocking Stuffer".
//
// Input
//
// A string of length 8 containing only the lowercase
// letters 'a' to 'z'. For example:
//
//	iwrupvqb
//
// Because this is more or less a brute-force solution,
// it may take a long time before returning an answer.
func Day04(input string) (answer1, answer2 string, err error) {

	// suppose we have an addends channel
	// which contains all the possible addends for input
	// i.e., input+<-addned
	processes := 8 // how many processes in parallel?
	addends := func() chan int {
		result := make(chan int, processes*2)
		go func() {
			moar := 1
			// add moar to addend channel
			for {
				result <- moar
				moar++
			}
		}()
		return result
	}

	resultFive, resultSix := make(chan int), make(chan int) // result for 5 and 6 zeroes
	addendFive, addendSix := addends(), addends()
	foundFive, foundSix := make(chan struct{}), make(chan struct{}) // will close once results are found

	evaluateFive := func() {
		for {
			addend := <-addendFive
			select {
			case <-foundFive:
				return
			default:
				if checkHash(5, input, addend) {
					resultFive <- addend
					close(foundFive) // will always send default
					return
				}
			}
		}
	}

	evaluateSix := func() {
		for {
			addend := <-addendSix
			select {
			case <-foundSix:
				return
			default:
				if checkHash(6, input, addend) {
					resultSix <- addend
					close(foundSix) // will always send default
					return
				}
			}
		}
	}

	for ii := 0; ii < processes; ii++ {
		go evaluateFive()
		go evaluateSix()
	}

	answer1 = strconv.Itoa(<-resultFive)
	answer2 = strconv.Itoa(<-resultSix)

	return
}

// Day04ST solves the fourth day puzzle "The Ideal Stocking Stuffer"
// but is a single-threaded solution.
func Day04ST(input string) (answer1, answer2 string, err error) {

	addend := 1        // to be appended to input
	foundFive := false // have we found five zeroes?
	foundSix := false  // have we found six zeroes?
	for {
		builtString := []byte(fmt.Sprintf("%v%v", input, addend))
		if !foundFive {
			firstFive := fmt.Sprintf("%x", md5.Sum(builtString))[:5]
			if firstFive == "00000" {
				foundFive = true
				answer1 = strconv.Itoa(addend)
			}
		}
		if !foundSix {
			firstSix := fmt.Sprintf("%x", md5.Sum(builtString))[:6]
			if firstSix == "000000" {
				foundSix = true
				answer2 = strconv.Itoa(addend)
			}

		}

		if foundFive && foundSix {
			break
		}

		addend++
	}
	return
}
