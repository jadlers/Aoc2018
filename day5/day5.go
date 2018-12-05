package main

import (
	"fmt"
	util "github.com/jadlers/AoC2018/util"
	"strings"
)

func main() {
	input := util.ReadLines()
	p1, p2 := Day5(input[0])

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day5(input string) (p1, p2 int) {
	var res string = input
	var removed int
	for {
		res, removed = removeReacting(res)
		if removed == 0 {
			break
		}
	}

	p1 = len(res)

	p2 = 0
	return
}

func removeReacting(polymer string) (remainder string, removed int) {
	chars := strings.Split(polymer, "")
	for i, char := range chars {
		if i+1 == len(polymer) {
			break
		}

		if IsReacting(char, chars[i+1]) {
			chars[i] = " "
			chars[i+1] = " "
			removed += 2
		}
	}
	remainder = strings.Join(chars, "")
	remainder = util.TrimAllWhitespace(remainder)

	return
}

func IsReacting(a, b string) bool {
	if strings.ToUpper(a) == strings.ToUpper(b) {
		if a != b {
			return true
		}
	}
	return false
}
