// The AdventOfCode project contains my solutions
// for the Advent of Code series of puzzles from 2015 to 2019.
// Progress varies between these years,
// but I shall finish this project in due time.
//
// Usage
//
// If a year or day is invalid, or if the puzzle for some day of some year
// has yet to be implemented, the program will panic.
//
//	./AdventOfCode [other flags] -year YEAR -day DAY -input FILE
//
// For example:
//
//	./AdventOfCode [other flags] -year 2015 -day 4 -input inputs/2015_04.txt
//
// Profiling flags
//
// There are three profiling flags available: -trace, -cpuprof, and -memprof.
// Only one of them may be turned on at one time, with -trace having leading priority.
// They will write their outputs to trace.out, cpu.pprof, and mem.pprof respectively
// in the current folder.
//
// Glog flags
//
// This package uses the Google logging library https://github.com/golang/glog.
// A solver function may use glog to log what it's doing to Info or Warning.
// The user may set -stderrthreshold to modify the level of verbosity.
//
// Structuring of packages
//
// This project contains several packages: those containing the solutions
// (which start with aoc), and utilities packages (in the internal folder).
//
// The aoc20XX Packages are those that contain the solutions for the Day
// and the Year they are written. For example, the solution for Day 4 of
// Year 2015 is aoc2015.Day04, which is in day04.go in the aoc2015 folder.
// Each of these "Solver Functions" are written such that they take in
// a *bufio.Scanner, which contains the text for the puzzle input,
// and returns two strings answer1 and answer2, and an error.
//
//	// Day04 solves the fourth day puzzle "The Ideal Stocking Stuffer".
//	//
//	// Input
//	//
//	// A string of length 8 containing only the lowercase
//	// letters 'a' to 'z'. For example:
//	//
//	//	iwrupvqb
//	func Day04(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
//		// ...implementation
//	}
//
// For some puzzles, I have written more than one solver function.
// These other solver functions are written as alternative solutions,
// but I believe that the "primary" one is the most optimal.
// For example, there is also an aoc2015.Day04ST,
// which is just like aoc2015.Day04 but is a single-threaded solution.
// Alternative solutions are only there in special cases,
// such as documenting single threaded solutions or using alternative algorithms.
// These are also benchmarked using my input as well as those from other people.
//
// For some puzzles, answer1 or answer2 do not give the exact answer,
// but rather require a human to look at the output directly
// and infer the actual answer from there.
// These puzzles are often in the form of
// "your output will display an ASCII image containing a word which you should input".
// Parsing these outputs are difficult so I have left them be.
//
// I have not written every Solver Function there is yet,
// and for those that have yet to be implemented,
// a dummy Solver Function is written that just throws an error that says "unimplemented".
//
// These Solver Functions could not act alone, and thus would need
// the assistance of helper structs and functions.
// These Helpers are also visible in the .go files in their respective Day and Year,
// and each set of Helpers are unique to a certain Solver Function.
// While these Helpers are left unexported, they can still be viewed in godoc using ?m=all.
// Helper packages also exist in the aoc20XX folders
// wherein multiple puzzles in a year use a set of tools.
//
// Most Helpers and all Solver Functions are tested using the internal package.
// The internal package is only meant to be used for testing.
// In the future, this may be deprecated so that the aoc20XX packages could stand alone.
// These tests also include the inputs that are given to me.
//
// Each aoc20XX package will also have an AllSolutions in their solutions.go,
// which is a way to store all Solver Functions in one slice,
// which this Program imports to choose the correct Solver Function given
// a Day and a Year.
// Note that some of these functions are unimplemented.
//
// The packages written can be used outside of this project,
// by importing the appropriate aoc20XX package.
//
// Input file
//
// The program is designed to read from an input file which is the user's "puzzle input".
// Check the inputs/ folder in this repository for my inputs.
// In addition, I have also placed these inputs as strings in each package
// primarily for testing purposes.
//
// If the puzzle input is not "valid", undefined behavior may follow.
// Often, the Solver function may throw out an error and the program's execution may halt.
// Othertimes, the Solver function may not care and parse the input as-is.
// The tests written mostly do not check for invalid inputs.
// Great care is taken so that none of the solver functions panic.
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
	"github.com/janreggie/aoc/aoc2015"
	"github.com/janreggie/aoc/aoc2016"
	"github.com/janreggie/aoc/aoc2017"
	"github.com/janreggie/aoc/aoc2018"
	"github.com/janreggie/aoc/aoc2019"
	"github.com/janreggie/aoc/aoc2020"
	"github.com/pkg/profile"
)

// allAoCs store all AdventOfCode solutions
var allAoCs = [][]func(*bufio.Scanner) (string, string, error){
	aoc2015.AllSolutions, // 2015
	aoc2016.AllSolutions, // 2016
	aoc2017.AllSolutions, // 2017
	aoc2018.AllSolutions, // 2018
	aoc2019.AllSolutions, // 2019
	aoc2020.AllSolutions, // 2020
}

func main() {
	var year, day int
	var input string
	var profileCPU, profileMem, profileTrace bool
	const minYear, maxYear = 2015, 2020 // minimum and maximum years so far
	const minDay, maxDay = 1, 25        // minimum and maximum days

	// Year, Day, and Input: the essentials
	flag.IntVar(&year, "year", 0, "Year to solve")
	flag.IntVar(&day, "day", 0, "Day to solve")
	flag.StringVar(&input, "input", "", "Input file")

	// to trace, or not to trace?
	flag.BoolVar(&profileCPU, "cpuprof", false, "Profile CPU")
	flag.BoolVar(&profileMem, "memprof", false, "Profile memory")
	flag.BoolVar(&profileTrace, "trace", false, "Enable execution tracing")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\t%s -year YEAR -day DAY -input INPUT_FILE\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Other flags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if year < minYear || year > maxYear {
		flag.Usage()
		glog.Fatalf("invalid year %v", year)
	}
	if day < minDay || day > maxDay {
		flag.Usage()
		glog.Fatalf("invalid day: %v", day)
	}
	if input == "" {
		flag.Usage()
		glog.Fatalf("input file cannot be empty")
	}

	// Enable profiling
	if profileTrace {
		profileCPU = false
		profileMem = false
		defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	} else if profileCPU {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	} else if profileMem {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}

	// open the file
	glog.Infof("To solve: Year %v Day %v\n", year, day)
	glog.Infof("Opening file %v...\n", input)
	osFile, err := os.Open(input)
	defer osFile.Close()
	if err != nil {
		glog.Fatalf("cannot open %v: %+v", input, err)
	}
	bufferFile := bufio.NewScanner(osFile)

	// retrieve solverFunc.
	// Just to be safe, check yearInd, dayInd.
	yearInd, dayInd := year-minYear, day-minDay
	if yearInd >= len(allAoCs) || yearInd < 0 {
		glog.Fatalf("couldn't access allAoCs[%d]; check your program", yearInd)
	}
	if dayInd >= len(allAoCs[yearInd]) || dayInd < 0 {
		glog.Fatalf("couldn't access allAoCs[%d][%d]; check your program", yearInd, dayInd)
	}
	solverFunc := allAoCs[yearInd][day-minDay]
	glog.Infof("Solving puzzle...\n")

	start := time.Now()
	if result1, result2, err := solverFunc(bufferFile); err != nil {
		glog.Fatalf("could not solve Y%04dD%02d: %+v", year, day, err)
	} else {
		fmt.Printf("First answer is %v.\n", result1)
		fmt.Printf("Second answer is %v.\n", result2)
	}
	fmt.Printf("It took %v time to solve Y%04dD%02d.\n", time.Since(start), year, day)
}
