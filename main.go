package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/janreggie/AdventOfCode/aoc2015"
	"github.com/pkg/profile"
)

var year int     // year (must be between 2015 to 2018)
var day int      // day (must be between 1 to 25)
var input string // filename of input
// to trace or not to trace?
var profileCPU bool
var profileMem bool
var profileTrace bool

const minYear, maxYear = 2015, 2018 // minimum and maximum years so far
const minDay, maxDay = 1, 25        // minimum and maximum days

var allAoCs = [][]func(*bufio.Scanner) (string, string, error){
	aoc2015.AllSolutions}

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
	start := time.Now()
	if result1, result2, err := solverFunc(bufferFile); err != nil {
		glog.Fatalf("some error: %v", err)
	} else {
		fmt.Printf("First answer is %v.\n", result1)
		fmt.Printf("Second answer is %v.\n", result2)
	}
	fmt.Printf("It took %v time to solve Y%04vD%02v.\n", time.Since(start), year, day)
}
