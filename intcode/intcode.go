package intcode

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

// Program is a program.
type Program struct {
	mem        []int   // bytecode array
	inst       [4]int  // instruction array
	errors     []error // error list
	ip         int     // instruction pointer
	mode       string  // position/immediate mode for a, b, c
	ra, rb     int     // registers
	iBus, oBus int     // input and output BUS
}

const (
	instAdd  = 1  // Add
	instMul  = 2  // Multiply
	instInp  = 3  // Input
	instOut  = 4  // Output
	instJmpT = 5  // Jump if true
	instJmpF = 6  // Jump if false
	instLess = 7  // Less than
	instEq   = 8  // Equal to
	instHalt = 99 // Halt
)

// NewProgram takes a reader that contains the source code of the program.
func NewProgram(r io.Reader) *Program {
	codes := make([]int, 0)
	errs := make([]error, 0)

	src, err := ioutil.ReadAll(r)
	if err != nil {
		errs = append(errs, err)
	}

	for _, code := range strings.Split(string(src), ",") {
		x, err := strconv.ParseInt(code, 10, 64)
		if err != nil {
			errs = append(errs, err)
		}
		codes = append(codes, int(x))
	}

	return &Program{mem: codes, errors: errs, ip: 0}
}

func (p *Program) fetch() {
	copy(p.inst[:], p.mem[p.ip:])
}

func (p *Program) decode() {
	bytecode := fmt.Sprintf("%05d", p.inst[0])
	p.mode = bytecode[0:3]
	inst, _ := strconv.ParseInt(bytecode[3:], 10, 64)
	p.inst[0] = int(inst)
}

func (p *Program) execute() {
	switch p.inst[0] {
	case instAdd:
		p.executeAdd()
		p.ip += 4
	case instMul:
		p.executeMul()
		p.ip += 4
	case instInp:
		p.executeInp()
		p.ip += 2
	case instOut:
		p.executeOut()
		p.ip += 2
	case instJmpF:
		p.execJumpFalse()
	case instJmpT:
		p.execJumpTrue()
	case instLess:
		p.execLessThan()
		p.ip += 4
	case instEq:
		p.execEqualTo()
		p.ip += 4
	case instHalt:
		p.ip = len(p.mem)
	default:
		err := fmt.Errorf("%d: unknown opcode: %d", p.ip, p.inst[0])
		p.errors = append(p.errors, err)
		p.ip++
	}
}

// Run executes the program in memory
func (p *Program) Run() {
	for p.ip < len(p.mem) {
		p.fetch()
		p.decode()
		p.execute()
	}
}

// Set will set the value at location to the value in x
func (p *Program) Set(loc, x int) {
	p.mem[loc] = x
}

// Input takes the passed value and stores it in register A
func (p *Program) Input(s string) {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		p.errors = append(p.errors, err)
		return
	}

	p.iBus = int(x)
}

// Output takes the value in register A and returns it as a string
func (p *Program) Output() string {
	return strconv.FormatInt(int64(p.oBus), 10)
}

// MemoryDump will return the entire outpuut of memory.
func (p *Program) MemoryDump() string {
	tmp := make([]string, len(p.mem))
	for i, code := range p.mem {
		tmp[i] = strconv.FormatInt(int64(code), 10)
	}
	return strings.Join(tmp, ",")
}

// Errors returns a list of all errors.
func (p *Program) Errors() []error {
	return p.errors
}
