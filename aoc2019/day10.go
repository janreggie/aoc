package aoc2019

import (
	"bufio"
	"fmt"
	"image"
	"math/big"
	"strconv"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

// gcd and lcm from https://play.golang.org/p/SmzvkDjYlb
// author unknown.

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	// make sure a > b
	if b > a {
		return gcd(b, a)
	}
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

// lcmUpto generates the lcm of 1, 2, up to limit.
// If limit is less than 2, it will return limit
func lcmUpto(limit int) int {
	if limit <= 1 {
		return 1
	}
	slice := make([]int, 0, limit)
	for ii := 1; ii < limit; ii++ {
		slice = append(slice, ii)
	}
	return lcm(1, limit, slice...)
}

// findNextCoprime finds the next image.Point
// where its coordinates are coprime
// and where pt.X > pt.Y.
// If pt.X < 2 or pt.Y < 1, pt becomes (2,1).
// If pt.X <= pt.Y, pt becomes (pt.X+1, 1).
// For example, putting in (6,1) returns immediately
// but putting in (6,2) will cause it to iterate
// to (6,2), (6,3), until it reaches (6,5).
// Putting in (6,5) will cause it to iterate to (6,6),
// which is invalid, thus turning it into (7,1).
func findNextCoprime(pt *image.Point) {
	// make sure a is at least 2
	if pt.X < 2 || pt.Y < 1 {
		*pt = image.Point{2, 1}
		return
	}
	// make sure a > b
	if pt.X <= pt.Y {
		*pt = image.Point{pt.X + 1, 1}
		return
	}
	// now for the iteration
	pt.Y = pt.Y + 1
	for !isCoprime(*pt) {
		pt.Y = pt.Y + 1
	}
	if pt.Y >= pt.X {
		pt.X++
		pt.Y = 1
	}
}

// isCoprime checks the coordinates of image.Point
// and returns true if they are coprime.
// If pt.X or pt.Y is nonpositive,
// or if they are equal to each other, return false
func isCoprime(pt image.Point) bool {
	x := pt.X
	y := pt.Y
	if x <= 0 || y <= 0 { // negative?!
		return false
	}
	if x == y { // no way.
		return false
	}
	return gcd(x, y) == 1
}

// generateCoprime generates a channel of some size
// that returns coprime points.
// The pattern is as follows:
// (2,1), (3,1), (3,2),
// (4,1), (4,3), (5,1), (5,2), and so on.
// Note that (1,0) is included, and (1,1) is skipped.
func generateCoprime(size int) chan image.Point {
	if size < 1 {
		size = 1
	}
	channel := make(chan image.Point, size)
	tracker := image.Pt(2, 1)
	channel <- tracker
	go func() {
		for {
			findNextCoprime(&tracker)
			channel <- tracker // add in perpetua.
		}
	}()
	return channel
}

// asteroidField represents a field of asteroids
type asteroidField struct {
	asteroids     [][]bool
	width, height int
}

// newAsteroidField parses a bufio.Scanner
// and returns an asteroidField object.
//
// The scanned text must look like the following:
// .#..##
// ......
// #####.
// ....#.
// ...##.
// with a fixed width (6) and a height (5).
// In case its width and height are not fixed,
// or if its width or height appear to be 0,
// it will return an error.
// Make sure that the string only contains `.` or `#`
// because any character other than `#` will result to false
// or non-presence of asteroid.
func newAsteroidField(scanner *bufio.Scanner) (*asteroidField, error) {
	// okay... how do we store this
	asteroids := make([][]bool, 0) // true: asteroid present; false: no asteroid
	// asteroids[2][3] or asteroids(3,2) represents the 3rd asteroid in the 2nd row (counting from row 0, row 1,...)
	var fieldWidth, fieldHeight int // width and height as observed.
	// we know the input is 26x26 though... but we want to support other sizes.
	for scanner.Scan() {
		text := scanner.Text()
		if fieldWidth == 0 {
			fieldWidth = len(text)
		}
		if fieldWidth != len(text) {
			return nil, fmt.Errorf("width of %v should b %v", text, fieldWidth)
		}
		asteroids = append(asteroids, make([]bool, fieldWidth))
		for ii := range asteroids[fieldHeight] {
			if text[ii] == '#' {
				// default is false.
				asteroids[fieldHeight][ii] = true
			}
		}
		fieldHeight++
	}
	// check if asteroids is valid
	if fieldHeight == 0 {
		return nil, errors.New("height is zero")
	}
	if fieldWidth == 0 {
		return nil, errors.New("width is zero")
	}
	return &asteroidField{asteroids: asteroids, width: fieldHeight, height: fieldHeight}, nil
}

// at returns true if there is an asteroid at a position point(x,y),
// false if there is none or if point is out of bounds
func (field *asteroidField) at(point image.Point) bool {
	if point.X >= field.width || point.X < 0 {
		return false
	}
	if point.Y >= field.height || point.Y < 0 {
		return false
	}
	return field.asteroids[point.Y][point.X]
}

// bounds returns an image.Rect representing the bounds of the asteroid field
func (field *asteroidField) bounds() image.Rectangle {
	return image.Rect(0, 0, field.width, field.height)
}

// countLineOfSight counts the number of asteroids a position index sees in its line of sight.
func (field *asteroidField) countLineOfSight(index image.Point) int {
	count := 0
	glog.Infof("counting surrounding %v", index)
	for range field.sweepClockwise(index) {
		count++
	}
	return count
	// // Centering on index,
	// // suppose let us divide the 2-dimensional grid into eight parts,
	// // where each part is between one of the major axes
	// // and one of the lines that pass thru the origin and with slope -1 or 1.
	// responses := make(chan int, 8)                        // responses on these, um... parts
	// border := image.Rect(0, 0, field.width, field.height) // the edges of the play grid

	// // There are eight possible points from index which can be checked
	// // where point is a generated coprime point:
	// // (index.X+point.X, index.Y+point.Y)
	// // (index.X+point.X, index.Y-point.Y)
	// // (index.X+point.Y, index.Y+point.X)
	// // (index.X+point.Y, index.Y-point.X)
	// // (index.X-point.X, index.Y+point.Y)
	// // (index.X-point.X, index.Y-point.Y)
	// // (index.X-point.Y, index.Y+point.X)
	// // (index.X-point.Y, index.Y-point.X)
	// // Now how do we generate these
	// for ii := 0; ii < 8; ii++ {
	// 	go func(flags int) {
	// 		addFxn := func(pt image.Point) image.Point {
	// 			// these three booleans should be able to represent all configurations.
	// 			transpose := flags%2 == 0 // transpose
	// 			// true: index.X adds/subtracts to point.X and index.Y adds/subtracts to point.Y
	// 			// false: index.Y adds/subtracts to point.X and index.X adds/subtracts to point.Y
	// 			invert := (flags/2)%2 == 0 // invert
	// 			// true: index.X gets added and index.Y gets subtracted or vice versa
	// 			// false: index.X and index.Y both get added or subtracted
	// 			subtract := (flags/4)%2 == 0 // subtract
	// 			// true: index.X gets subtracted
	// 			// false: index.X gets added
	// 			if transpose {
	// 				pt.X, pt.Y = pt.Y, pt.X
	// 			}
	// 			if invert {
	// 				pt.Y = -pt.Y
	// 			}
	// 			if subtract {
	// 				pt = pt.Mul(-1)
	// 			}
	// 			return pt
	// 		}

	// 		coprimeChannel := generateCoprime(1)
	// 		countFound := 0 // how many asteroids have we spotted using coprimeChannel
	// 		for {
	// 			affix := <-coprimeChannel // the next point to be examined
	// 			point := index.Add(addFxn(affix))
	// 			// how do we know if point exceeds the half-quadrant?
	// 			if !point.In(border) {
	// 				if affix.Y == 1 { // a fresh depth
	// 					break // we know we're done.
	// 				}
	// 				continue // no point in continuing
	// 			}
	// 			for point.In(border) {
	// 				// guaranteed field.at will not return some error
	// 				if field.at(point) {
	// 					countFound++
	// 					break
	// 				}
	// 				point = point.Add(addFxn(affix))
	// 			}

	// 		}
	// 		responses <- countFound
	// 	}(ii)
	// }

	// // but we haven't determined the position diagonal to index
	// // as well as that of the direct sides

	// // check its adjacent sides
	// for _, affix := range []image.Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
	// 	point := index.Add(affix)
	// 	for point.In(border) {
	// 		// guaranteed field.at will not return some error
	// 		if field.at(point) {
	// 			count++
	// 			break
	// 		}
	// 		point = point.Add(affix)
	// 	}
	// }

	// // check its diagonals
	// for _, affix := range []image.Point{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}} {
	// 	point := index.Add(affix)
	// 	for point.In(border) {
	// 		// guaranteed field.at will not return some error
	// 		if field.at(point) {
	// 			count++
	// 			break
	// 		}
	// 		point = point.Add(affix)
	// 	}
	// }

	// for ii := 0; ii < 8; ii++ {
	// 	count += <-responses // will block until one has been sent...
	// }
	// return count
}

// rationalToPoint returns an image.Point representation of a rational number.
// Note that it truncates the number to the capacity of int
// (which may vary from system to system)
func rationalToPoint(rational *big.Rat) image.Point {
	return image.Pt(int(rational.Num().Int64()), int(rational.Denom().Int64()))
}

// Consider some point P(X,Y) on a lattice grid of asteroids
// where X denotes the distance from the left edge
// and Y denotes the distance from the bottom edge.
// Let a Sison point be a point Q_{i,j} (x,y) in the lattice grid
// such that the segment connecting (i,j) and (x,y) does not intersect any other points
// in such a lattice grid.
// For instance, the Sison points of (0,0) of a 3x3 grid
// are (0,1), (1,0), (1,1), (2,1), and (1,2).

// generateSisonPoints creates a channel of size 4
// that generates Sison points for the upper-right quadrant of the Cartesian coordinate system
// such that it starts from (0, 0), (1, height), ..., (1, 1), ..., (width, 1)
// and then it closes it.
// If both width and height are non-positive
// the function returns an already closed channel.
// If only the height is positive, return (0,1).
//
// TODO: Optimize. Current algorithm brute forces by iterating thru
// lcmUpto(width) and lcmUpto(height).
// Average case is exponential time!
// (https://math.stackexchange.com/questions/1111334/reason-for-lcm-of-all-numbers-from-1-n-equals-roughly-en)
func generateSisonPoints(width, height int) chan image.Point {
	channel := make(chan image.Point, 4)

	go func() {
		// okay what if width and height are -1
		if width <= 0 && height <= 0 {
			close(channel)
			return
		}

		// if height is positive?
		if width <= 0 {
			channel <- image.Pt(0, 1)
			close(channel)
			return
		}

		denom := lcmUpto(height)
		pointer := big.NewRat(0, int64(denom))
		for pointer.Cmp(big.NewRat(1, 1)) < 0 { // while less than 1
			// note that two checks are performed here,
			// that is, if the denominator is no more than the height
			// and the numerator is no more than the width.
			// The checks are performed to ensure that the points added
			// are in the bounds of the height and width.
			if pointer.Denom().Int64() <= int64(height) &&
				pointer.Num().Int64() <= int64(width) {
				channel <- rationalToPoint(pointer)
			}
			pointer.Add(pointer, big.NewRat(1, int64(denom)))
		}
		// okay now let's go from 1,1 to 0,0
		denom = lcmUpto(width)
		for pointer.Cmp(big.NewRat(0, 1)) > 0 { // while greater than 0
			if pointer.Denom().Int64() <= int64(width) &&
				pointer.Num().Int64() <= int64(height) {
				// the coordinates of the point have to be switched
				// so that they would not be redundant
				// to the previous for loop
				pt := rationalToPoint(pointer)
				pt.X, pt.Y = pt.Y, pt.X
				channel <- pt
			}
			pointer.Add(pointer, big.NewRat(-1, int64(denom)))
		}
		// ok we're done
		close(channel)
	}()

	return channel
}

// sweepClockwise creates a channel of size 4
// that gets fed with coordinates of asteroids
// that are also Sison points of a relativeTo point
// in a clockwise direction
// starting from the point's north (relativeTo(x,y-1))
// and closes after one sweep.
func (field *asteroidField) sweepClockwise(relativeTo image.Point) chan image.Point {
	if !relativeTo.In(field.bounds()) {
		return nil // there's no point (ehehe...)
	}
	channel := make(chan image.Point, 4) // 4 elements is good enough

	// now how do we generate Sison points...
	go func() {
		func() { // generate for upper-right quadrant
			quarterWidth := field.width - relativeTo.X - 1
			quarterHeight := relativeTo.Y
			glog.Infof("For %v quadrant, width: %v, height:%v", "upper-right", quarterWidth, quarterHeight)
			for sisonPoint := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(sisonPoint.X, -sisonPoint.Y)
				glog.Infof("generated addend %v", addend)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					glog.Infof("absolute position is %v", absolute)
					if field.at(absolute) {
						glog.Infof("%v saw %v at upper right\n", relativeTo, absolute)
						channel <- absolute
						break
					}
				}
				glog.Infof("Hey thanks!")
			}
		}()

		func() { // generate for lower-right quadrant
			// As we are sweeping counterclockwise,
			// the height of the Sison Point Generator would be the horizontal distance based on field.width,...,
			// that is, field.width - relativeTo.X - 1,
			// and the width of the Generator would be the vertical distance based on field.height.
			quarterWidth := field.height - relativeTo.Y - 1
			quarterHeight := field.width - relativeTo.X - 1
			glog.Infof("For %v quadrant, width: %v, height:%v", "lower-right", quarterWidth, quarterHeight)
			for sisonPont := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(+sisonPont.Y, +sisonPont.X)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					if field.at(absolute) {
						glog.Infof("%v saw %v at lower right\n", relativeTo, absolute)
						channel <- absolute
						break
					}
				}
			}
		}()

		func() { // generate for the lower-left quadrant
			quarterWidth := relativeTo.X
			quarterHeight := field.height - relativeTo.Y - 1
			glog.Infof("For %v quadrant, width: %v, height:%v", "lower-left", quarterWidth, quarterHeight)
			for sisonPoint := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(-sisonPoint.X, +sisonPoint.Y)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					if field.at(absolute) {
						glog.Infof("%v saw %v at lower left\n", relativeTo, absolute)
						channel <- absolute
						break
					}
				}
			}
		}()

		func() { // generate for the upper-left quadrant
			// Similar to the lower-right quadrant,
			// the quarter width and height appear flipped
			quarterWidth := relativeTo.Y
			quarterHeight := relativeTo.X
			glog.Infof("For %v quadrant, width: %v, height:%v", "upper-left", quarterWidth, quarterHeight)
			for sisonPoint := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(-sisonPoint.Y, -sisonPoint.X)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					if field.at(absolute) {
						glog.Infof("%v saw %v at upper left\n", relativeTo, absolute)
						channel <- absolute
						break
					}
				}
			}
		}()

		close(channel) // oh right... close the channel.
	}()

	return channel
}

