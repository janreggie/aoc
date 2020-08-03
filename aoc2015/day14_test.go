package aoc2015

import (
	"testing"

	"github.com/janreggie/AdventOfCode/internal"
	"github.com/stretchr/testify/assert"
)

const day14myInput = `Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.
Rudolph can fly 3 km/s for 15 seconds, but then must rest for 28 seconds.
Donner can fly 19 km/s for 9 seconds, but then must rest for 164 seconds.
Blitzen can fly 19 km/s for 9 seconds, but then must rest for 158 seconds.
Comet can fly 13 km/s for 7 seconds, but then must rest for 82 seconds.
Cupid can fly 25 km/s for 6 seconds, but then must rest for 145 seconds.
Dasher can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.
Dancer can fly 3 km/s for 16 seconds, but then must rest for 37 seconds.
Prancer can fly 25 km/s for 6 seconds, but then must rest for 143 seconds.
`

func Test_newRacingReindeer(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		description string
		want        racingReindeer
	}{
		{description: "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
			want: racingReindeer{
				name:        "Comet",
				flyingSpeed: 14,
				flyingTime:  10,
				restingTime: 127,
			},
		},
	}
	for _, tt := range tests {
		result, err := newRacingReindeer(tt.description)
		assert.NoError(err, tt.description)
		assert.Equal(tt.want, result, tt.description)
	}
}

func Test_racingReindeer_distance_isFlying(t *testing.T) {
	assert := assert.New(t)
	comet, err := newRacingReindeer("Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.")
	assert.NoError(err, "generating comet")
	dancer, err := newRacingReindeer("Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.")
	assert.NoError(err, "generating dancer")

	tests := []struct {
		time         uint
		cometDist    uint
		cometFlying  bool
		dancerDist   uint
		dancerFlying bool
	}{
		{time: 1,
			cometDist:    14,
			cometFlying:  true,
			dancerDist:   16,
			dancerFlying: true,
		},
		{time: 10,
			cometDist:    140,
			cometFlying:  true,
			dancerDist:   160,
			dancerFlying: true,
		},
		{time: 11,
			cometDist:    140,
			cometFlying:  false,
			dancerDist:   176,
			dancerFlying: true,
		},
		{time: 12,
			cometDist:    140,
			cometFlying:  false,
			dancerDist:   176,
			dancerFlying: false,
		},
		{time: 138,
			cometDist:    154,
			cometFlying:  true,
			dancerDist:   176,
			dancerFlying: false,
		},
		{time: 1000,
			cometDist:    1120,
			cometFlying:  false,
			dancerDist:   1056,
			dancerFlying: false,
		},
	}
	for _, tt := range tests {
		assert.Equalf(tt.cometDist, comet.distance(tt.time), "comet distance %v", tt.time)
		assert.Equalf(tt.dancerDist, dancer.distance(tt.time), "dancer distance %v", tt.time)
		assert.Equalf(tt.cometFlying, comet.isFlying(tt.time), "comet flying %v", tt.time)
		assert.Equalf(tt.dancerFlying, dancer.isFlying(tt.time), "dancer flying %v", tt.time)
	}
}

func Test_reindeerOlympics_iterate(t *testing.T) {
	assert := assert.New(t)
	comet, err := newRacingReindeer("Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.")
	assert.NoError(err, "generating comet")
	dancer, err := newRacingReindeer("Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.")
	assert.NoError(err, "generating dancer")
	olympics := newReindeerOlympics([]racingReindeer{comet, dancer})

	type distances struct {
		cometDistance  uint
		dancerDistance uint
	}
	type points struct {
		cometPoints  uint
		dancerPoints uint
	}

	tests := []struct {
		time          uint
		wantDistances distances
		wantPoints    points
	}{
		{time: 1,
			wantDistances: distances{14, 16},
			wantPoints:    points{0, 1},
		},
		{time: 10,
			wantDistances: distances{140, 160},
			wantPoints:    points{0, 10},
		},
		{time: 11,
			wantDistances: distances{140, 176},
			wantPoints:    points{0, 11},
		},
		{time: 12,
			wantDistances: distances{140, 176},
			wantPoints:    points{0, 12},
		},
		{time: 138,
			wantDistances: distances{154, 176},
			wantPoints:    points{0, 138},
		},
		{time: 1000,
			wantDistances: distances{1120, 1056},
			wantPoints:    points{312, 689},
		},
	}
	for _, tt := range tests {
		olympics.iterateUntil(tt.time)
		assert.Equalf(tt.wantDistances, distances{olympics.travelled[0], olympics.travelled[1]}, "travelled after %v sec", tt.time)
		assert.Equalf(tt.wantPoints, points{olympics.points[0], olympics.points[1]}, "points after %v sec", tt.time)
		olympics.reset()
	}
}

func TestDay14(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2015D14 my input",
			Input:   day14myInput,
			Result1: "2660",
			Result2: "1256"},
	}
	for _, tt := range testCases {
		tt.Test(Day14, assert)
	}
}

func BenchmarkDay14(b *testing.B) {
	internal.Benchmark(Day14, b, day14myInput)
}
