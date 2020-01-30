// The AdventOfCode project contains my solutions for the Advent of Code series of puzzles from 2015 to 2019.
// Progress varies between these years, but I shall attempt to finish this
// project in due time.
//
// Usage
//
// If a year or day is invalid, or if the puzzle for some day of some year
// has yet to be implemented, the program will panic.
//
//	./AdventOfCode [profiling flags] -year YEAR -day DAY -input FILE
//
// For example:
//
//	./AdventOfCode [profiling flags] -year 2015 -day 4 -input inputs/2015_04.txt
//
// Profiling flags
//
// There are three profiling flags available: -trace, -cpuprof, and -memprof.
// Only one of them may be turned on at one time, with -trace having leading priority.
// They will write their outputs to trace.out, cpu.pprof, and mem.pprof respectively
// in the current folder.
//
// Structuring of packages
//
// This project contains several packages: those containing the solutions
// (which start with aoc), and the utilities packages (those in the tools folder).
//
// The Solutions Packages are those that contain the solutions for the Day
// and the Year they are written. For example, the solution for Day 4 of
// Year 2015 is aoc2015.Day04, which is in day04.go in the aoc2015 folder.
// Each of these "Solver Functions" are written such that they take in
// a *bufio.Scanner, which contains the text for the puzzle input,
// and returns two strings answer1 and answer2, and an error.
//
// I have not written every Solver Function there is yet,
// and for those that have yet to be implemented,
// a dummy Solver Function is written that just throws an error
// that says "unimplemented".
//
// These Solver Functions could not act alone, and thus would need
// the assistance of helper structs and functions.
// These Helpers are also visible in the .go files in their respective Day and Year,
// and each set of Helpers are unique to a certain Solver Function.
// While these Helpers are left unexported, they can still be viewed in godoc using ?m=all.
//
// Most Helpers and all Solver Functions are tested using the utilities
// packages (seen below). These tests also include my inputs,
//
// Each Solutions Package will also have an AllSolutions in their init.go,
// which is a way to store all Solver Functions in one slice,
// which this Program imports to choose the correct Solver Function given
// a Day and a Year.
//
// The utilities packages are there to facilitate the operation of several
// Solver Functions, such as the Intcode package used for several days in
// aoc2019. A utility package is also written to facilitate my testing.
//
// The packages written can be used outside of this project,
// by, for instance, importing the appropriate Solutions Package.
//
// Input file
//
// The program is designed to read from an input file which is the user's "puzzle input".
// Check the inputs/ folder in this repository for example inputs of the author.
// If the puzzle input is not "valid", undefined behavior may follow.
// The Solver function may throw out an error and the program's execution may halt.
//
// Each Solver Function should provide documentation on what the input file looks like.
//
// License
//
// MIT License:
//
//	Copyright (c) 2019 janreggie
//
//	Permission is hereby granted, free of charge, to any person obtaining a copy
//	of this software and associated documentation files (the "Software"), to deal
//	in the Software without restriction, including without limitation the rights
//	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//	copies of the Software, and to permit persons to whom the Software is
//	furnished to do so, subject to the following conditions:
//
//	The above copyright notice and this permission notice shall be included in all
//	copies or substantial portions of the Software.
//
//	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//	SOFTWARE.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/janreggie/AdventOfCode/aoc2016"
	"github.com/janreggie/AdventOfCode/aoc2017"
	"github.com/janreggie/AdventOfCode/aoc2018"
	"github.com/janreggie/AdventOfCode/aoc2019"
	"github.com/pkg/profile"
)

var year int     // year (must be between 2015 to 2018)
var day int      // day (must be between 1 to 25)
var input string // filename of input
// to trace or not to trace?
var profileCPU bool
var profileMem bool
var profileTrace bool

const minYear, maxYear = 2015, 2019 // minimum and maximum years so far
const minDay, maxDay = 1, 25        // minimum and maximum days

var allAoCs = [][]func(*bufio.Scanner) (string, string, error){
	aoc2015.AllSolutions, // 2015
	aoc2016.AllSolutions, // 2016
	aoc2017.AllSolutions, // 2017
	aoc2018.AllSolutions, // 2018
	aoc2019.AllSolutions, // 2019
}

func init() {
	flag.IntVar(&year, "year", 0, "Year to solve")
	flag.IntVar(&day, "day", 0, "Day to solve")
	flag.StringVar(&input, "input", "", "Input file")
	flag.BoolVar(&profileCPU, "cpuprof", false, "Profile CPU")
	flag.BoolVar(&profileMem, "memprof", false, "Profile memory")
	flag.BoolVar(&profileTrace, "trace", false, "Enable execution tracing")
	flag.Parse()

	if year < minYear || year > maxYear {
		glog.Fatalf("invalid year %v", year)
	}
	if day < minDay || day > maxDay {
		glog.Fatalf("invalid day: %v", day)
	}
	if input == "" {
		glog.Fatalf("input file cannot be empty")
	}
	if profileTrace {
		profileCPU = false
		profileMem = false
	}
}

func main() {
	// Enable profiling
	if profileTrace {
		defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	} else if profileCPU {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	} else if profileMem {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}

	// implement flags
	glog.Infof("To solve: Year %v Day %v\n", year, day)
	glog.Infof("Opening file %v...\n", input)
	osFile, err := os.Open(input)
	defer osFile.Close()
	if err != nil {
		glog.Fatalf("cannot open %v: %v", input, err)
	}
	bufferFile := bufio.NewScanner(osFile)

	// now this is where it gets interesting
	solverFunc := allAoCs[year-minYear][day-minDay]
	glog.Infof("Solving puzzle...\n")
	start := time.Now()
	if result1, result2, err := solverFunc(bufferFile); err != nil {
		glog.Fatalf("some error: %v", err)
	} else {
		fmt.Printf("First answer is %v.\n", result1)
		fmt.Printf("Second answer is %v.\n", result2)
	}
	fmt.Printf("It took %v time to solve Y%04vD%02v.\n", time.Since(start), year, day)
}
