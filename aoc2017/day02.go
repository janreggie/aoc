package aoc2017

import (
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// spreadsheet represents the spreadsheet for Day02.
//
type spreadsheet struct {
	nums [][]int64
}

// newSpreadsheet returns a spreadsheet using the input provided
func newSpreadsheet(input string) (*spreadsheet, error) {

	rows := aoc.SplitLines(input)
	nums := make([][]int64, len(rows))

	for ii, row := range rows {
		entries := strings.Fields(row)
		nums[ii] = make([]int64, len(entries))
		for jj, entry := range entries {
			num, err := strconv.ParseInt(entry, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, "could not parse input at row %v col %v", ii, jj)
			}
			nums[ii][jj] = num
		}
	}

	return &spreadsheet{nums: nums}, nil
}

// checksum returns the checksum of a given spreadsheet
func (sheet *spreadsheet) checksum() int64 {

	var result int64

	for _, row := range sheet.nums {
		if len(row) == 0 {
			continue
		}
		min, max := row[0], row[0]
		for _, elem := range row {
			if elem < min {
				min = elem
			}
			if elem > max {
				max = elem
			}
		}
		result += (max - min)
	}

	return result
}

// quotientSum returns the "quotient sum" of a given spreadsheet.
// A "quotient sum" is the sum of each "quotient" per row.
// A row's "quotient" is the quotient of the only *two* numbers in that row that evenly divide each other.
//
func (sheet *spreadsheet) quotientSum() int64 {

	var result int64

FOREACHROW:
	for _, row := range sheet.nums {
		for ii := 0; ii < len(row); ii++ {
			for jj := ii + 1; jj < len(row); jj++ {

				n1, n2 := row[ii], row[jj]
				if n1 > n2 {
					n1, n2 = n2, n1
				}
				if n1 == 0 {
					continue
				}
				quot := n2 / n1
				if quot*n1 == n2 {
					result += quot
					continue FOREACHROW
				}
			}
		}
	}

	return result
}

// Day02 solves the second day puzzle
// "Corruption Checksum".
//
// Input
//
// A set of numbers displayed in a grid separated by tabs. For example:
//
// 	116	1259	1045	679	1334	157	277	1217	218	641	1089	136	247	1195	239	834
// 	269	1751	732	3016	260	6440	5773	4677	306	230	6928	7182	231	2942	2738	3617
// 	644	128	89	361	530	97	35	604	535	297	599	121	567	106	114	480
// 	105	408	120	363	430	102	137	283	123	258	19	101	181	477	463	279
// 	873	116	840	105	285	238	540	22	117	125	699	953	920	106	113	259
//
// All numbers are at most four digits long.
//
func Day02(input string) (answer1, answer2 string, err error) {

	sheet, err := newSpreadsheet(input)
	if err != nil {
		err = errors.Wrapf(err, "could not instantiate spreadsheet from input")
		return
	}

	answer1 = strconv.FormatInt(sheet.checksum(), 10)
	answer2 = strconv.FormatInt(sheet.quotientSum(), 10)
	return
}
