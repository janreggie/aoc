package aoc2019

import (
	"fmt"
	"strconv"

	aoc "github.com/janreggie/aoc/internal"
)

type spaceImage [6][25]int8 // 6 tall, 25 wide

func newSpaceImage(raw string) (sp spaceImage, err error) {
	if len(raw) != 150 {
		err = fmt.Errorf("raw is length %v, should be 150", len(raw))
		return
	}
	for ii := 0; ii < 6; ii++ {
		for jj := 0; jj < 25; jj++ {
			switch val := raw[ii*25+jj]; val {
			case '0':
				sp[ii][jj] = 0
			case '1':
				sp[ii][jj] = 1
			case '2':
				sp[ii][jj] = 2
			default:
				err = fmt.Errorf("invalid byte %v on position %v", string(val), ii*25+jj)
				return
			}
		}
	}
	return
}

func (sp *spaceImage) count() (count [3]int) {
	for ii := range sp {
		for jj := range sp[ii] {
			switch sp[ii][jj] {
			case 0:
				count[0]++
			case 1:
				count[1]++
			case 2:
				count[2]++
			}
		}
	}
	return
}

// Day08 solves the eighth day puzzle "Space Image Format".
//
// Input
//
// A string of length 15000, which only consists of '0', '1', and '2'.
//
// Notes
//
// The second answer should not be used as-is,
// but requires the interpretation of a human.
// The output will resemble pixel art made out of blocks,
// which spells out a word that the user must input manually.
// My output is as follows:
//
// 	 ██ ██  ███  ██ ██ █ ██ █
// 	 ██ █ ██ █ ██ █ ██ █ ██ █
// 	 ██ █ ████ ████ ██ █    █
// 	 ██ █ █  █ ████ ██ █ ██ █
// 	 ██ █ ██ █ ██ █ ██ █ ██ █
// 	█  ███   ██  ███  ██ ██ █
//
// It may be hard to read but it does spell out UGCUH.
func Day08(input string) (answer1, answer2 string, err error) {
	allLayers := make([]spaceImage, 0)

	for _, line := range aoc.SplitLines(input) {
		// split for every 150 bytes
		for i := 0; (i+1)*150 <= len(line); i++ {
			sp, e := newSpaceImage(line[i*150 : (i+1)*150])
			if e != nil {
				err = fmt.Errorf("can't create image from %v: %v", line, e)
				return
			}
			allLayers = append(allLayers, sp)
		}
	}

	fewestZeroes := allLayers[0].count()[0] // got it? good.
	for _, layer := range allLayers {
		if counts := layer.count(); counts[0] < fewestZeroes {
			fewestZeroes = counts[0]
			answer1 = strconv.Itoa(counts[1] * counts[2])
		}
	}
	// now for the second answer...
	answer2 += "the following:\n"
	for ii := 0; ii < 6; ii++ {
		for jj := 0; jj < 25; jj++ {
			for _, layer := range allLayers {
				if layer[ii][jj] == 0 {
					answer2 += "█"
					break
				} else if layer[ii][jj] == 1 {
					answer2 += " "
					break
				}
			}
		}
		answer2 += "\n"
	}
	return
}
