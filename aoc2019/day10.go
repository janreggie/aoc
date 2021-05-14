package aoc2019

import (
	"bufio"
	"fmt"
	"image"
	"math/big"
	"sort"
	"strconv"
	"strings"

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

// lcmUpTo generates the lcm of 1, 2, up to limit.
// If limit is less than 2, it will return limit
func lcmUpTo(limit int) int {
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
// If pt.X or pt.Y is 1, return true.
// If pt.X or pt.Y is nonpositive,
// or if they are equal to each other, return false.
func isCoprime(pt image.Point) bool {
	x := pt.X
	y := pt.Y
	if x <= 0 || y <= 0 { // negative?!
		return false
	}
	if x == 1 || y == 1 {
		return true
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

// get returns true if there is an asteroid get a position point(x,y),
// false if there is none or if point is out of bounds
func (field *asteroidField) get(point image.Point) bool {
	if !point.In(field.bounds()) {
		return false
	}
	return field.asteroids[point.Y][point.X]
}

// set sets the value at some point.
// If point is out of bounds it does nothing.
func (field *asteroidField) set(point image.Point, value bool) {
	if !point.In(field.bounds()) {
		return
	}
	field.asteroids[point.Y][point.X] = value
}

// bounds returns an image.Rect representing the bounds of the asteroid field
func (field *asteroidField) bounds() image.Rectangle {
	return image.Rect(0, 0, field.width, field.height)
}

// countLineOfSight counts the number of asteroids a position index sees in its line of sight.
func (field *asteroidField) countLineOfSight(index image.Point) int {
	count := 0
	for range field.sweepClockwise(index) {
		count++
	}
	return count
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
func generateSisonPoints(width, height int) chan image.Point {
	channel := make(chan image.Point, 4) // 4 is a good enough size

	go func() {
		// okay what if width and height are -1
		if width <= 0 && height <= 0 {
			close(channel)
			return
		}

		// if height is positive?
		channel <- image.Pt(0, 1)
		if width <= 0 {
			close(channel)
			return
		}

		// Let's grab all points from (1,1) to (width, height).
		// That's width*height points
		allPoints := make([]image.Point, 0, width*height/3) // good estimate
		for ii := 1; ii <= width; ii++ {
			for jj := 1; jj <= height; jj++ {
				pt := image.Pt(ii, jj)
				if isCoprime(pt) {
					allPoints = append(allPoints, pt)
				}
			}
		}

		// sort them all together...
		sort.Slice(allPoints, func(i, j int) bool {
			pi := allPoints[i]
			pj := allPoints[j]
			return (pi.X)*(pj.Y) < (pj.X)*(pi.Y)
		})

		for _, pt := range allPoints {
			channel <- pt
		}
		close(channel)
	}()

	return channel
}

// generateSisonPointsSlow generates Sison points slowly
// by brute forcing through lcmUpTo(width) and lcmUpTo(height).
func generateSisonPointsSlow(width, height int) chan image.Point {
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

		denom := lcmUpTo(height)
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
		denom = lcmUpTo(width)
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
	clockwiseGenerator := func() {
		func() { // generate for upper-right quadrant
			quarterWidth := field.width - relativeTo.X - 1
			quarterHeight := relativeTo.Y
			for sisonPoint := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(sisonPoint.X, -sisonPoint.Y)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					if field.get(absolute) {
						channel <- absolute
						break
					}
				}
			}
		}()

		func() { // generate for lower-right quadrant
			// As we are sweeping counterclockwise,
			// the height of the Sison Point Generator would be the horizontal distance based on field.width,...,
			// that is, field.width - relativeTo.X - 1,
			// and the width of the Generator would be the vertical distance based on field.height.
			quarterWidth := field.height - relativeTo.Y - 1
			quarterHeight := field.width - relativeTo.X - 1
			for sisonPont := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(+sisonPont.Y, +sisonPont.X)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					if field.get(absolute) {
						channel <- absolute
						break
					}
				}
			}
		}()

		func() { // generate for the lower-left quadrant
			quarterWidth := relativeTo.X
			quarterHeight := field.height - relativeTo.Y - 1
			for sisonPoint := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(-sisonPoint.X, +sisonPoint.Y)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					if field.get(absolute) {
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
			for sisonPoint := range generateSisonPoints(quarterWidth, quarterHeight) {
				addend := image.Pt(-sisonPoint.Y, -sisonPoint.X)
				for absolute := relativeTo.Add(addend); absolute.In(field.bounds()); absolute = absolute.Add(addend) {
					if field.get(absolute) {
						channel <- absolute
						break
					}
				}
			}
		}()

		close(channel) // oh right... close the channel.
	}
	go clockwiseGenerator()

	return channel
}

// sweepAsteroids vaporizes some number of asteroids from the north clockwise
// and returns the coordinates of the last vaporized asteroid.
// If toVaporize is 0 return image.Point{0,0}.
func (field *asteroidField) sweepAsteroids(station image.Point, toVaporize uint) image.Point {
	if toVaporize == 0 {
		return image.Pt(0, 0)
	}
	var vaporized uint
	for {
		for pt := range field.sweepClockwise(station) {
			field.set(pt, false)
			vaporized++
			if vaporized == toVaporize {
				return pt
			}
		}
	}
}

// Day10 solves the tenth day problem "Monitoring Station".
//
// Input
//
// An asteroid field containing HEIGHT lines, each of exactly WIDTH characters long
// and only containing the characters `#` or `.`. For example:
//
//	.#..##.###...#######
//	##.############..##.
//	.#.######.########.#
//	.###.#######.####.#.
//	#####.##.#.##.###.##
//	..#####..#.#########
//	####################
//	#.####....###.#.#.##
//	##.#################
//	#####.##.###..####..
//	..######..##.#######
//	####.##.####...##..#
//	.#####..#.######.###
//	##...#.##########...
//	#.##########.#######
//	.####.#.###.###.#.##
//	....##.##.###..#####
//	.#.#.###########.###
//	#.#.#.#####.####.###
//	###.##.####.##.#..##
//
// If the input contains any other character than the ones specified,
// or if it contains lines of uneven width, it may return an error.
//
// It is guaranteed that WIDTH and HEIGHT are both at least 1 and at most 30.
func Day10(input string) (answer1, answer2 string, err error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	asteroids, err := newAsteroidField(scanner)
	if err != nil {
		err = errors.Wrap(err, "Y2019D10: could not create asteroidField")
		return
	}

	mostSeen := 0
	bestAsteroid := image.Pt(0, 0)
	for ii := 0; ii < asteroids.width; ii++ {
		for jj := 0; jj < asteroids.height; jj++ {
			if !asteroids.get(image.Pt(ii, jj)) {
				continue // don't care about non-asteroids
			}
			// for range asteroids.sweepClockwise(image.Pt(ii, jj)) {
			// 	/// do nothing...
			// }
			if seen := asteroids.countLineOfSight(image.Pt(ii, jj)); seen > mostSeen {
				mostSeen = seen
				bestAsteroid = image.Pt(ii, jj)
			}
		}
	}
	answer1 = strconv.Itoa(mostSeen)

	// now for Part 2
	lastAsteroid := asteroids.sweepAsteroids(bestAsteroid, 200)
	answer2 = strconv.Itoa(lastAsteroid.X*100 + lastAsteroid.Y)

	return
}
