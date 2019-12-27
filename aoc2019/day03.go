package aoc2019

import (
	"bufio"
	"fmt"
	"image"
	"strconv"
	"strings"
)

func maxInt(a int, b ...int) (r int) {
	r = a
	for _, v := range b {
		if v > r {
			r = v
		}
	}
	return
}

type direction int

const (
	up    direction = iota // 00
	right                  // 01
	down                   // 10
	left                   // 11
)

func (dn direction) velocity() (velocity image.Point) {
	switch dn {
	case up:
		velocity.Y = 1
	case right:
		velocity.X = 1
	case down:
		velocity.Y = -1
	case left:
		velocity.X = -1
	}
	return
}

type path struct {
	drn direction // R1006 -> right
	dst int32     // distance R1006 -> 1006
}

func strToPath(raw string) (p path, err error) {
	if len(raw) < 2 {
		err = fmt.Errorf("path %v too short", raw)
		return
	}
	switch b1 := raw[0]; b1 {
	case 'U':
		p.drn = up
	case 'R':
		p.drn = right
	case 'D':
		p.drn = down
	case 'L':
		p.drn = left
	default:
		err = fmt.Errorf("invalid direction for %v: %v", raw, b1)
		return
	}
	d, e := strconv.ParseInt(raw[1:], 10, 16)
	if e != nil {
		err = fmt.Errorf("invalid number for %v: %v", raw, raw[1:])
		return
	}
	p.dst = int32(d)
	return
}

func (p path) move(input image.Point) (output image.Point) {
	// ix, iy becomes ox, oy after being moved through p
	output = input.Add(p.drn.velocity().Mul(int(p.dst)))
	return
}

func determineDimensions(wire []path, otherWires ...[]path) (min, max image.Point) {
	for pointer, ind := image.Pt(0, 0), 0; ind < len(wire); ind++ {
		pointer = wire[ind].move(pointer)
		if pointer.X > max.X {
			max.X = pointer.X
		}
		if pointer.Y > max.Y {
			max.Y = pointer.Y
		}
		if pointer.X < min.X {
			min.X = pointer.X
		}
		if pointer.Y < min.Y {
			min.Y = pointer.Y
		}
	}
	// consider then otherWires
	for _, eachWire := range otherWires {
		for pointer, ind := image.Pt(0, 0), 0; ind < len(eachWire); ind++ {
			pointer = eachWire[ind].move(pointer)
			if pointer.X > max.X {
				max.X = pointer.X
			}
			if pointer.Y > max.Y {
				max.Y = pointer.Y
			}
			if pointer.X < min.X {
				min.X = pointer.X
			}
			if pointer.Y < min.Y {
				min.Y = pointer.Y
			}
		}
	}
	// add 1,1 to max for good measure
	max = max.Add(image.Pt(1, 1))
	return
}

// tracePath will trace wire on the trace (will modify!!)
// as well as count how many steps it takes to reach critPoints
func tracePath(wire []path, trace [][]bool, origin image.Point, critPoints map[image.Point]int32) {
	// With regards to tracing,
	// minPoint (declared in Day03) is the bottom leftmost point of the grid
	// which can either be the origin or something else with negative coordinates.
	// When performing operations on trace1 and trace2,
	// tracing has to take note of this minimum point,
	// which can be used as our *offset*.
	// What I meant by this is that,
	// if it is evaluated that the trace reaches points (-6,-9) and (-8, -7),
	// minPoint shall be (-8, -9).
	// Noting that we can't access trace[-6][-9],
	// we can use the offset as how much needs to beadded to the Point
	// such that, if minPoint = (-8, -9),
	// our "origin" can be represented as (8, 9)
	// offsetting pt by origin.
	var counter int32

	tracePoint := func(pt image.Point) {
		pt = pt.Add(origin)
		trace[pt.X][pt.Y] = true
		if _, ok := critPoints[pt]; ok {
			critPoints[pt] += counter // += since will be used by trace1 and trace2
		}
	}
	pointer := image.Pt(0, 0)
	for _, currentPath := range wire {
		// trace everything from pointer to currentPath.move(pointer)
		for step := int32(0); step < currentPath.dst; step++ {
			pointer = pointer.Add(currentPath.drn.velocity())
			counter++
			tracePoint(pointer)
		}
	}
}

