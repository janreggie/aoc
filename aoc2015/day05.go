package aoc2015

import (
	"bufio"
	"fmt"
	"strconv"
)

// Day05 solves the fifth day puzzle
// "Doesn't He Have Intern-Elves For This?"
func Day05(scanner *bufio.Scanner) (string, string, error) {
	answer1, answer2 := "", ""
	var totalNiceOne, totalNiceTwo int64
	// Is it more worth to check each ind := range currentString?
	// Not necessarily since it will run the same number of comparisons
	// O(len(currentString)*3) == O(3*len(currentString))

	// checkNiceOne checks whether string is "nice" according to Part One criteria
	checkNiceOne := func(input string) bool {
		// checkThreeVowels checks whether input has at least three vowels
		// and returns true if so.
		checkThreeVowels := func(input string) bool {
			checkVowel := func(char byte) bool {
				return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
			}
			numberOfVowels := 0
			for ind := range input {
				if checkVowel(input[ind]) {
					numberOfVowels++
				}
				if numberOfVowels >= 3 {
					return true
				}
			}
			return false
		}
		// checkTwiceInARow checks whether input has a letter appearing twice in a row
		// and returns true if so.
		checkTwiceInARow := func(input string) bool {
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
		checkSanity := func(input string) bool {
			invalidStrings := make(map[byte]byte)
			invalidStrings['a'] = 'b'
			invalidStrings['c'] = 'd'
			invalidStrings['p'] = 'q'
			invalidStrings['x'] = 'y'
			for ind := range input {
				if ind == len(input)-1 {
					break
				}
				if invalidStrings[input[ind]] == input[ind+1] {
					return false
				}
			}
			return true
		}

		return checkSanity(input) &&
			checkThreeVowels(input) &&
			checkTwiceInARow(input)
	}

	// checkNiceTwo checks whether string is "nice" according to Part Two criteria
	checkNiceTwo := func(input string) bool {
		// checkRepeatPair checks whether a pair of any two letters
		// appears twice in the string without overlapping
		checkRepeatPair := func(input string) bool {
			// checkPatternExists checks whether pattern is in text
			checkPatternExists := func(text, pattern string) bool {
				indexText := 0
				indexPat := 0
				for indexText < len(text) {
					// compare text[indexText] and pattern[indexPat]
					if text[indexText] != pattern[indexPat] {
						indexPat = 0
					}
					if text[indexText] == pattern[indexPat] {
						indexPat++
					}
					if indexPat == len(pattern) {
						return true
					}
					indexText++
				}
				return false
			}
			for ind := 0; ind < len(input)-1; ind++ {
				substr := input[ind : ind+2] // length 2
				if checkPatternExists(input[ind+2:], substr) {
					return true
				}
			}
			return false
		}
		// checkRepeatChar checks whether at least one letter
		// repeats with exactly one letter between them
		checkRepeatChar := func(input string) bool {
			for ind := 0; ind < len(input)-2; ind++ {
				// check input[ind] and input[ind+2]
				if input[ind] == input[ind+2] {
					return true
				}
			}
			return false
		}

		return checkRepeatChar(input) &&
			checkRepeatPair(input)
	}

	for scanner.Scan() {
		currentString := scanner.Text()
		if checkNiceOne(currentString) {
			totalNiceOne++
		}
		if checkNiceTwo(currentString) {
			totalNiceTwo++
		}
		if err := scanner.Err(); err != nil {
			return "", "", fmt.Errorf("scanner error: %v", err)
		}
	}
	answer1 = strconv.FormatInt(totalNiceOne, 10)
	answer2 = strconv.FormatInt(totalNiceTwo, 10)
	return answer1, answer2, nil
}
