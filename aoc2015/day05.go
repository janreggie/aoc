package aoc2015

import (
	"strconv"
	"strings"
)

// checkThreeVowels returns true if there are at least three vowels in input
func checkThreeVowels(input string) bool {

	checkVowel := func(char byte) bool {
		return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
	}

	vowelCount := 0
	for ind := range input {
		if checkVowel(input[ind]) {
			vowelCount++
		}
		if vowelCount >= 3 {
			return true
		}
	}

	return false
}

// checkTwiceInARow returns true if two characters appear twice in a row in input
func checkTwiceInARow(input string) bool {

	for ind := range input {
		if ind == len(input)-1 {
			break
		}
		if input[ind] == input[ind+1] {
			return true
		}
	}

	return false
}

// checkSanity checks for the strings
// ab, cd, pq, or xy
// and returns false if found
func checkSanity(input string) bool {

	for ind := range input {
		if ind == len(input)-1 {
			break
		}

		currentAndNext := input[ind : ind+2]
		if currentAndNext == "ab" ||
			currentAndNext == "cd" ||
			currentAndNext == "pq" ||
			currentAndNext == "xy" {
			return false
		}
	}

	return true
}

// checkRepeatPair checks whether any pair of any two letters
// appears twice in the string without overlapping
func checkRepeatPair(input string) bool {

	// checkPatternExists checks whether pattern is in text
	checkPatternExists := func(text, pattern string) bool {
		indexText := 0
		indexPat := 0
		for indexText < len(text) {

			if text[indexText] != pattern[indexPat] { // Reset counter
				indexPat = 0
			}
			if text[indexText] == pattern[indexPat] { // But see if this one matches
				indexPat++
			}
			if indexPat == len(pattern) {
				return true
			}
			indexText++
		}
		return false
	}

	for ii := 0; ii < len(input)-1; ii++ {
		substr := input[ii : ii+2] // length 2
		if checkPatternExists(input[ii+2:], substr) {
			return true
		}
	}

	return false
}

// checkRepeatChar checks whether at least one letter
// repeats with exactly one letter between them
func checkRepeatChar(input string) bool {
	for ind := 0; ind < len(input)-2; ind++ {
		// check input[ind] and input[ind+2]
		if input[ind] == input[ind+2] {
			return true
		}
	}
	return false
}

// checkNiceOne checks whether string is "nice" according to Part One criteria
func checkNiceOne(input string) bool {
	return checkSanity(input) &&
		checkThreeVowels(input) &&
		checkTwiceInARow(input)
}

// checkNiceTwo checks whether string is "nice" according to Part Two criteria
func checkNiceTwo(input string) bool {
	return checkRepeatChar(input) &&
		checkRepeatPair(input)
}

// Day05 solves the fifth day puzzle "Doesn't He Have Intern-Elves For This?".
//
// Input
//
// A file containing 1000 lines each of length 16 only containing the
// lowercase letters 'a' to 'z'. For example:
//
//	zvnvprxxjkhnkkpq
//	nejwxbhjxxdbenid
//	chryiccsebdbcnkc
//	guoeefaeafhlgvxh
//	nzapxrfrrqhsingx
//	mkzvquzvqvwsejqs
//	kozmlmbchydtxeeo
//	keylygnoqhmfzrfp
//	srwzoxccndoxylxe
//	uqjzalppoorosxxo
//	potmkinyuqxsfdfw
//	qkkwrhpbhypxhiun
//	wgfvnogarjmdbxyh
//
func Day05(input string) (answer1, answer2 string, err error) {

	// Goroutines are not worth the trouble here.

	var totalNiceOne, totalNiceTwo int64
	for _, currentString := range strings.Fields(input) {
		if checkNiceOne(currentString) {
			totalNiceOne++
		}
		if checkNiceTwo(currentString) {
			totalNiceTwo++
		}
	}

	answer1 = strconv.FormatInt(totalNiceOne, 10)
	answer2 = strconv.FormatInt(totalNiceTwo, 10)
	return
}
