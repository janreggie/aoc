package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// conwayCube represents a Conway cube
type conwayCube [][][]bool

func (cube conwayCube) String() string {
	var sb strings.Builder

	for zz := range cube {
		for yy := range cube[zz] {
			for xx := range cube[zz][yy] {
				if cube[zz][yy][xx] {
					sb.WriteByte('#')
				} else {
					sb.WriteByte('.')
				}
			}
			sb.WriteByte('\n')
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

// width returns the width (x-dim) of a cube.
func (cube conwayCube) width() int {
	if cube.depth() == 0 || cube.height() == 0 {
		return 0
	}
	return len(cube[0][0])
}

// height returns the height (y-dim) of a cube.
func (cube conwayCube) height() int {
	if cube.depth() == 0 {
		return 0
	}
	return len(cube[0])
}

// depth returns the depth (z-dim) of a cube.
func (cube conwayCube) depth() int {
	return len(cube)
}

// volume returns the volume of the cube.
func (cube conwayCube) volume() int {
	return cube.width() * cube.height() * cube.depth()
}

// countActive returns how many active cubes are there
func (cube conwayCube) countActive() int {
	result := 0
	for xx := 0; xx < cube.width(); xx++ {
		for yy := 0; yy < cube.height(); yy++ {
			for zz := 0; zz < cube.depth(); zz++ {
				if cube.at(xx, yy, zz) {
					result++
				}
			}
		}
	}
	return result
}

// newConwayCube returns a new Conway Cube from a string
// whose width each line is equal and only contains `#` and `.`.
func newConwayCube(input string) (conwayCube, error) {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		// Well... we have an empty cube!
		return [][][]bool{{{false}}}, nil
	}

	cube := make([][][]bool, 1)
	cube[0] = make([][]bool, len(lines))
	width := len(lines[0])
	for ii, line := range lines {
		if len(line) != width {
			return nil, errors.Errorf("line %s is length %d, expected %d", line, len(line), width)
		}
		cube[0][ii] = make([]bool, width)
		for jj, pt := range line {
			switch pt {
			case '#':
				cube[0][ii][jj] = true
			case '.':
				cube[0][ii][jj] = false
			default:
				return nil, errors.Errorf("invalid character %c in line %s", pt, line)
			}
		}
	}
	return cube, nil
}

// blankConwayCube makes a blank conwayCube with some width, height, and depth
func blankConwayCube(w, h, d int) conwayCube {
	result := make(conwayCube, d)
	for zz := range result {
		result[zz] = make([][]bool, h)
		for yy := range result[zz] {
			result[zz][yy] = make([]bool, w)
		}
	}
	return result
}

// reduce reduces the cube and returns a new cube
func (cube conwayCube) reduce() conwayCube {
	// every element in the cube is checked
	// and {min|max}{X|Y|Z} is narrowed down
	minX, minY, minZ := cube.width(), cube.height(), cube.depth()
	maxX, maxY, maxZ := 0, 0, 0
	for zz := 0; zz < cube.depth(); zz++ {
		for yy := 0; yy < cube.height(); yy++ {
			for xx := 0; xx < cube.width(); xx++ {
				if !cube.at(xx, yy, zz) {
					continue
				}

				// update mins and maxs
				if xx < minX {
					minX = xx
				}
				if xx > maxX {
					maxX = xx
				}
				if yy < minY {
					minY = yy
				}
				if yy > maxY {
					maxY = yy
				}
				if zz < minZ {
					minZ = zz
				}
				if zz > maxZ {
					maxZ = zz
				}
			}
		}
	}

	// -1 since maxX goes from 0 to cube.width()-1, and similarly maxY and maxZ
	if maxX-minX == cube.width()-1 &&
		maxY-minY == cube.height()-1 &&
		maxZ-minZ == cube.depth()-1 {
		return cube
	}

	newCube := blankConwayCube(maxX-minX+1, maxY-minY+1, maxZ-minZ+1)
	for xx := minX; xx <= maxX; xx++ {
		for yy := minY; yy <= maxY; yy++ {
			for zz := minZ; zz <= maxZ; zz++ {
				newCube.set(xx-minX, yy-minY, zz-minZ, cube.at(xx, yy, zz))
			}
		}
	}

	return newCube
}

// at represents if there is a value at some point in the cube.
// If x, y, or z is out of bounds, return false.
func (cube conwayCube) at(x, y, z int) bool {
	if z < 0 || z >= cube.depth() ||
		y < 0 || y >= cube.height() ||
		x < 0 || x >= cube.width() {
		return false
	}
	return cube[z][y][x]
}

// set sets the value of some position to val.
// If x, y, or z is out of bounds, do nothing.
func (cube *conwayCube) set(x, y, z int, val bool) {
	if z < 0 || z >= len(*cube) ||
		y < 0 || y >= len((*cube)[z]) ||
		x < 0 || x >= len((*cube)[z][y]) {
		return
	}
	(*cube)[z][y][x] = val
}

// surrounds counts the number of active cubes that surround (x,y,z)
func (cube conwayCube) surrounds(x, y, z int) int {
	countTrue := func(bs ...bool) int {
		result := 0
		for _, bb := range bs {
			if bb {
				result++
			}
		}
		return result
	}

	return countTrue(
		cube.at(x-1, y-1, z-1), cube.at(x, y-1, z-1), cube.at(x+1, y-1, z-1),
		cube.at(x-1, y, z-1), cube.at(x, y, z-1), cube.at(x+1, y, z-1),
		cube.at(x-1, y+1, z-1), cube.at(x, y+1, z-1), cube.at(x+1, y+1, z-1),

		cube.at(x-1, y-1, z), cube.at(x, y-1, z), cube.at(x+1, y-1, z),
		cube.at(x-1, y, z), false, cube.at(x+1, y, z),
		cube.at(x-1, y+1, z), cube.at(x, y+1, z), cube.at(x+1, y+1, z),

		cube.at(x-1, y-1, z+1), cube.at(x, y-1, z+1), cube.at(x+1, y-1, z+1),
		cube.at(x-1, y, z+1), cube.at(x, y, z+1), cube.at(x+1, y, z+1),
		cube.at(x-1, y+1, z+1), cube.at(x, y+1, z+1), cube.at(x+1, y+1, z+1),
	)
}

// iterate returns a new conwayCube based on the original after iterating based on the rules
func (cube conwayCube) iterate() conwayCube {
	newCube := make(conwayCube, len(cube)+2)
	for ii := range newCube {
		newCube[ii] = make([][]bool, len(cube[0])+2)
		for jj := range newCube[ii] {
			newCube[ii][jj] = make([]bool, len(cube[0][0])+2)
		}
	}

	// newCube.at(x+1,y+1,z+1) <-> cube.at(x,y,z)
	for zz := range newCube {
		for yy := range newCube[zz] {
			for xx := range newCube[zz][yy] {
				val := cube.at(xx-1, yy-1, zz-1)
				sur := cube.surrounds(xx-1, yy-1, zz-1)
				if (val && (sur == 2 || sur == 3)) || (!val && sur == 3) {
					newCube.set(xx, yy, zz, true)
				} else {
					newCube.set(xx, yy, zz, false) // unneeded really since default is false
				}
			}
		}
	}

	return newCube.reduce()
}

// Day17 solves the seventeenth day puzzle "Conway Cubes"
//
// Input
//
// A 2D grid representing the initial state of the Conway cube. For example:
//
//  some sample input indented to become a code block
//
// It is guaranteed that the input only contains `.` and `#`'s.
func Day17(scanner *bufio.Scanner) (answer1, answer2 string, err error) {
	// Code goes here.
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text())
		sb.WriteString("\n")
	}
	input := strings.TrimSpace(sb.String())
	cube, err := newConwayCube(input)
	if err != nil {
		err = errors.Wrapf(err, "could not derive cube from input %s", input)
		return
	}
	hcube, err := newConwayHypercube(input)
	if err != nil {
		err = errors.Wrapf(err, "could not derive hypercube from input %s", input)
		return
	}

	for ii := 0; ii < 6; ii++ {
		cube = cube.iterate()
		hcube = hcube.iterate()
	}
	answer1 = strconv.Itoa(cube.countActive())
	answer2 = strconv.Itoa(hcube.countActive())
	return
}
