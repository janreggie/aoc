package aoc2016

import (
	"crypto/md5"
	"strconv"
	"sync"
)

// checkHash checks the first n nibbles (half-octets) of the MD5 of Augend+string(Addend)
// and returns true and the entire hash if the first n nibbles are all zero.
// Otherwise it will return false and the hash.
func checkHash(n int, augend string, addend int) (bool, [16]byte) {
	if n < 0 || n > 31 {
		return false, [16]byte{}
	}
	hash := md5.Sum([]byte(augend + strconv.Itoa(addend)))

	// This is where counting gets weird.
	// Note that hash is a byte array of length 16,
	// and a nibble is half a byte, thus having 32 of them...
	// Let us consider the ii'th nibble as the LSBs and MSBs of hash[ii/2].
	for ii := 0; ii < n; ii++ {
		concern := hash[ii/2]
		if ii%2 == 0 { // MSBs
			if concern&0xF0 != 0 {
				return false, hash
			}
		} else { // LSBs
			if concern&0x0F != 0 {
				return false, hash
			}
		}
	}

	// once we know first n bits is zero...
	return true, hash
}

// Day05 solves the fifth day puzzle "How About a Nice Game of Chess?".
//
// Input
//
// A single line that represents the Door ID. For example:
//
//	ffykfhsq
//
// It is guaranteed that the input contains at least one line.
func Day05(input string) (answer1, answer2 string, err error) {

	firstPassword := make([]byte, 8)
	secondPassword := make([]byte, 8)

	addends := func() chan int {
		result := make(chan int)
		go func() {
			moar := 0
			// add moar to addend channel
			for {
				result <- moar
				moar++
			}
		}()
		return result
	}

	firstAddends := addends()
	firstPasswordChannel := make(chan byte, 8)
	addToFirst := func(toAdd byte) {
		if toAdd < 0x0A {
			firstPasswordChannel <- (0x30 + toAdd) // 0->'0', ...
		} else {
			firstPasswordChannel <- (0x57 + toAdd) // a->'a', ...
		}
	}

	secondPasswordChannel := make(chan struct {
		position uint8 // 0 to 7 only
		value    byte  // the character (0~9, a~f)
	}, 8)
	addToSecond := func(toAdd byte) {
		// toAdd has eight bits
		// toAdd[4:7] (MSBs) encode the actual value to be added
		// toAdd[0:3] (LSBs) encode the position
		result := struct {
			position uint8
			value    byte
		}{position: toAdd & 0x0F, value: (toAdd & 0xF0) >> 4}
		if result.value < 0x0A {
			result.value += 0x30 // 0->'0', ...
		} else {
			result.value += 0x57 // a->'a', ...
		}
		if result.position > 0x07 { // invalid!
			return // do nothing
		}
		secondPasswordChannel <- result
	}

	for ii := 0; ii < 4; ii++ {
		go func() {
			for {
				toAdd := <-firstAddends
				if isZero, result := checkHash(5, input, toAdd); isZero {

					// process (in accordance to Day 1)
					addToFirst(result[2])

					// process (in accordance to Day 2)
					addToSecond((result[3] & 0xF0) | result[2])
				}
			}
		}()
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for ii := 0; ii < 8; ii++ {
			firstPassword[ii] = <-firstPasswordChannel
		}
		wg.Done()
		for {
			<-firstPasswordChannel // just keep on taking...
		}
	}()

	go func() {
		completed := [8]bool{} // all false
	FORSECOND:
		for {
			out := <-secondPasswordChannel
			if !completed[out.position] {
				secondPassword[out.position] = out.value
				completed[out.position] = true
			}
			for _, done := range completed {
				if !done {
					continue FORSECOND
				}
			}
			wg.Done()
			break FORSECOND // we're done.
		}
		for {
			<-secondPasswordChannel
		}
	}()

	wg.Wait()
	answer1 = string(firstPassword)
	answer2 = string(secondPassword)
	return
}