// checkIntersect checks all cells within distance distance from origin
// and returns true if it is common to both trace1 and trace2
// as well as all intersections that it can find
func checkIntersect(trace1, trace2 [][]bool, distance int, origin, maxPoint image.Point) (has bool, pts []image.Point) {
	checkPoint := func(pt image.Point) bool {
		return trace1[pt.X][pt.Y] == true && trace2[pt.X][pt.Y] == true
	}
	bounds := image.Rectangle{image.Pt(0, 0), maxPoint.Add(origin)}
	if distance > maxPoint.X+maxPoint.Y && distance > origin.X+origin.Y {
		// there's no point in continuing really
		return
	}
	// first quadrant: check from +x-axis to +y-axis
	pointer := image.Pt(origin.X+distance, origin.Y)
	if pointer.X >= origin.X+maxPoint.X {
		pointer.Y += pointer.X - (origin.X + maxPoint.X - 1)
		pointer.X = maxPoint.X - 1
	}
	for ; pointer.In(bounds) && pointer.X > origin.X; pointer = pointer.Add(image.Pt(-1, 1)) {
		if checkPoint(pointer) {
			pts = append(pts, pointer)
			has = true
		}
	}
	// second quadrant: from +y-axis to -x-axis
	pointer = image.Pt(origin.X, origin.Y+distance)
	if pointer.Y >= origin.Y+maxPoint.Y {
		pointer.X -= pointer.Y - (origin.Y + maxPoint.Y - 1)
		pointer.Y = maxPoint.Y - 1
	}
	for ; pointer.In(bounds) && pointer.Y > origin.Y; pointer = pointer.Add(image.Pt(-1, -1)) {
		if checkPoint(pointer) {
			pts = append(pts, pointer)
			has = true
		}
	}
	// third quadrant: from -x-axis to -y-axis
	pointer = image.Pt(origin.X-distance, origin.Y)
	if pointer.X < 0 {
		pointer.Y -= 0 - pointer.X // pointer.Y will be subtracted by abs(pointer.X)
		pointer.X = 0
	}
	for ; pointer.In(bounds) && pointer.X < origin.X; pointer = pointer.Add(image.Pt(1, -1)) {
		if checkPoint(pointer) {
			pts = append(pts, pointer)
			has = true
		}
	}
	// fourth quadrant: from -y-axis to +x-axis
	pointer = image.Pt(origin.X, origin.Y-distance)
	if pointer.Y < 0 {
		pointer.X += 0 - pointer.Y
		pointer.Y = 0
	}
	for ; pointer.In(bounds) && pointer.Y < origin.Y; pointer = pointer.Add(image.Pt(1, 1)) {
		if checkPoint(pointer) {
			pts = append(pts, pointer)
			has = true
		}
	}
	return
}

// Day03 solves the third day puzzle
// "Crossed Wires"
//
// WARNING: This may consume half a gigabyte of RAM.
// But it works. And more often than not that's what matters.
//
// TODO (janreggie): Optimize.
func Day03(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	wire1 := make([]path, 0) // we don't know how long it is just yet
	wire2 := make([]path, 0)
	var trace1, trace2 [][]bool           // "traces" for the wire; 0 if non-traced; 1 if traced
	var minPoint, maxPoint image.Point    // maximum and minimum ranges for both trace1 and trace2
	crits1 := make(map[image.Point]int32) // critical points for later
	crits2 := make(map[image.Point]int32)
	var origin image.Point

	if scan := scanner.Scan(); !scan {
		err = fmt.Errorf("premature scan error")
		return
	}
	for _, eachPath := range strings.Split(scanner.Text(), ",") {
		p, e := strToPath(eachPath)
		if e != nil {
			err = fmt.Errorf("cannot extract path: %v", e)
			return
		}
		wire1 = append(wire1, p)
	}

	if scan := scanner.Scan(); !scan {
		err = fmt.Errorf("premature scan error")
		return
	}
	for _, eachPath := range strings.Split(scanner.Text(), ",") {
		p, e := strToPath(eachPath)
		if e != nil {
			err = fmt.Errorf("cannot extract path: %v", e)
			return
		}
		wire2 = append(wire2, p)
	}

	// now what are the maximum dimensions of the paths for wire1 and wire2?
	minPoint, maxPoint = determineDimensions(wire1, wire2)
	trace1 = make([][]bool, maxPoint.X-minPoint.X)
	for ind := range trace1 {
		trace1[ind] = make([]bool, maxPoint.Y-minPoint.Y)
	}
	trace2 = make([][]bool, maxPoint.X-minPoint.X)
	for ind := range trace2 {
		trace2[ind] = make([]bool, maxPoint.Y-minPoint.Y)
	}
	origin = minPoint.Mul(-1)
	// glog.Info("minPoint: ", minPoint)
	// glog.Info("maxPoint: ", maxPoint)

	// from here on do the actual tracing
	// right now allIntersections contains nothing
	tracePath(wire1, trace1, origin, crits1)
	tracePath(wire2, trace2, origin, crits2) // it points to nowhere right now
	// // print the traces out?
	// boolsToStr := func(bools []bool) string {
	// 	s := make([]string, len(bools))
	// 	for i, v := range bools {
	// 		if v {
	// 			s[i] = "X"
	// 		} else {
	// 			s[i] = "."
	// 		}
	// 	}
	// 	return strings.Join(s, "")
	// }
	// glog.Info("trace 1:")
	// for i, v := range trace1 {
	// 	glog.Info(i, ":", boolsToStr(v))
	// }
	// glog.Info("trace 2:")
	// for i, v := range trace2 {
	// 	glog.Info(i, ":", boolsToStr(v))
	// }

	// now that trace1 and trace2 have been established
	// let us determine if there are any points that are common
	// between trace1 and trace2
	// ii steps from the origin
	for ii := 0; ii < maxInt(maxPoint.X+maxPoint.Y+2, origin.X+origin.Y); ii++ {
		has, pts := checkIntersect(trace1, trace2, ii, origin, maxPoint)
		if has && answer1 == "" {
			answer1 = strconv.Itoa(ii)
		}
		// glog.Infof("%v found some points: %v", ii, pts)
		for _, v := range pts {

			crits1[v] = 0 // initializing!
			crits2[v] = 0 // initializing too!
		}
	}

	// part 2
	for ii := range trace1 {
		for jj := range trace1[ii] {
			trace1[ii][jj] = false
			trace2[ii][jj] = false // they have the same bounds!
		}
	}
	tracePath(wire1, trace1, origin, crits1)
	tracePath(wire2, trace2, origin, crits2) // should result to something
	// glog.Info("crit points: ", crits1)       // they're the same either way
	for k, v := range crits2 {
		crits1[k] += v
	}
	// now determine which is the lowest value
	lowest := int32(0)
	for _, v := range crits1 {
		if lowest == 0 {
			lowest = v
		}
		if v < lowest {
			lowest = v
		}
	}
	answer2 = strconv.Itoa(int(lowest))

	return
}
