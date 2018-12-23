package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
)

var operations []instrFunc = []instrFunc{
	addi, addr,
	muli, mulr,
	bani, banr,
	bori, borr,
	seti, setr,
	gtir, gtri, gtrr,
	eqir, eqri, eqrr}

func main() {
	p1Input := util.ReadLinesFromFile("input_p1.txt")
	p1 := Day16_p1(p1Input)

	p2Input := util.ReadLinesFromFile("input_p2.txt")
	p2 := Day16_p2(p1Input, p2Input)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
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
		equalOps := EqualOps(dump)
		// fmt.Println(equalOps)
		if len(equalOps) >= 3 {
			p1++
		}
	}

	// Guesses: 407 (too low)

	return
}

func Day16_p2(dumpInput, program []string) (p2 int) {
	dumps := []Dump{}
	opCodeDumps := map[int][]Dump{}
	for i := 0; i < len(dumpInput); i += 4 {
		var ba, bb, bc, bd int
		fmt.Sscanf(dumpInput[i], "Before: [%d, %d, %d, %d]", &ba, &bb, &bc, &bd)
		bef := [4]int{ba, bb, bc, bd}

		var oa, ob, oc, od int
		fmt.Sscanf(dumpInput[i+1], "%d %d %d %d", &oa, &ob, &oc, &od)
		ops := []int{oa, ob, oc, od}

		var aa, ab, ac, ad int
		fmt.Sscanf(dumpInput[i+2], "After: [%d, %d, %d, %d]", &aa, &ab, &ac, &ad)
		aft := [4]int{aa, ab, ac, ad}

		dumps = append(dumps, Dump{bef, ops, aft})
		opCodeDumps[oa] = append(opCodeDumps[oa], Dump{bef, ops, aft})
	}

	opToAllFunc := map[int][][]int{}
	for _, dump := range dumps {
		curOpCode := dump.operation[0]
		opToAllFunc[curOpCode] = append(opToAllFunc[curOpCode], EqualOps(dump))
	}

	// Remove all possible instructions which does not pass every dump
	opToFunc := map[int]map[int]bool{}
	for opCode, allPossibleFuncs := range opToAllFunc {
		newList := map[int]bool{}
		for _, possibleFuncs := range allPossibleFuncs {
			if len(newList) == 0 {
				for _, funcs := range possibleFuncs {
					newList[funcs] = true
				}
			} else {
				for val := range newList {
					if !includes(possibleFuncs, val) {
						delete(newList, val)
					}
				}
			}
		}
		opToFunc[opCode] = newList
	}

	for opCode, possibleOps := range opToFunc {
		fmt.Printf("opToFunc[%v]: %v\n", opCode, possibleOps)
	}

	// key is opCode in dump data and val is index in operations slice
	// determined := map[int]int{}

	// determinedPrevLen := -1
	// for determinedPrevLen != len(determined) {
	// 	determinedPrevLen = len(determined)
	// 	fmt.Println("len(opToFunc):", len(opToFunc))
	// 	fmt.Println("determined:", determined)
	// 	for opCode, possibleOps := range opToFunc {
	// 		fmt.Printf("opToFunc[%v]: %v\n", opCode, possibleOps)
	// 		if len(possibleOps) == 1 {
	// 			for possibleOp := range possibleOps {
	// 				determined[opCode] = possibleOp
	// 				delete(opToFunc, opCode)
	// 			}
	// 		}
	// 	}

	// 	// Remove determined from all possible ops
	// 	for opDetermined := range determined {
	// 		for _, possibleOps := range opToFunc {
	// 			delete(possibleOps, opDetermined)
	// 		}
	// 	}
	// }

	// fmt.Println("determined:", determined)
	// for opCode, possibleOps := range opToFunc {
	// 	fmt.Printf("opToFunc[%v]: %v\n", opCode, possibleOps)
	// }

	// Did this manually
	determined := map[int]int{
		0:  2,
		1:  7,
		2:  11,
		3:  14,
		4:  12,
		5:  13,
		6:  0,
		7:  9,
		8:  3,
		9:  1,
		10: 6,
		11: 4,
		12: 8,
		13: 15,
		14: 5,
		15: 10,
	}

	instructions := make([][]int, len(program))
	for i, instr := range program {
		var a, b, c, d int
		fmt.Sscanf(instr, "%d %d %d %d", &a, &b, &c, &d)
		instructions[i] = []int{a, b, c, d}
	}

	registers := [4]int{0, 0, 0, 0}
	for i, instruction := range instructions {
		registers = operations[determined[instruction[0]]](registers, instruction)
		fmt.Printf("%3d: registers: %v\n", i, registers)
	}

	fmt.Println(determined)
	fmt.Println("intstructions:", instructions[:10])
	p2 = registers[0]

	return
}

func RemoveOccurancesOf(list []int, val int) []int {
	newList := []int{}
	for _, item := range list {
		if item != val {
			newList = append(newList, item)
		}
	}
	return newList
}

func includes(list []int, val int) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}

func EqualOps(dump Dump) (res []int) {

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
