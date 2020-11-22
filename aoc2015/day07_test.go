package aoc2015

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/janreggie/aoc/internal"
	"github.com/stretchr/testify/assert"
)

const day07myInput = `NOT dq -> dr
kg OR kf -> kh
ep OR eo -> eq
44430 -> b
NOT gs -> gt
dd OR do -> dp
eg AND ei -> ej
y AND ae -> ag
jx AND jz -> ka
lf RSHIFT 2 -> lg
z AND aa -> ac
dy AND ej -> el
bj OR bi -> bk
kk RSHIFT 3 -> km
NOT cn -> co
gn AND gp -> gq
cq AND cs -> ct
eo LSHIFT 15 -> es
lg OR lm -> ln
dy OR ej -> ek
NOT di -> dj
1 AND fi -> fj
kf LSHIFT 15 -> kj
NOT jy -> jz
NOT ft -> fu
fs AND fu -> fv
NOT hr -> hs
ck OR cl -> cm
jp RSHIFT 5 -> js
iv OR jb -> jc
is OR it -> iu
ld OR le -> lf
NOT fc -> fd
NOT dm -> dn
bn OR by -> bz
aj AND al -> am
cd LSHIFT 15 -> ch
jp AND ka -> kc
ci OR ct -> cu
gv AND gx -> gy
de AND dk -> dm
x RSHIFT 5 -> aa
et RSHIFT 2 -> eu
x RSHIFT 1 -> aq
ia OR ig -> ih
bk LSHIFT 1 -> ce
y OR ae -> af
NOT ca -> cb
e AND f -> h
ia AND ig -> ii
ck AND cl -> cn
NOT jh -> ji
z OR aa -> ab
1 AND en -> eo
ib AND ic -> ie
NOT eh -> ei
iy AND ja -> jb
NOT bb -> bc
ha OR gz -> hb
1 AND cx -> cy
NOT ax -> ay
ev OR ew -> ex
bn RSHIFT 2 -> bo
er OR es -> et
eu OR fa -> fb
jp OR ka -> kb
ea AND eb -> ed
k AND m -> n
et RSHIFT 3 -> ev
et RSHIFT 5 -> ew
hz RSHIFT 1 -> is
ki OR kj -> kk
NOT h -> i
lv LSHIFT 15 -> lz
as RSHIFT 1 -> bl
hu LSHIFT 15 -> hy
iw AND ix -> iz
lf RSHIFT 1 -> ly
fp OR fv -> fw
1 AND am -> an
ap LSHIFT 1 -> bj
u LSHIFT 1 -> ao
b RSHIFT 5 -> f
jq AND jw -> jy
iu RSHIFT 3 -> iw
ih AND ij -> ik
NOT iz -> ja
de OR dk -> dl
iu OR jf -> jg
as AND bd -> bf
b RSHIFT 3 -> e
jq OR jw -> jx
iv AND jb -> jd
cg OR ch -> ci
iu AND jf -> jh
lx -> a
1 AND cc -> cd
ly OR lz -> ma
NOT el -> em
1 AND bh -> bi
fb AND fd -> fe
lf OR lq -> lr
bn RSHIFT 3 -> bp
bn AND by -> ca
af AND ah -> ai
cf LSHIFT 1 -> cz
dw OR dx -> dy
gj AND gu -> gw
jg AND ji -> jj
jr OR js -> jt
bl OR bm -> bn
gj RSHIFT 2 -> gk
cj OR cp -> cq
gj OR gu -> gv
b OR n -> o
o AND q -> r
bi LSHIFT 15 -> bm
dy RSHIFT 1 -> er
cu AND cw -> cx
iw OR ix -> iy
hc OR hd -> he
0 -> c
db OR dc -> dd
kk RSHIFT 2 -> kl
eq LSHIFT 1 -> fk
dz OR ef -> eg
NOT ed -> ee
lw OR lv -> lx
fw AND fy -> fz
dz AND ef -> eh
jp RSHIFT 3 -> jr
lg AND lm -> lo
ci RSHIFT 2 -> cj
be AND bg -> bh
lc LSHIFT 1 -> lw
hm AND ho -> hp
jr AND js -> ju
1 AND io -> ip
cm AND co -> cp
ib OR ic -> id
NOT bf -> bg
fo RSHIFT 5 -> fr
ip LSHIFT 15 -> it
jt AND jv -> jw
jc AND je -> jf
du OR dt -> dv
NOT fx -> fy
aw AND ay -> az
ge LSHIFT 15 -> gi
NOT ak -> al
fm OR fn -> fo
ff AND fh -> fi
ci RSHIFT 5 -> cl
cz OR cy -> da
NOT ey -> ez
NOT ju -> jv
NOT ls -> lt
kk AND kv -> kx
NOT ii -> ij
kl AND kr -> kt
jk LSHIFT 15 -> jo
e OR f -> g
NOT bs -> bt
hi AND hk -> hl
hz OR ik -> il
ek AND em -> en
ao OR an -> ap
dv LSHIFT 1 -> ep
an LSHIFT 15 -> ar
fo RSHIFT 1 -> gh
NOT im -> in
kk RSHIFT 1 -> ld
hw LSHIFT 1 -> iq
ec AND ee -> ef
hb LSHIFT 1 -> hv
kb AND kd -> ke
x AND ai -> ak
dd AND do -> dq
aq OR ar -> as
iq OR ip -> ir
dl AND dn -> do
iu RSHIFT 5 -> ix
as OR bd -> be
NOT go -> gp
fk OR fj -> fl
jm LSHIFT 1 -> kg
NOT cv -> cw
dp AND dr -> ds
dt LSHIFT 15 -> dx
et RSHIFT 1 -> fm
dy RSHIFT 3 -> ea
fp AND fv -> fx
NOT p -> q
dd RSHIFT 2 -> de
eu AND fa -> fc
ba AND bc -> bd
dh AND dj -> dk
lr AND lt -> lu
he RSHIFT 1 -> hx
ex AND ez -> fa
df OR dg -> dh
fj LSHIFT 15 -> fn
NOT kx -> ky
gk OR gq -> gr
dy RSHIFT 2 -> dz
gh OR gi -> gj
lj AND ll -> lm
x OR ai -> aj
bz AND cb -> cc
1 AND lu -> lv
as RSHIFT 3 -> au
ce OR cd -> cf
il AND in -> io
dd RSHIFT 1 -> dw
NOT lo -> lp
c LSHIFT 1 -> t
dd RSHIFT 3 -> df
dd RSHIFT 5 -> dg
lh AND li -> lk
lf RSHIFT 5 -> li
dy RSHIFT 5 -> eb
NOT kt -> ku
at OR az -> ba
x RSHIFT 3 -> z
NOT lk -> ll
lb OR la -> lc
1 AND r -> s
lh OR li -> lj
ln AND lp -> lq
kk RSHIFT 5 -> kn
ea OR eb -> ec
ci AND ct -> cv
b RSHIFT 2 -> d
jp RSHIFT 1 -> ki
NOT cr -> cs
NOT jd -> je
jp RSHIFT 2 -> jq
jn OR jo -> jp
lf RSHIFT 3 -> lh
1 AND ds -> dt
lf AND lq -> ls
la LSHIFT 15 -> le
NOT fg -> fh
at AND az -> bb
au AND av -> ax
kw AND ky -> kz
v OR w -> x
kk OR kv -> kw
ks AND ku -> kv
kh LSHIFT 1 -> lb
1 AND kz -> la
NOT kc -> kd
x RSHIFT 2 -> y
et OR fe -> ff
et AND fe -> fg
NOT ac -> ad
jl OR jk -> jm
1 AND jj -> jk
bn RSHIFT 1 -> cg
NOT kp -> kq
ci RSHIFT 3 -> ck
ev AND ew -> ey
1 AND ke -> kf
cj AND cp -> cr
ir LSHIFT 1 -> jl
NOT gw -> gx
as RSHIFT 2 -> at
iu RSHIFT 1 -> jn
cy LSHIFT 15 -> dc
hg OR hh -> hi
ci RSHIFT 1 -> db
au OR av -> aw
km AND kn -> kp
gj RSHIFT 1 -> hc
iu RSHIFT 2 -> iv
ab AND ad -> ae
da LSHIFT 1 -> du
NOT bw -> bx
km OR kn -> ko
ko AND kq -> kr
bv AND bx -> by
kl OR kr -> ks
1 AND ht -> hu
df AND dg -> di
NOT ag -> ah
d OR j -> k
d AND j -> l
b AND n -> p
gf OR ge -> gg
gg LSHIFT 1 -> ha
bn RSHIFT 5 -> bq
bo OR bu -> bv
1 AND gy -> gz
s LSHIFT 15 -> w
NOT ie -> if
as RSHIFT 5 -> av
bo AND bu -> bw
hz AND ik -> im
bp AND bq -> bs
b RSHIFT 1 -> v
NOT l -> m
bp OR bq -> br
g AND i -> j
br AND bt -> bu
t OR s -> u
hz RSHIFT 5 -> ic
gk AND gq -> gs
fl LSHIFT 1 -> gf
he RSHIFT 3 -> hg
gz LSHIFT 15 -> hd
hf OR hl -> hm
1 AND gd -> ge
fo OR fz -> ga
id AND if -> ig
fo AND fz -> gb
gr AND gt -> gu
he OR hp -> hq
fq AND fr -> ft
ga AND gc -> gd
fo RSHIFT 2 -> fp
gl OR gm -> gn
hg AND hh -> hj
NOT hn -> ho
gl AND gm -> go
he RSHIFT 5 -> hh
NOT gb -> gc
hq AND hs -> ht
hz RSHIFT 3 -> ib
hz RSHIFT 2 -> ia
fq OR fr -> fs
hx OR hy -> hz
he AND hp -> hr
gj RSHIFT 5 -> gm
hf AND hl -> hn
hv OR hu -> hw
NOT hj -> hk
gj RSHIFT 3 -> gl
fo RSHIFT 3 -> fq
he RSHIFT 2 -> hf
`

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
