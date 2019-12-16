package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
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

// Day08 solves the sixth day puzzle
// "Space Image Format"
func Day08(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	allLayers := make([]spaceImage, 0)
	// scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// 	// This is bufio.ScanBytes but with 150 bytes at a time
	// 	// For some reason it doesn't work...
	// 	if atEOF && len(data) == 0 {
	// 		return 0, nil, nil
	// 	}
	// 	if len(data) < 150 {
	// 		// glog.Infof("data is %v; length %v", data, len(data))
	// 		return len(data), data[0:len(data)], nil
	// 	}
	// 	return 150, data[0:150], nil
	// })
	// scanner.Buffer([]byte{}, 100000)
	for scanner.Scan() {
		raw := scanner.Text()
		// split for every 150 bytes
		for i := 0; (i+1)*150 <= len(raw); i++ {
			sp, e := newSpaceImage(raw[i*150 : (i+1)*150])
			if e != nil {
				err = fmt.Errorf("can't create image from %v: %v", raw, e)
				return
			}
			allLayers = append(allLayers, sp)
		}
		// For custom SplitFunc
		// sp, err := newSpaceImage(raw)
		// if err != nil {
		// 	e = fmt.Errorf("can't create image from %v: %v", raw, err)
		// 	return
		// }
		// allLayers = append(allLayers, sp)
		// glog.Infof("current: %v; count: %v", raw, len(allLayers))
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
