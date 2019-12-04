package aoc2019

import "testing"

func Test_checkAdjacent(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"12", 12, false},
		{"11", 11, true},
		{"123455", 123455, true},
		{"124464", 124464, true},
		{"123456", 123456, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkAdjacent(tt.input); got != tt.want {
				t.Errorf("checkAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkNonDecreasing(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"12", 12, true},
		{"11", 11, true},
		{"10", 10, false},
		{"123455", 123455, true},
		{"124464", 124464, false},
		{"123456", 123456, true},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkNonDecreasing(tt.input); got != tt.want {
				t.Errorf("checkNonDecreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkDuality(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{"112233", 112233, true},
		{"123444", 123444, false},
		{"111122", 111122, true},
		{"123456", 123456, false},
		{"123345", 123345, true},
		{"122333", 122333, true},
		{"114444", 114444, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDuality(tt.input); got != tt.want {
				t.Errorf("checkDuality() = %v, want %v", got, tt.want)
			}
		})
	}
}