// sweepAsteroids vaporizes some number of asteroids from the north
// and returns the coordinates of the last vaporized asteroid
func (field *asteroidField) sweepAsteroids(station image.Point, toVaporize int) image.Point {
	// how do we do this...
	// vaporizedCount := 0 // number that has been vaporized

	return image.Point{}
}

// Day10 solves the tenth day problem
// "Monitoring Station"
func Day10(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	asteroids, err := newAsteroidField(scanner)
	if err != nil {
		err = errors.Wrap(err, "Y2019D10: could not create asteroidField")
		return
	}

	mostSeen := 0
	bestAsteroid := image.Pt(0, 0)
	for ii := 0; ii < asteroids.width; ii++ {
		for jj := 0; jj < asteroids.height; jj++ {
			if !asteroids.at(image.Pt(ii, jj)) {
				continue // don't care about non-asteroids
			}
			// for range asteroids.sweepClockwise(image.Pt(ii, jj)) {
			// 	/// do nothing...
			// }
			if seen := asteroids.countLineOfSight(image.Pt(ii, jj)); seen > mostSeen {
				glog.Infof("asteroid %v sees %v\n", image.Pt(ii, jj), seen)
				mostSeen = seen
				bestAsteroid = image.Pt(ii, jj)
			}
		}
	}
	answer1 = strconv.Itoa(mostSeen)

	// now for Part 2
	// Okay this is epic.
	// Apparently you can return the points going clockwise/counterclockwise using fractions.
	// This needs to be more well documented...
	// allSisons := asteroids.sweepClockwise(bestAsteroid)
	// for sisonPt := range allSisons {
	// 	// fmt.Println(sisonPt)
	// }

	answer2 = fmt.Sprint(bestAsteroid)

	return
}
