package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
var allAoCs = [][]func(input string) (string, string, error){
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
	fileContents, err := ioutil.ReadFile(input)
	if err != nil {
		glog.Fatalf("cannot open %v: %+v", input, err)
	}
	if fileContents[len(fileContents)-1] == '\n' {
		fileContents = fileContents[:len(fileContents)-1] // Remove trailing newline
	}

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
	if result1, result2, err := solverFunc(string(fileContents)); err != nil {
		glog.Fatalf("could not solve Y%04dD%02d: %+v", year, day, err)
	} else {
		fmt.Printf("First answer is %v.\n", result1)
		fmt.Printf("Second answer is %v.\n", result2)
	}
	fmt.Printf("It took %v time to solve Y%04dD%02d.\n", time.Since(start), year, day)
}
