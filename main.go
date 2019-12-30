package main

import (
	"fmt"
	"strings"
)

func main() {
	src := "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,19,9,23,1,23,13,27,1,10,27,31,2,31,13,35,1,10,35,39,2,9,39,43,2,43,9,47,1,6,47,51,1,10,51,55,2,55,13,59,1,59,10,63,2,63,13,67,2,67,9,71,1,6,71,75,2,75,9,79,1,79,5,83,2,83,13,87,1,9,87,91,1,13,91,95,1,2,95,99,1,99,6,0,99,2,14,0,0"

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			p := newIntcodeProgram(strings.NewReader(src))
			p.set(1, i)
			p.set(2, j)
			p.run()

			if p.mem[0] == 19690720 {
				fmt.Println(p.out())
				break
			}
		}
	}
}
