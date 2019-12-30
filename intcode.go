package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type intcodeProgram struct {
	mem    []int   // bytecode array
	inst   [4]int  // instruction array
	errors []error // error list
	ip     int     // instruction pointer
	ra, rb int     // registers
}

const (
	instAdd  = 1  // Add
	instMul  = 2  // Multiply
	instHalt = 99 // Halt
)

func newIntcodeProgram(r io.Reader) *intcodeProgram {
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

	return &intcodeProgram{mem: codes, errors: errs, ip: 0}
}

func (p *intcodeProgram) fetch() {
	copy(p.inst[:], p.mem[p.ip:])
}

func (p *intcodeProgram) decode() {
	// no-op for now
}

func (p *intcodeProgram) execute() {
	switch p.inst[0] {
	case instAdd:
		p.executeAdd()
		p.ip += 4
	case instMul:
		p.executeMul()
		p.ip += 4
	case instHalt:
		p.ip = len(p.mem)
	default:
		err := fmt.Errorf("%d: unknown opcode: %d", p.ip, p.inst[0])
		p.errors = append(p.errors, err)
		p.ip++
	}
}

func (p *intcodeProgram) run() {
	for p.ip < len(p.mem) {
		p.fetch()
		p.decode()
		p.execute()
	}
}

func (p *intcodeProgram) set(loc, x int) {
	p.mem[loc] = x
}

func (p *intcodeProgram) executeAdd() {
	p.ra = p.mem[p.inst[1]]
	p.rb = p.mem[p.inst[2]]

	p.ra += p.rb

	p.mem[p.inst[3]] = p.ra
}

func (p *intcodeProgram) executeMul() {
	p.ra = p.mem[p.inst[1]]
	p.rb = p.mem[p.inst[2]]

	p.ra *= p.rb

	p.mem[p.inst[3]] = p.ra
}

func (p *intcodeProgram) out() string {
	tmp := make([]string, len(p.mem))
	for i, code := range p.mem {
		tmp[i] = strconv.FormatInt(int64(code), 10)
	}
	return strings.Join(tmp, ",")
}
