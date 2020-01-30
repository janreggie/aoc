package aoc2015

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/janreggie/AdventOfCode/tools"
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
		WantErr bool
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
			if err := tt.instr.findChildren(tt.args.imap); (err != nil) != tt.WantErr {
				t.Errorf("instruction.findChildren() error = %v, WantErr %v", err, tt.WantErr)
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
		WantErr bool
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
			if err := tt.instr.populateChildren(tt.args.imap); (err != nil) != tt.WantErr {
				t.Errorf("instruction.populateChildren() error = %v, WantErr %v", err, tt.WantErr)
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
		WantErr bool
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
				WantErr bool
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
			if err := tt.instr.performInstruction(tt.args.mem); (err != nil) != tt.WantErr {
				t.Errorf("instruction.performInstruction() error = %v, WantErr %v", err, tt.WantErr)
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
		WantErr bool
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
			if err := tt.instr.performAll(tt.args.mem); (err != nil) != tt.WantErr {
				t.Errorf("instruction.performAll() error = %v, WantErr %v", err, tt.WantErr)
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
		WantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.imap.lookup(tt.args.id)
			if (err != nil) != tt.WantErr {
				t.Errorf("instructionMap.lookup() error = %v, WantErr %v", err, tt.WantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("instructionMap.lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDay07(t *testing.T) {
	assert := assert.New(t)
	testCases := []tools.TestCase{
		{Details: "Y2016D07 my input",
			Input: "NOT dq -> dr\n" +
				"kg OR kf -> kh\n" +
				"ep OR eo -> eq\n" +
				"44430 -> b\n" +
				"NOT gs -> gt\n" +
				"dd OR do -> dp\n" +
				"eg AND ei -> ej\n" +
				"y AND ae -> ag\n" +
				"jx AND jz -> ka\n" +
				"lf RSHIFT 2 -> lg\n" +
				"z AND aa -> ac\n" +
				"dy AND ej -> el\n" +
				"bj OR bi -> bk\n" +
				"kk RSHIFT 3 -> km\n" +
				"NOT cn -> co\n" +
				"gn AND gp -> gq\n" +
				"cq AND cs -> ct\n" +
				"eo LSHIFT 15 -> es\n" +
				"lg OR lm -> ln\n" +
				"dy OR ej -> ek\n" +
				"NOT di -> dj\n" +
				"1 AND fi -> fj\n" +
				"kf LSHIFT 15 -> kj\n" +
				"NOT jy -> jz\n" +
				"NOT ft -> fu\n" +
				"fs AND fu -> fv\n" +
				"NOT hr -> hs\n" +
				"ck OR cl -> cm\n" +
				"jp RSHIFT 5 -> js\n" +
				"iv OR jb -> jc\n" +
				"is OR it -> iu\n" +
				"ld OR le -> lf\n" +
				"NOT fc -> fd\n" +
				"NOT dm -> dn\n" +
				"bn OR by -> bz\n" +
				"aj AND al -> am\n" +
				"cd LSHIFT 15 -> ch\n" +
				"jp AND ka -> kc\n" +
				"ci OR ct -> cu\n" +
				"gv AND gx -> gy\n" +
				"de AND dk -> dm\n" +
				"x RSHIFT 5 -> aa\n" +
				"et RSHIFT 2 -> eu\n" +
				"x RSHIFT 1 -> aq\n" +
				"ia OR ig -> ih\n" +
				"bk LSHIFT 1 -> ce\n" +
				"y OR ae -> af\n" +
				"NOT ca -> cb\n" +
				"e AND f -> h\n" +
				"ia AND ig -> ii\n" +
				"ck AND cl -> cn\n" +
				"NOT jh -> ji\n" +
				"z OR aa -> ab\n" +
				"1 AND en -> eo\n" +
				"ib AND ic -> ie\n" +
				"NOT eh -> ei\n" +
				"iy AND ja -> jb\n" +
				"NOT bb -> bc\n" +
				"ha OR gz -> hb\n" +
				"1 AND cx -> cy\n" +
				"NOT ax -> ay\n" +
				"ev OR ew -> ex\n" +
				"bn RSHIFT 2 -> bo\n" +
				"er OR es -> et\n" +
				"eu OR fa -> fb\n" +
				"jp OR ka -> kb\n" +
				"ea AND eb -> ed\n" +
				"k AND m -> n\n" +
				"et RSHIFT 3 -> ev\n" +
				"et RSHIFT 5 -> ew\n" +
				"hz RSHIFT 1 -> is\n" +
				"ki OR kj -> kk\n" +
				"NOT h -> i\n" +
				"lv LSHIFT 15 -> lz\n" +
				"as RSHIFT 1 -> bl\n" +
				"hu LSHIFT 15 -> hy\n" +
				"iw AND ix -> iz\n" +
				"lf RSHIFT 1 -> ly\n" +
				"fp OR fv -> fw\n" +
				"1 AND am -> an\n" +
				"ap LSHIFT 1 -> bj\n" +
				"u LSHIFT 1 -> ao\n" +
				"b RSHIFT 5 -> f\n" +
				"jq AND jw -> jy\n" +
				"iu RSHIFT 3 -> iw\n" +
				"ih AND ij -> ik\n" +
				"NOT iz -> ja\n" +
				"de OR dk -> dl\n" +
				"iu OR jf -> jg\n" +
				"as AND bd -> bf\n" +
				"b RSHIFT 3 -> e\n" +
				"jq OR jw -> jx\n" +
				"iv AND jb -> jd\n" +
				"cg OR ch -> ci\n" +
				"iu AND jf -> jh\n" +
				"lx -> a\n" +
				"1 AND cc -> cd\n" +
				"ly OR lz -> ma\n" +
				"NOT el -> em\n" +
				"1 AND bh -> bi\n" +
				"fb AND fd -> fe\n" +
				"lf OR lq -> lr\n" +
				"bn RSHIFT 3 -> bp\n" +
				"bn AND by -> ca\n" +
				"af AND ah -> ai\n" +
				"cf LSHIFT 1 -> cz\n" +
				"dw OR dx -> dy\n" +
				"gj AND gu -> gw\n" +
				"jg AND ji -> jj\n" +
				"jr OR js -> jt\n" +
				"bl OR bm -> bn\n" +
				"gj RSHIFT 2 -> gk\n" +
				"cj OR cp -> cq\n" +
				"gj OR gu -> gv\n" +
				"b OR n -> o\n" +
				"o AND q -> r\n" +
				"bi LSHIFT 15 -> bm\n" +
				"dy RSHIFT 1 -> er\n" +
				"cu AND cw -> cx\n" +
				"iw OR ix -> iy\n" +
				"hc OR hd -> he\n" +
				"0 -> c\n" +
				"db OR dc -> dd\n" +
				"kk RSHIFT 2 -> kl\n" +
				"eq LSHIFT 1 -> fk\n" +
				"dz OR ef -> eg\n" +
				"NOT ed -> ee\n" +
				"lw OR lv -> lx\n" +
				"fw AND fy -> fz\n" +
				"dz AND ef -> eh\n" +
				"jp RSHIFT 3 -> jr\n" +
				"lg AND lm -> lo\n" +
				"ci RSHIFT 2 -> cj\n" +
				"be AND bg -> bh\n" +
				"lc LSHIFT 1 -> lw\n" +
				"hm AND ho -> hp\n" +
				"jr AND js -> ju\n" +
				"1 AND io -> ip\n" +
				"cm AND co -> cp\n" +
				"ib OR ic -> id\n" +
				"NOT bf -> bg\n" +
				"fo RSHIFT 5 -> fr\n" +
				"ip LSHIFT 15 -> it\n" +
				"jt AND jv -> jw\n" +
				"jc AND je -> jf\n" +
				"du OR dt -> dv\n" +
				"NOT fx -> fy\n" +
				"aw AND ay -> az\n" +
				"ge LSHIFT 15 -> gi\n" +
				"NOT ak -> al\n" +
				"fm OR fn -> fo\n" +
				"ff AND fh -> fi\n" +
				"ci RSHIFT 5 -> cl\n" +
				"cz OR cy -> da\n" +
				"NOT ey -> ez\n" +
				"NOT ju -> jv\n" +
				"NOT ls -> lt\n" +
				"kk AND kv -> kx\n" +
				"NOT ii -> ij\n" +
				"kl AND kr -> kt\n" +
				"jk LSHIFT 15 -> jo\n" +
				"e OR f -> g\n" +
				"NOT bs -> bt\n" +
				"hi AND hk -> hl\n" +
				"hz OR ik -> il\n" +
				"ek AND em -> en\n" +
				"ao OR an -> ap\n" +
				"dv LSHIFT 1 -> ep\n" +
				"an LSHIFT 15 -> ar\n" +
				"fo RSHIFT 1 -> gh\n" +
				"NOT im -> in\n" +
				"kk RSHIFT 1 -> ld\n" +
				"hw LSHIFT 1 -> iq\n" +
				"ec AND ee -> ef\n" +
				"hb LSHIFT 1 -> hv\n" +
				"kb AND kd -> ke\n" +
				"x AND ai -> ak\n" +
				"dd AND do -> dq\n" +
				"aq OR ar -> as\n" +
				"iq OR ip -> ir\n" +
				"dl AND dn -> do\n" +
				"iu RSHIFT 5 -> ix\n" +
				"as OR bd -> be\n" +
				"NOT go -> gp\n" +
				"fk OR fj -> fl\n" +
				"jm LSHIFT 1 -> kg\n" +
				"NOT cv -> cw\n" +
				"dp AND dr -> ds\n" +
				"dt LSHIFT 15 -> dx\n" +
				"et RSHIFT 1 -> fm\n" +
				"dy RSHIFT 3 -> ea\n" +
				"fp AND fv -> fx\n" +
				"NOT p -> q\n" +
				"dd RSHIFT 2 -> de\n" +
				"eu AND fa -> fc\n" +
				"ba AND bc -> bd\n" +
				"dh AND dj -> dk\n" +
				"lr AND lt -> lu\n" +
				"he RSHIFT 1 -> hx\n" +
				"ex AND ez -> fa\n" +
				"df OR dg -> dh\n" +
				"fj LSHIFT 15 -> fn\n" +
				"NOT kx -> ky\n" +
				"gk OR gq -> gr\n" +
				"dy RSHIFT 2 -> dz\n" +
				"gh OR gi -> gj\n" +
				"lj AND ll -> lm\n" +
				"x OR ai -> aj\n" +
				"bz AND cb -> cc\n" +
				"1 AND lu -> lv\n" +
				"as RSHIFT 3 -> au\n" +
				"ce OR cd -> cf\n" +
				"il AND in -> io\n" +
				"dd RSHIFT 1 -> dw\n" +
				"NOT lo -> lp\n" +
				"c LSHIFT 1 -> t\n" +
				"dd RSHIFT 3 -> df\n" +
				"dd RSHIFT 5 -> dg\n" +
				"lh AND li -> lk\n" +
				"lf RSHIFT 5 -> li\n" +
				"dy RSHIFT 5 -> eb\n" +
				"NOT kt -> ku\n" +
				"at OR az -> ba\n" +
				"x RSHIFT 3 -> z\n" +
				"NOT lk -> ll\n" +
				"lb OR la -> lc\n" +
				"1 AND r -> s\n" +
				"lh OR li -> lj\n" +
				"ln AND lp -> lq\n" +
				"kk RSHIFT 5 -> kn\n" +
				"ea OR eb -> ec\n" +
				"ci AND ct -> cv\n" +
				"b RSHIFT 2 -> d\n" +
				"jp RSHIFT 1 -> ki\n" +
				"NOT cr -> cs\n" +
				"NOT jd -> je\n" +
				"jp RSHIFT 2 -> jq\n" +
				"jn OR jo -> jp\n" +
				"lf RSHIFT 3 -> lh\n" +
				"1 AND ds -> dt\n" +
				"lf AND lq -> ls\n" +
				"la LSHIFT 15 -> le\n" +
				"NOT fg -> fh\n" +
				"at AND az -> bb\n" +
				"au AND av -> ax\n" +
				"kw AND ky -> kz\n" +
				"v OR w -> x\n" +
				"kk OR kv -> kw\n" +
				"ks AND ku -> kv\n" +
				"kh LSHIFT 1 -> lb\n" +
				"1 AND kz -> la\n" +
				"NOT kc -> kd\n" +
				"x RSHIFT 2 -> y\n" +
				"et OR fe -> ff\n" +
				"et AND fe -> fg\n" +
				"NOT ac -> ad\n" +
				"jl OR jk -> jm\n" +
				"1 AND jj -> jk\n" +
				"bn RSHIFT 1 -> cg\n" +
				"NOT kp -> kq\n" +
				"ci RSHIFT 3 -> ck\n" +
				"ev AND ew -> ey\n" +
				"1 AND ke -> kf\n" +
				"cj AND cp -> cr\n" +
				"ir LSHIFT 1 -> jl\n" +
				"NOT gw -> gx\n" +
				"as RSHIFT 2 -> at\n" +
				"iu RSHIFT 1 -> jn\n" +
				"cy LSHIFT 15 -> dc\n" +
				"hg OR hh -> hi\n" +
				"ci RSHIFT 1 -> db\n" +
				"au OR av -> aw\n" +
				"km AND kn -> kp\n" +
				"gj RSHIFT 1 -> hc\n" +
				"iu RSHIFT 2 -> iv\n" +
				"ab AND ad -> ae\n" +
				"da LSHIFT 1 -> du\n" +
				"NOT bw -> bx\n" +
				"km OR kn -> ko\n" +
				"ko AND kq -> kr\n" +
				"bv AND bx -> by\n" +
				"kl OR kr -> ks\n" +
				"1 AND ht -> hu\n" +
				"df AND dg -> di\n" +
				"NOT ag -> ah\n" +
				"d OR j -> k\n" +
				"d AND j -> l\n" +
				"b AND n -> p\n" +
				"gf OR ge -> gg\n" +
				"gg LSHIFT 1 -> ha\n" +
				"bn RSHIFT 5 -> bq\n" +
				"bo OR bu -> bv\n" +
				"1 AND gy -> gz\n" +
				"s LSHIFT 15 -> w\n" +
				"NOT ie -> if\n" +
				"as RSHIFT 5 -> av\n" +
				"bo AND bu -> bw\n" +
				"hz AND ik -> im\n" +
				"bp AND bq -> bs\n" +
				"b RSHIFT 1 -> v\n" +
				"NOT l -> m\n" +
				"bp OR bq -> br\n" +
				"g AND i -> j\n" +
				"br AND bt -> bu\n" +
				"t OR s -> u\n" +
				"hz RSHIFT 5 -> ic\n" +
				"gk AND gq -> gs\n" +
				"fl LSHIFT 1 -> gf\n" +
				"he RSHIFT 3 -> hg\n" +
				"gz LSHIFT 15 -> hd\n" +
				"hf OR hl -> hm\n" +
				"1 AND gd -> ge\n" +
				"fo OR fz -> ga\n" +
				"id AND if -> ig\n" +
				"fo AND fz -> gb\n" +
				"gr AND gt -> gu\n" +
				"he OR hp -> hq\n" +
				"fq AND fr -> ft\n" +
				"ga AND gc -> gd\n" +
				"fo RSHIFT 2 -> fp\n" +
				"gl OR gm -> gn\n" +
				"hg AND hh -> hj\n" +
				"NOT hn -> ho\n" +
				"gl AND gm -> go\n" +
				"he RSHIFT 5 -> hh\n" +
				"NOT gb -> gc\n" +
				"hq AND hs -> ht\n" +
				"hz RSHIFT 3 -> ib\n" +
				"hz RSHIFT 2 -> ia\n" +
				"fq OR fr -> fs\n" +
				"hx OR hy -> hz\n" +
				"he AND hp -> hr\n" +
				"gj RSHIFT 5 -> gm\n" +
				"hf AND hl -> hn\n" +
				"hv OR hu -> hw\n" +
				"NOT hj -> hk\n" +
				"gj RSHIFT 3 -> gl\n" +
				"fo RSHIFT 3 -> fq\n" +
				"he RSHIFT 2 -> hf\n",
			Result1: "3176",
			Result2: "14710"},
	}
	for _, tt := range testCases {
		tt.Test(Day07, assert)
	}
}
