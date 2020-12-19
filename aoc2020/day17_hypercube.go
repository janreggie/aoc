package aoc2020

// day17_hypercube is a helper for day17 to isolate the "hypercube" component (part 2).

import (
	"strings"

	"github.com/pkg/errors"
)

// conwaHypercube represents a Conway hypercube
type conwayHypercube [][][][]bool

func (hcube conwayHypercube) String() string {
	var sb strings.Builder
	for ww := range hcube {
		for zz := range hcube[ww] {
			for yy := range hcube[ww][zz] {
				for xx := range hcube[ww][zz][yy] {
					if hcube[ww][zz][yy][xx] {
						sb.WriteByte('#')
					} else {
						sb.WriteByte('.')
					}
				}
				sb.WriteByte('\n')
			}
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

// width returns the width (x-dim) of a hcube.
func (hcube conwayHypercube) width() int {
	if hcube.height() == 0 {
		return 0
	}
	return len(hcube[0][0][0])
}

// height returns the height (y-dim) of a hcube.
func (hcube conwayHypercube) height() int {
	if hcube.depth() == 0 {
		return 0
	}
	return len(hcube[0][0])
}

// depth returns the depth (z-dim) of a hcube.
func (hcube conwayHypercube) depth() int {
	if hcube.hyperdepth() == 0 {
		return 0
	}
	return len(hcube[0])
}

// hyperdepth returns the hyperdepth (w-dim) of a hcube
func (hcube conwayHypercube) hyperdepth() int {
	return len(hcube)
}

// volume returns the volume of the hcube.
func (hcube conwayHypercube) volume() int {
	return hcube.hyperdepth() * hcube.width() * hcube.height() * hcube.depth()
}

// countActive returns how many active hcubes are there
func (hcube conwayHypercube) countActive() int {
	result := 0
	for xx := 0; xx < hcube.width(); xx++ {
		for yy := 0; yy < hcube.height(); yy++ {
			for zz := 0; zz < hcube.depth(); zz++ {
				for ww := 0; ww < hcube.hyperdepth(); ww++ {
					if hcube.at(xx, yy, zz, ww) {
						result++
					}
				}
			}
		}
	}
	return result
}

// newConwayHypercube returns a new Conway Cube from a string
// whose width each line is equal and only contains `#` and `.`.
func newConwayHypercube(input string) (conwayHypercube, error) {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		// Well... we have an empty hcube!
		return blankConwayHypercube(0, 0, 0, 0), nil
	}

	hcube := make([][][][]bool, 1)
	hcube[0] = make([][][]bool, 1)
	hcube[0][0] = make([][]bool, len(lines))
	width := len(lines[0])
	for ii, line := range lines {
		if len(line) != width {
			return nil, errors.Errorf("line %s is length %d, expected %d", line, len(line), width)
		}
		hcube[0][0][ii] = make([]bool, width)
		for jj, pt := range line {
			switch pt {
			case '#':
				hcube[0][0][ii][jj] = true
			case '.':
				hcube[0][0][ii][jj] = false
			default:
				return nil, errors.Errorf("invalid character %c in line %s", pt, line)
			}
		}
	}
	return hcube, nil
}

// blankConwayCube makes a blank conwayCube with some width, height, and depth
func blankConwayHypercube(w, h, d, hd int) conwayHypercube {
	result := make(conwayHypercube, hd)
	for ww := range result {
		result[ww] = make([][][]bool, d)
		for zz := range result[ww] {
			result[ww][zz] = make([][]bool, h)
			for yy := range result[ww][zz] {
				result[ww][zz][yy] = make([]bool, w)
			}
		}
	}
	return result
}

// reduce reduces the hcube and returns a new hcube
func (hcube conwayHypercube) reduce() conwayHypercube {
	// every element in the hcube is checked
	// and {min|max}{X|Y|Z} is narrowed down
	minX, minY, minZ, minW := hcube.width(), hcube.height(), hcube.depth(), hcube.hyperdepth()
	maxX, maxY, maxZ, maxW := 0, 0, 0, 0
	for ww := 0; ww < hcube.hyperdepth(); ww++ {
		for zz := 0; zz < hcube.depth(); zz++ {
			for yy := 0; yy < hcube.height(); yy++ {
				for xx := 0; xx < hcube.width(); xx++ {
					if !hcube.at(xx, yy, zz, ww) {
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
					if ww < minW {
						minW = ww
					}
					if ww > maxW {
						maxW = ww
					}
				}
			}
		}
	}

	// -1 since maxX goes from 0 to hcube.width()-1, and similarly maxY and maxZ
	if maxX-minX == hcube.width()-1 &&
		maxY-minY == hcube.height()-1 &&
		maxZ-minZ == hcube.depth()-1 &&
		maxW-minW == hcube.hyperdepth()-1 {
		return hcube
	}

	newCube := blankConwayHypercube(maxX-minX+1, maxY-minY+1, maxZ-minZ+1, maxW-minW+1)
	for xx := minX; xx <= maxX; xx++ {
		for yy := minY; yy <= maxY; yy++ {
			for zz := minZ; zz <= maxZ; zz++ {
				for ww := minW; ww <= maxW; ww++ {
					newCube.set(xx-minX, yy-minY, zz-minZ, ww-minW, hcube.at(xx, yy, zz, ww))
				}
			}
		}
	}

	return newCube
}

// at represents if there is a value at some point in the hcube.
// If x, y, or z is out of bounds, return false.
func (hcube conwayHypercube) at(x, y, z, w int) bool {
	if w < 0 || w >= hcube.hyperdepth() ||
		z < 0 || z >= hcube.depth() ||
		y < 0 || y >= hcube.height() ||
		x < 0 || x >= hcube.width() {
		return false
	}
	return hcube[w][z][y][x]
}

// set sets the value of some position to val.
// If x, y, or z is out of bounds, do nothing.
func (hcube *conwayHypercube) set(x, y, z, w int, val bool) {
	if w < 0 || w >= hcube.hyperdepth() ||
		z < 0 || z >= hcube.depth() ||
		y < 0 || y >= hcube.height() ||
		x < 0 || x >= hcube.width() {
		return
	}
	(*hcube)[w][z][y][x] = val
}

// surrounds counts the number of active hcubes that surround (x,y,z)
func (hcube conwayHypercube) surrounds(x, y, z, w int) int {
	result := 0
	offsets := []int{-1, 0, 1}
	boolToInt := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	for _, ox := range offsets {
		for _, oy := range offsets {
			for _, oz := range offsets {
				for _, ow := range offsets {
					if ox == 0 && oy == 0 && oz == 0 && ow == 0 {
						continue
					}
					result = result + boolToInt(hcube.at(x+ox, y+oy, z+oz, w+ow))
				}
			}
		}
	}
	return result
}

// iterate returns a new conwayCube based on the original after iterating based on the rules
func (hcube conwayHypercube) iterate() conwayHypercube {
	newCube := blankConwayHypercube(hcube.width()+2, hcube.height()+2, hcube.depth()+2, hcube.hyperdepth()+2)

	// newCube.at(x+1,y+1,z+1,w+1) <-> hcube.at(x,y,z,w)
	for ww := 0; ww < newCube.hyperdepth(); ww++ {
		for zz := 0; zz < newCube.depth(); zz++ {
			for yy := 0; yy < newCube.height(); yy++ {
				for xx := 0; xx < newCube.width(); xx++ {
					val := hcube.at(xx-1, yy-1, zz-1, ww-1)
					sur := hcube.surrounds(xx-1, yy-1, zz-1, ww-1)
					if (val && (sur == 2 || sur == 3)) || (!val && sur == 3) {
						newCube.set(xx, yy, zz, ww, true)
					} else {
						newCube.set(xx, yy, zz, ww, false) // unneeded really since default is false
					}
				}
			}
		}
	}

	return newCube.reduce()
}
