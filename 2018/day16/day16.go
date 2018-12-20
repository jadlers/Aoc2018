package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
)

func main() {
	p1Input := util.ReadLinesFromFile("input_p1.txt")
	p1 := Day16_p1(p1Input)

	fmt.Println("Part 1:", p1)
	// fmt.Println("Part 2:", p2)
}

type Dump struct {
	before    [4]int
	operation []int
	after     [4]int
}

func Day16_p1(input []string) (p1 int) {
	dumps := []Dump{}
	for i := 0; i < len(input); i += 4 {
		var ba, bb, bc, bd int
		fmt.Sscanf(input[i], "Before: [%d, %d, %d, %d]", &ba, &bb, &bc, &bd)
		bef := [4]int{ba, bb, bc, bd}

		var oa, ob, oc, od int
		fmt.Sscanf(input[i+1], "%d %d %d %d", &oa, &ob, &oc, &od)
		ops := []int{oa, ob, oc, od}

		var aa, ab, ac, ad int
		fmt.Sscanf(input[i+2], "After: [%d, %d, %d, %d]", &aa, &ab, &ac, &ad)
		aft := [4]int{aa, ab, ac, ad}

		dumps = append(dumps, Dump{bef, ops, aft})
	}

	for _, dump := range dumps {
		equalOps := NumEqualOps(dump)
		// fmt.Println(equalOps)
		if len(equalOps) >= 3 {
			p1++
		}
	}

	// Guesses: 407 (too low)

	return
}

func NumEqualOps(dump Dump) (res []int) {
	operations := []instrFunc{
		addi, addr,
		muli, mulr,
		bani, banr,
		bori, borr,
		seti, setr,
		gtir, gtri, gtrr,
		eqir, eqri, eqrr}

	for i, op := range operations {
		exec := op(dump.before, dump.operation)
		// fmt.Printf("Performing operation %v with %v on %v results in %v (%v)\n", i, dump.operation, dump.before, exec, dump.after)
		if exec == dump.after {
			// fmt.Println("equal: ", i)
			res = append(res, i)
		}
	}

	return
}

type instrFunc func([4]int, []int) [4]int

func addi(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] + arg[2]
	return registers
}

func addr(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] + registers[arg[2]]
	return registers
}

func muli(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] * arg[2]
	return registers
}

func mulr(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] * registers[arg[2]]
	return registers
}

func bani(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] & arg[2]
	return registers
}

func banr(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] & registers[arg[2]]
	return registers
}

func bori(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] | arg[2]
	return registers
}

func borr(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]] | registers[arg[2]]
	return registers
}

func seti(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = arg[1]
	return registers
}

func setr(registers [4]int, arg []int) [4]int {
	registers[arg[3]] = registers[arg[1]]
	return registers
}

func gtir(registers [4]int, arg []int) [4]int {
	if arg[1] > registers[arg[2]] {
		registers[arg[3]] = 1
	} else {
		registers[arg[3]] = 0
	}
	return registers
}

func gtri(registers [4]int, arg []int) [4]int {
	if registers[arg[1]] > arg[2] {
		registers[arg[3]] = 1
	} else {
		registers[arg[3]] = 0
	}
	return registers
}

func gtrr(registers [4]int, arg []int) [4]int {
	if registers[arg[1]] > registers[arg[2]] {
		registers[arg[3]] = 1
	} else {
		registers[arg[3]] = 0
	}
	return registers
}

func eqir(registers [4]int, arg []int) [4]int {
	if arg[1] == registers[arg[2]] {
		registers[arg[3]] = 1
	} else {
		registers[arg[3]] = 0
	}
	return registers
}

func eqri(registers [4]int, arg []int) [4]int {
	if registers[arg[1]] == arg[2] {
		registers[arg[3]] = 1
	} else {
		registers[arg[3]] = 0
	}
	return registers
}

func eqrr(registers [4]int, arg []int) [4]int {
	if registers[arg[1]] == registers[arg[2]] {
		registers[arg[3]] = 1
	} else {
		registers[arg[3]] = 0
	}
	return registers
}
