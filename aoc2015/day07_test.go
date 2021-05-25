package aoc2015

import (
	"reflect"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

func Test_parseParams(t *testing.T) {
	// populate memory
	mem := make(map[identifier]signal)
	mem["a"] = 34
	mem["b"] = 69
	type args struct {
		mem    map[identifier]signal
		params []identifier
	}
	tests := []struct {
		name    string
		args    args
		want    []signal
		WantErr bool
	}{
		{
			"one identifier and one signal",
			args{
				mem,
				[]identifier{"10", "a"}},
			[]signal{10, 34},
			false,
		},
		{
			"something is not found in memory",
			args{
				mem,
				[]identifier{"4", "c"}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseParams(tt.args.mem, tt.args.params)
			if (err != nil) != tt.WantErr {
				t.Errorf("parseParams() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeNumbers(t *testing.T) {
	type args struct {
		params []identifier
	}
	tests := []struct {
		name string
		args args
		want []identifier
	}{
		{
			"only one test for this",
			args{[]identifier{"a", "4", "c"}},
			[]identifier{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNumbers(tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assign(t *testing.T) {
	assert := assert.New(t)
	mem := make(map[identifier]signal)
	mem["a"] = 16
	type args struct {
		mem     map[identifier]signal
		address identifier
		params  []identifier
	}
	tests := []struct {
		name    string
		args    args
		WantErr bool
	}{
		{
			"params not length 1",
			args{mem, "b", []identifier{"a", "43"}},
			true,
		},
		{
			"successfully assign b with val 43",
			args{mem, "b", []identifier{"43"}},
			false,
		},
		{
			"successfully assign c with val 16",
			args{mem, "c", []identifier{"a"}},
			false,
		},
		{
			"doesn't have a value yet!",
			args{mem, "d", []identifier{"f"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := assign(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.WantErr {
				t.Errorf("assign() error = %v, WantErr %v", err, tt.WantErr)
			}
		})
	}
	assert.Equal(mem,
		map[identifier]signal{
			"a": 16,
			"b": 43,
			"c": 16,
		},
		"check map equality")
}

func Test_and(t *testing.T) {
	assert := assert.New(t)
	mem := map[identifier]signal{
		"a": 2,
		"b": 5,
	}
	type args struct {
		mem     map[identifier]signal
		address identifier
		params  []identifier
	}
	tests := []struct {
		name    string
		args    args
		WantErr bool
	}{
		{
			"two Identifiers",
			args{mem, "c", []identifier{"a", "b"}},
			false,
		},
		{
			"identifier and signal",
			args{mem, "d", []identifier{"b", "9"}},
			false,
		},
		{
			"three parameters",
			args{mem, "d", []identifier{"a", "b", "c"}},
			true,
		},
		{
			"memory does not exist",
			args{mem, "d", []identifier{"e", "b"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := and(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.WantErr {
				t.Errorf("and() error = %v, WantErr %v", err, tt.WantErr)
			}
		})
	}
	assert.Equal(mem, map[identifier]signal{
		"a": 2,
		"b": 5,
		"c": 2 & 5,
		"d": 5 & 9,
	}, "result of and statement")
}

func Test_or(t *testing.T) {
	assert := assert.New(t)
	mem := map[identifier]signal{
		"a": 2,
		"b": 5,
	}
	type args struct {
		mem     map[identifier]signal
		address identifier
		params  []identifier
	}
	tests := []struct {
		name    string
		args    args
		WantErr bool
	}{
		{
			"two Identifiers",
			args{mem, "c", []identifier{"a", "b"}},
			false,
		},
		{
			"identifier and signal",
			args{mem, "d", []identifier{"b", "9"}},
			false,
		},
		{
			"three parameters",
			args{mem, "d", []identifier{"a", "b", "c"}},
			true,
		},
		{
			"memory does not exist",
			args{mem, "d", []identifier{"e", "b"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := or(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.WantErr {
				t.Errorf("or() error = %v, WantErr %v", err, tt.WantErr)
			}
		})
	}
	assert.Equal(mem, map[identifier]signal{
		"a": 2,
		"b": 5,
		"c": 2 | 5,
		"d": 5 | 9,
	})
}

func Test_not(t *testing.T) {
	assert := assert.New(t)
	mem := map[identifier]signal{
		"a": 143,
	}

	type args struct {
		mem     map[identifier]signal
		address identifier
		params  []identifier
	}
	tests := []struct {
		name    string
		args    args
		WantErr bool
	}{
		{
			"identifier",
			args{mem, "b", []identifier{"a"}},
			false,
		},
		{
			"signal",
			args{mem, "c", []identifier{"43"}},
			false,
		},
		{
			"two parameters",
			args{mem, "d", []identifier{"b", "c"}},
			true,
		},
		{
			"memory does not exist",
			args{mem, "e", []identifier{"d"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := not(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.WantErr {
				t.Errorf("not() error = %v, WantErr %v", err, tt.WantErr)
			}
		})
	}
	assert.Equal(mem, map[identifier]signal{
		"a": 143,
		"b": ^signal(143),
		"c": ^signal(43),
	})
}

func Test_lshift(t *testing.T) {
	assert := assert.New(t)
	mem := map[identifier]signal{
		"a": 2,
		"b": 5,
	}
	type args struct {
		mem     map[identifier]signal
		address identifier
		params  []identifier
	}
	tests := []struct {
		name    string
		args    args
		WantErr bool
	}{
		{
			"two Identifiers",
			args{mem, "c", []identifier{"a", "b"}},
			false,
		},
		{
			"identifier and signal",
			args{mem, "d", []identifier{"b", "9"}},
			false,
		},
		{
			"three parameters",
			args{mem, "d", []identifier{"a", "b", "c"}},
			true,
		},
		{
			"memory does not exist",
			args{mem, "d", []identifier{"e", "b"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lshift(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.WantErr {
				t.Errorf("lshift() error = %v, WantErr %v", err, tt.WantErr)
			}
		})
	}
	assert.Equal(mem, map[identifier]signal{
		"a": 2,
		"b": 5,
		"c": 2 << 5,
		"d": 5 << 9,
	})
}

func Test_rshift(t *testing.T) {
	assert := assert.New(t)
	mem := map[identifier]signal{
		"a": 2,
		"b": 132,
	}
	type args struct {
		mem     map[identifier]signal
		address identifier
		params  []identifier
	}
	tests := []struct {
		name    string
		args    args
		WantErr bool
	}{
		{
			"two Identifiers",
			args{mem, "c", []identifier{"b", "a"}},
			false,
		},
		{
			"identifier and signal",
			args{mem, "d", []identifier{"b", "9"}},
			false,
		},
		{
			"three parameters",
			args{mem, "d", []identifier{"a", "b", "c"}},
			true,
		},
		{
			"memory does not exist",
			args{mem, "d", []identifier{"e", "b"}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := rshift(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.WantErr {
				t.Errorf("rshift() error = %v, WantErr %v", err, tt.WantErr)
			}
		})
	}
	assert.Equal(mem, map[identifier]signal{
		"a": 2,
		"b": 132,
		"c": 132 >> 2,
		"d": 132 >> 9,
	})
}

func Test_parseInstruction(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    *instruction
		WantErr bool
	}{
		{
			"test assign instruction `lv -> a`",
			args{"lv -> b"},
			&instruction{
				"b",
				[]identifier{"lv"},
				[]identifier{"lv"},
				assign,
				false},
			false,
		},
		{
			"test and instruction `5 AND x -> b`",
			args{"5 AND x -> b"},
			&instruction{
				"b",
				[]identifier{"5", "x"},
				[]identifier{"x"},
				and,
				false},
			false,
		},
		{
			"test or instruction `o OR b -> c`",
			args{"o OR b -> c"},
			&instruction{
				"c",
				[]identifier{"o", "b"},
				[]identifier{"o", "b"},
				or,
				false},
			false,
		},
		{
			"test not instruction `NOT x -> d`",
			args{"NOT x -> d"},
			&instruction{
				"d",
				[]identifier{"x"},
				[]identifier{"x"},
				not,
				false},
			false,
		},
		{
			"test lshift instruction `o LSHIFT 4 -> c`",
			args{"o LSHIFT 4 -> c"},
			&instruction{
				"c",
				[]identifier{"o", "4"},
				[]identifier{"o"},
				lshift,
				false},
			false,
		},
		{
			"test rshift instruction `2 RSHIFT b -> c`",
			args{"2 RSHIFT b -> c"},
			&instruction{
				"c",
				[]identifier{"2", "b"},
				[]identifier{"b"},
				rshift,
				false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInstruction(tt.args.raw)
			if (err != nil) != tt.WantErr {
				t.Errorf("parseInstruction() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if !reflect.DeepEqual(got.id, tt.want.id) {
				t.Errorf("parseInstruction().id = %v, want %v", got.id, tt.want.id)
			}
			if !reflect.DeepEqual(got.parameters, tt.want.parameters) {
				t.Errorf("parseInstruction().parameters = %v, want %v", got.parameters, tt.want.parameters)
			}
			if !reflect.DeepEqual(got.childrenID, tt.want.childrenID) {
				t.Errorf("parseInstruction().childrenID = %v, want %v", got.childrenID, tt.want.childrenID)
			}
			// right now we can't compare functions
			if !reflect.DeepEqual(got.operationName(), tt.want.operationName()) {
				t.Errorf("parseInstruction().function = %v, want %v", got.operationName(), tt.want.operationName())
			}
		})
	}
}
func Test_instruction_String(t *testing.T) {
	tests := []struct {
		name  string
		instr *instruction
		want  string
	}{
		// no cases for now...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.instr.String(); got != tt.want {
				t.Errorf("instruction.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instruction_operationName(t *testing.T) {
	tests := []struct {
		name  string
		instr *instruction
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.instr.operationName(); got != tt.want {
				t.Errorf("instruction.operationName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	testCases := []internal.TestCase{
		{Details: "Y2016D07 my input",
			Input:   day07myInput,
			Result1: "3176",
			Result2: "14710"},
	}
	for _, tt := range testCases {
		tt.Test(Day07, assert)
	}
}

func BenchmarkDay07(b *testing.B) {
	internal.Benchmark(Day07, b, day07myInput)
}
