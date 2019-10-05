package aoc2015

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isInteger(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"normal number", args{"14"}, 14, true},
		{"not a number", args{"a"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := isInteger(tt.args.str)
			if got != tt.want {
				t.Errorf("isInteger() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("isInteger() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
		wantErr bool
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
			if (err != nil) != tt.wantErr {
				t.Errorf("parseParams() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr bool
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
			if err := assign(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("assign() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr bool
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
			if err := and(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("and() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr bool
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
			if err := or(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("or() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr bool
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
			if err := not(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("not() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr bool
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
			if err := lshift(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("lshift() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr bool
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
			if err := rshift(tt.args.mem, tt.args.address, tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("rshift() error = %v, wantErr %v", err, tt.wantErr)
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
		wantErr bool
	}{
		{
			"test assign instruction `lv -> a`",
			args{"lv -> b"},
			&instruction{
				"b",
				[]identifier{"lv"},
				[]identifier{"lv"},
				make([]*instruction, 0),
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
				make([]*instruction, 0),
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
				make([]*instruction, 0),
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
				make([]*instruction, 0),
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
				make([]*instruction, 0),
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
				make([]*instruction, 0),
				rshift,
				false},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseInstruction(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseInstruction() error = %v, wantErr %v", err, tt.wantErr)
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
			if !reflect.DeepEqual(got.children, tt.want.children) {
				t.Errorf("parseInstruction().children = %v, want %v", got.children, tt.want.children)
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

func Test_instruction_findChildren(t *testing.T) {
	// let us use test case in
	// https://adventofcode.com/2015/day/7
	assert := assert.New(t)
	imap := newInstructionMap()
	input := `123 -> x
	456 -> y
	x AND y -> d
	x OR y -> e
	x LSHIFT 2 -> f
	y RSHIFT 2 -> g
	NOT x -> h
	NOT y -> i`
	for _, eachLine := range strings.Split(input, "\n") {
		instr, err := parseInstruction(eachLine)
		assert.NoError(err)
		imap.append(instr)
	}
	instrOnly := func(instr *instruction, err error) *instruction {
		assert.NoError(err, "make sure there is no error in returning only *instruction")
		return instr
	}
	lookup := func(id identifier) *instruction {
		return instrOnly(imap.lookup(id))
	}
	checkChildren := func(instr *instruction) {
		for _, eachID := range instr.childrenID {
			assert.Contains(instr.children, lookup(eachID), "check children")
		}
	}
	type args struct {
		imap *instructionMap
	}
	instrD := lookup("d")
	instrG := lookup("g")
	tests := []struct {
		name    string
		instr   *instruction
		args    args
		wantErr bool
	}{
		{
			"find children of d",
			instrD,
			args{imap},
			false,
		},
		{
			"find children of g",
			instrG,
			args{imap},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.instr.findChildren(tt.args.imap); (err != nil) != tt.wantErr {
				t.Errorf("instruction.findChildren() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		// now test whether they have children
		checkChildren(tt.instr)
	}
}

func Test_instruction_populateChildren(t *testing.T) {
	assert := assert.New(t)
	imap := newInstructionMap()
	input := `123 -> x
	456 -> y
	x AND y -> d
	x OR y -> e
	d LSHIFT 2 -> f
	y RSHIFT 2 -> g
	NOT f -> h
	NOT g -> i` // modified slightly to be more dependent on each other
	for _, eachLine := range strings.Split(input, "\n") {
		instr, err := parseInstruction(eachLine)
		assert.NoError(err)
		imap.append(instr)
	}
	instrOnly := func(instr *instruction, err error) *instruction {
		assert.NoError(err, "make sure there is no error")
		return instr
	}
	lookup := func(id identifier) *instruction {
		return instrOnly(imap.lookup(id))
	}
	checkChildren := func(instr *instruction) {
		for _, eachID := range instr.childrenID {
			assert.Contains(instr.children, lookup(eachID),
				fmt.Sprintf("check %v is in %v", eachID, instr.id))
		}
	}
	type args struct {
		imap *instructionMap
	}
	instrH := lookup("h")
	instrI := lookup("i")
	tests := []struct {
		name    string
		instr   *instruction
		args    args
		wantErr bool
	}{
		{
			"populate tree with parent h",
			instrH,
			args{imap},
			false,
		},
		{
			"populate tree with parent i",
			instrI,
			args{imap},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.instr.populateChildren(tt.args.imap); (err != nil) != tt.wantErr {
				t.Errorf("instruction.populateChildren() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	imap.traverse(func(id identifier) {
		// supposed to fail if id == "e"
		// since "e" is not necessary to determine
		// values of H and I
		if id != "e" {
			checkChildren(lookup(id))
		}
	})
}

func Test_instruction_performInstruction(t *testing.T) {
	assert := assert.New(t)
	imap := newInstructionMap()
	mem := make(map[identifier]signal)
	type args struct {
		mem map[identifier]signal
	}
	input := `123 -> x
	456 -> y
	x AND y -> d
	x OR y -> e
	x LSHIFT 2 -> f
	y RSHIFT 2 -> g
	NOT x -> h
	NOT y -> i`
	tests := []struct {
		name    string
		instr   *instruction
		args    args
		wantErr bool
	}{}
	for _, eachLine := range strings.Split(input, "\n") {
		instr, err := parseInstruction(eachLine)
		assert.NoError(err)
		imap.append(instr)
		// note that eachLine is sequential
		tests = append(tests,
			struct {
				name    string
				instr   *instruction
				args    args
				wantErr bool
			}{

				"perform instruction " + string(instr.id),
				instr,
				args{mem},
				false,
			},
		)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.instr.performInstruction(tt.args.mem); (err != nil) != tt.wantErr {
				t.Errorf("instruction.performInstruction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	// now check whether mem is expected
	assert.Equal(mem,
		map[identifier]signal{
			"d": 72,
			"e": 507,
			"f": 492,
			"g": 114,
			"h": 65412,
			"i": 65079,
			"x": 123,
			"y": 456,
		})
}

func Test_instruction_performAll(t *testing.T) {
	assert := assert.New(t)
	imap := newInstructionMap()
	mem := make(map[identifier]signal)
	input := `123 -> x
	456 -> y
	x AND y -> d
	x OR y -> e
	d LSHIFT 2 -> f
	y RSHIFT 2 -> g
	NOT f -> h
	NOT g -> i` // modified slightly to be more dependent on each other
	for _, eachLine := range strings.Split(input, "\n") {
		instr, err := parseInstruction(eachLine)
		assert.NoError(err)
		imap.append(instr)
	}
	instrOnly := func(instr *instruction, err error) *instruction {
		assert.NoError(err, "make sure there is no error")
		return instr
	}
	lookup := func(id identifier) *instruction {
		return instrOnly(imap.lookup(id))
	}
	type args struct {
		mem map[identifier]signal
	}
	tests := []struct {
		name    string
		instr   *instruction
		args    args
		wantErr bool
	}{
		{
			"perform tree with node h",
			lookup("h"),
			args{mem},
			false,
		},
		{
			"perform tree with node i",
			lookup("i"),
			args{mem},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(tt.instr.populateChildren(imap))
			if err := tt.instr.performAll(tt.args.mem); (err != nil) != tt.wantErr {
				t.Errorf("instruction.performAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	// now determine whether map has values
	assert.Equal(mem, map[identifier]signal{
		"x": 123,
		"y": 456,
		"d": 72, // 123&456
		// "e": 507, // this is NOT important in determining h and i
		"f": 288,   //72<<2
		"g": 114,   // 456>>2
		"h": 65247, // ^288
		"i": 65421, // ^114
	})
}

func Test_newInstructionMap(t *testing.T) {
	tests := []struct {
		name string
		want *instructionMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newInstructionMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newInstructionMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instructionMap_String(t *testing.T) {
	tests := []struct {
		name string
		imap *instructionMap
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.imap.String(); got != tt.want {
				t.Errorf("instructionMap.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_instructionMap_append(t *testing.T) {
	type args struct {
		is *instruction
	}
	tests := []struct {
		name string
		imap *instructionMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.imap.append(tt.args.is)
		})
	}
}

func Test_instructionMap_delete(t *testing.T) {
	type args struct {
		id identifier
	}
	tests := []struct {
		name string
		imap *instructionMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.imap.delete(tt.args.id)
		})
	}
}

func Test_instructionMap_lookup(t *testing.T) {
	type args struct {
		id identifier
	}
	tests := []struct {
		name    string
		imap    *instructionMap
		args    args
		want    *instruction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.imap.lookup(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("instructionMap.lookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instructionMap.lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
