package aoc2015

import (
	"bufio"
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// lights represents a 1000x1000 grid of lights (Day 6)
type lights [1000][1000]struct {
	status     bool // on/off?
	brightness uint
}

// turnOn turns on all points in a rectangular region from initial to final inclusive
// and increases their brightness
func (lights *lights) turnOn(initial, final image.Point) {
	for ii := initial.X; ii <= final.X; ii++ {
		for jj := initial.Y; jj <= final.Y; jj++ {
			lights[jj][ii].status = true
			lights[jj][ii].brightness++
		}
	}
}

// turnOff turns off all points in a rectangular region from initial to final inclusive
// and reduces their brightness
func (lights *lights) turnOff(initial, final image.Point) {
	for ii := initial.X; ii <= final.X; ii++ {
		for jj := initial.Y; jj <= final.Y; jj++ {
			lights[jj][ii].status = false
			if lights[jj][ii].brightness > 0 {
				lights[jj][ii].brightness--
			}
		}
	}
}

// toggle toggles all points in a rectangular region from initial to final inclusive
// and increases their brightness by 2
func (lights *lights) toggle(initial, final image.Point) {
	for ii := initial.X; ii <= final.X; ii++ {
		for jj := initial.Y; jj <= final.Y; jj++ {
			lights[jj][ii].status = !lights[jj][ii].status
			lights[jj][ii].brightness += 2
		}
	}
}

// countLit counts the number of lit lights
func (lights *lights) countLit() uint {
	var totalLit uint
	for ii := 0; ii < 1000; ii++ {
		for jj := 0; jj < 1000; jj++ {
			if lights[jj][ii].status {
				totalLit++
			}
		}
	}
	return totalLit
}

// brightness returns the brightness
func (lights *lights) brightness() uint {
	var br uint
	for ii := 0; ii < 1000; ii++ {
		for jj := 0; jj < 1000; jj++ {
			br += lights[jj][ii].brightness
		}
	}
	return br
}

// Day06 solves the sixth day puzzle "Probably a Fire Hazard".
//
// Input
//
// A file containing 300 lines, each of which is an "instruction"
// that determines which lights to turn off, turn on, or toggle on a 1000x1000 grid.
// For example:
// 	turn on 818,296 through 818,681
// 	turn on 71,699 through 91,960
// 	turn off 838,578 through 967,928
// 	toggle 440,856 through 507,942
//
// Each instruction is guaranteed to be valid where each instructions is
// prefixed by "turn on", "turn off", or "toggle",
// the first coordinate is no greater than the third one and the second no greater than the fourth,
// and all numbers are no greater than 999.
func Day06(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	var allLights lights

	// parseRowCol turns dimensions initRaw and finalRaw
	// into two image points
	parseRowCol := func(initRaw, finalRaw string, initial, final *image.Point) error {
		// dimensions should be of form "ROW,COL"
		dims := strings.Split(initRaw, ",")
		if len(dims) != 2 {
			return fmt.Errorf("%q not comma separated", dims)
		}
		var err error
		if initial.X, err = strconv.Atoi(dims[0]); err != nil {
			return errors.Wrapf(err, "unknown row %q in query %q", dims[0], initRaw)
		}
		if initial.Y, err = strconv.Atoi(dims[1]); err != nil {
			return errors.Wrapf(err, "unknown col %q in query %q", dims[1], initRaw)
		}
		dims = strings.Split(finalRaw, ",")
		if len(dims) != 2 {
			return fmt.Errorf("%q not comma separated", dims)
		}
		if final.X, err = strconv.Atoi(dims[0]); err != nil {
			return errors.Wrapf(err, "unknown row %q in query %q", dims[0], initRaw)
		}
		if final.Y, err = strconv.Atoi(dims[1]); err != nil {
			return errors.Wrapf(err, "unknown col %q in query %q", dims[1], initRaw)
		}
		return nil
	}

	// read each line in scanner
	for scanner.Scan() {
		rawQuery := scanner.Text()
		query := strings.Fields(rawQuery)
		// note that rawQuery can either be one of the following
		// 	toggle ROW1,COL1 through ROW2,COL2
		// 	turn off ROW1,COL1 through ROW2,COL2
		// 	turn on ROW1,COL1 through ROW2,COL2
		// hence query is of length 4 or 5
		if len(query) != 4 && len(query) != 5 {
			err = fmt.Errorf("unknown spacing: %q", rawQuery)
			return
		}
		var initial, final image.Point
		if query[0] == "toggle" {
			if err = parseRowCol(query[1], query[3], &initial, &final); err != nil {
				err = errors.Wrapf(err, "could not parse query %v", rawQuery)
				return
			}
			allLights.toggle(initial, final)
		} else if query[0] == "turn" {
			if err = parseRowCol(query[2], query[4], &initial, &final); err != nil {
				err = errors.Wrapf(err, "could not parse query %v", rawQuery)
				return
			}
			if query[1] == "off" {
				allLights.turnOff(initial, final)
			} else if query[1] == "on" {
				allLights.turnOn(initial, final)
			} else {
				err = fmt.Errorf("unknown function: %q in %q neither off nor on", query[1], rawQuery)
				return
			}
		} else {
			err = fmt.Errorf("unknown function: %q in %q", query[0], rawQuery)
			return
		}
	}

	// now count
	answer1 = strconv.FormatUint(uint64(allLights.countLit()), 10)
	answer2 = strconv.FormatUint(uint64(allLights.brightness()), 10)
	return
}
