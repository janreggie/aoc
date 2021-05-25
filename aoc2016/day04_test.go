package aoc2016

import (
	"testing"

	aoc "github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_room(t *testing.T) {
	assert := assert.New(t)
	realRoomInput := "aaaaa-bbb-z-y-x-123[abxyz]"
	realRoom, err := newRoom(realRoomInput)
	assert.NoErrorf(err, "%v should be a valid room", realRoomInput)
	assert.Equal(room{
		checksum: "abxyz",
		letters: map[byte]uint16{
			'a': 5,
			'b': 3,
			'x': 1,
			'y': 1,
			'z': 1,
		},
		raw:      realRoomInput,
		sectorID: 123,
	}, realRoom, "realRoom should be expected")
	assert.True(realRoom.isReal(), "realRoom should be real")

	notRealRoomInput := "totally-real-room-200[decoy]"
	notRealRoom, err := newRoom(notRealRoomInput)
	assert.NoErrorf(err, "%v should be a valid room", notRealRoomInput)
	assert.Equal(room{
		checksum: "decoy",
		letters: map[byte]uint16{
			'a': 2,
			'e': 1,
			'l': 3,
			'm': 1,
			'o': 3,
			'r': 2,
			't': 2,
			'y': 1,
		},
		raw:      notRealRoomInput,
		sectorID: 200,
	}, notRealRoom, "notRealRoom should be expected")
	assert.False(notRealRoom.isReal(), "notRealRoom should not be real")
}

func Test_room_decrypt(t *testing.T) {
	assert := assert.New(t)
	rm, err := newRoom("qzmt-zixmtkozy-ivhz-343[zimth]")
	assert.NoError(err)
	assert.True(rm.isReal(), rm)
	assert.Equal("very encrypted name", rm.decrypt(), "decrypted message")
}

func TestDay04(t *testing.T) {
	assert := assert.New(t)
	testCases := []aoc.TestCase{
		{Details: "Y2016D04 sample input",
			Input:   day04sampleInput,
			Result1: "1514"},
		{Details: "Y2016D04 my input",
			Input:   day04myInput,
			Result1: "278221",
			Result2: "267"},
	}
	for _, tt := range testCases {
		tt.Test(Day04, assert)
	}
}

func BenchmarkDay04(b *testing.B) {
	aoc.Benchmark(Day04, b, day04myInput)
}
