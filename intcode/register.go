package intcode

func (p *Program) loadRegister(reg *int, index int) {
	if p.mode[3-index] == '1' {
		*reg = p.inst[index]
		return
	}

	*reg = p.mem[p.inst[index]]
}

func (p *Program) storeRegister(reg *int, index int) {
	if p.mode[3-index] == '1' {
		p.inst[index] = *reg
		return
	}

	p.mem[p.inst[index]] = *reg
}
