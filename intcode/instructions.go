package intcode

func (p *Program) executeAdd() {
	p.loadRegister(&p.ra, 1)
	p.loadRegister(&p.rb, 2)

	p.ra += p.rb

	p.storeRegister(&p.ra, 3)
}

func (p *Program) executeMul() {
	p.loadRegister(&p.ra, 1)
	p.loadRegister(&p.rb, 2)

	p.ra *= p.rb

	p.storeRegister(&p.ra, 3)
}

func (p *Program) executeInp() {
	p.ra = p.iBus
	p.storeRegister(&p.ra, 1)
}

func (p *Program) executeOut() {
	p.loadRegister(&p.ra, 1)
	p.oBus = p.ra
}

func (p *Program) execJumpFalse() {
	p.loadRegister(&p.ra, 1)
	p.loadRegister(&p.rb, 2)

	if p.ra == 0 {
		p.ip = p.rb
		return
	}

	p.ip += 3
}

func (p *Program) execJumpTrue() {
	p.loadRegister(&p.ra, 1)
	p.loadRegister(&p.rb, 2)

	if p.ra > 0 {
		p.ip = p.rb
		return
	}

	p.ip += 3
}

func (p *Program) execLessThan() {
	p.loadRegister(&p.ra, 1)
	p.loadRegister(&p.rb, 2)

	if p.ra < p.rb {
		p.ra = 1
	} else {
		p.ra = 0
	}

	p.storeRegister(&p.ra, 3)
}

func (p *Program) execEqualTo() {
	p.loadRegister(&p.ra, 1)
	p.loadRegister(&p.rb, 2)

	if p.ra == p.rb {
		p.ra = 1
	} else {
		p.ra = 0
	}

	p.storeRegister(&p.ra, 3)
}
