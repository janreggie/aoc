package aoc2015

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/pkg/errors"
)

// racingReindeer is a reindeer that participates in the Reindeer Olympics
type racingReindeer struct {
	name        string
	flyingSpeed uint
	flyingTime  uint
	restingTime uint
}

// reindeerOlympics represents the state of a Reindeer Olympics
// at some particular time
type reindeerOlympics struct {
	competitors []racingReindeer
	points      []uint
	travelled   []uint
	time        uint
}

// newRacingReindeer generates a racing reindeer by reading from a string that looks like:
// NAME can fly FLYINGSPEED km/s for FLYINGTIME seconds, but then must rest for RESTINGTIME seconds.
func newRacingReindeer(description string) (racingReindeer, error) {
	splitDesc := strings.Fields(description)
	if len(splitDesc) != 15 {
		return racingReindeer{}, fmt.Errorf("could not parse %v", description)
	}

	var out racingReindeer
	out.name = splitDesc[0]

	fs, err := strconv.Atoi(splitDesc[3])
	if err != nil {
		return out, errors.Wrapf(err, "could not parse FLYINGSPEED %v", splitDesc[3])
	}
	out.flyingSpeed = uint(fs)

	ft, err := strconv.Atoi(splitDesc[6])
	if err != nil {
		return out, errors.Wrapf(err, "could not parse FLYINGTIME: %v", splitDesc[6])
	}
	out.flyingTime = uint(ft)

	rt, err := strconv.Atoi(splitDesc[13])
	if err != nil {
		return out, errors.Wrapf(err, "could not parse RESTINGTIME: %v", splitDesc[13])
	}
	out.restingTime = uint(rt)

	return out, nil
}

// distance returns the distance that the reindeer would have travelled
// after some amount of time
func (reindeer racingReindeer) distance(time uint) uint {
	roundTime := reindeer.flyingTime + reindeer.restingTime

	if time > roundTime {
		rounds := time / roundTime
		remaining := time % roundTime
		return rounds*reindeer.flyingSpeed*reindeer.flyingTime + reindeer.distance(remaining)
	}

	if time > reindeer.flyingTime {
		return reindeer.flyingSpeed * reindeer.flyingTime
	}

	return reindeer.flyingSpeed * time
}

// isFlying returns if the reindeer is flying at a specific instance in time
func (reindeer racingReindeer) isFlying(time uint) bool {
	if time == 0 {
		return false
	}
	roundTime := reindeer.flyingTime + reindeer.restingTime
	if time > roundTime {
		return reindeer.isFlying(time % roundTime)
	}
	if time > reindeer.flyingTime {
		return false
	}
	return true
}

// newReindeerOlympics creates a reindeer olympics
func newReindeerOlympics(reindeer []racingReindeer) reindeerOlympics {
	var out reindeerOlympics

	out.competitors = make([]racingReindeer, len(reindeer))
	copy(out.competitors, reindeer)
	out.points = make([]uint, len(out.competitors))
	out.travelled = make([]uint, len(out.competitors))
	out.time = 0 // this is already default tho...

	return out
}

// reset resets the Reindeer Olympics
func (olympics *reindeerOlympics) reset() {
	for ii := range olympics.points {
		olympics.points[ii] = 0
		olympics.travelled[ii] = 0
	}
	olympics.time = 0
}

// iterate iterates the Reindeer Olympics by one second
func (olympics *reindeerOlympics) iterate() {
	olympics.time++
	highestDistance := uint(0)
	for ii, reindeer := range olympics.competitors {
		if reindeer.isFlying(olympics.time) {
			olympics.travelled[ii] += reindeer.flyingSpeed
		}
		if olympics.travelled[ii] > highestDistance {
			highestDistance = olympics.travelled[ii]

		}
	}
	for ii := range olympics.competitors {
		if olympics.travelled[ii] == highestDistance {
			olympics.points[ii]++
		}
	}
}

// iterateUntil iterates the Olympics until at some time.
// if olympics.time > time then nothing happens.
func (olympics *reindeerOlympics) iterateUntil(time uint) {
	for olympics.time < time {
		olympics.iterate()
	}
}

// Day14 solves the fourteenth day puzzle "Reindeer Olympics".
//
// Input
//
// A file containing nine lines, each of which
// describes the flying speed, flying time, and resting time for
// a particular reindeer. For example:
//
//	Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
//	Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.
//	Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.
//
// It is guaranteed that the numbers are non-negative integers,
// each no more than 200.
func Day14(input string) (answer1, answer2 string, err error) {

	allReindeer := make([]racingReindeer, 0)
	for _, line := range aoc.SplitLines(input) {

		rd, e := newRacingReindeer(line)
		if e != nil {
			err = errors.Wrapf(err, "could not create reindeer %v", line)
			return
		}
		allReindeer = append(allReindeer, rd)
	}

	olympics := newReindeerOlympics(allReindeer)
	olympics.iterateUntil(2503)
	var furthestDistance, mostPoints uint
	for ii := range olympics.competitors {
		if olympics.points[ii] > mostPoints {
			mostPoints = olympics.points[ii]
		}
		if olympics.travelled[ii] > furthestDistance {
			furthestDistance = olympics.travelled[ii]
		}
	}
	answer1 = strconv.FormatUint(uint64(furthestDistance), 10)
	answer2 = strconv.FormatUint(uint64(mostPoints), 10)

	return
}
