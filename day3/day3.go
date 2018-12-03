package main

import (
	"fmt"
	util "github.com/jadlers/AoC2018/util"
)

type claim struct {
	id      int
	offsetX int
	offsetY int
	width   int
	height  int
}

func main() {
	lines := util.ReadLines()

	fmt.Printf("Part 1: %v\n", Part1(lines))
	fmt.Printf("Part 2: %v\n", Part2(lines))
}

func Part1(lines []string) int {
	// Turn input into the form of claims
	claims := createClaimsSlice(lines)
	overlapping := [1000][1000]int{}

	for _, claim := range claims {
		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				overlapping[x+claim.offsetX][y+claim.offsetY]++
				// fmt.Printf("overlapping[%v][%v] = %v\n", x+claim.offsetX, y+claim.offsetY, overlapping[x+claim.offsetX][y+claim.offsetY])
			}
		}
	}

	sum := 0
	for x := 0; x < len(overlapping); x++ {
		for y := 0; y < len(overlapping[0]); y++ {
			if overlapping[x][y] > 1 {
				sum += 1
			}
		}
	}

	return sum
}

func Part2(lines []string) int {
	// Turn input into the form of claims
	claims := createClaimsSlice(lines)
	overlapping := [1000][1000][]int{}
	hasOverlap := make([]bool, len(lines)+1)
	hasOverlap[0] = true // No claim has id 0

	for _, claim := range claims {
		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				cur := &overlapping[x+claim.offsetX][y+claim.offsetY]
				*cur = append(*cur, claim.id)

				if len(*cur) > 1 {
					for _, id := range *cur {
						hasOverlap[id] = true
					}
				}
			}
		}
	}

	for id, overlaps := range hasOverlap {
		if !overlaps {
			return id
		}
	}

	return 0
}

// This solution builds directly on the solution for part 1
// which is probably a quicker way of solving
func Part2_alt(lines []string) int {
	// Turn input into the form of claims
	claims := createClaimsSlice(lines)
	overlapping := [1000][1000][]int{}

	for _, claim := range claims {
		for x := 0; x < claim.width; x++ {
			for y := 0; y < claim.height; y++ {
				cur := &overlapping[x+claim.offsetX][y+claim.offsetY]
				*cur = append(*cur, claim.id)
			}
		}
	}

	for _, claim := range claims {
		solo := true
		for x := 0; x < len(overlapping); x++ {
			for y := 0; y < len(overlapping[0]); y++ {
				if len(overlapping[x][y]) > 1 {
					solo = false
				}
			}
		}
		if solo {
			return claim.id
		}
	}

	return 0
}

func createClaimsSlice(lines []string) []claim {
	var claims []claim
	for _, line := range lines {
		var id, offsetX, offsetY, width, height int
		_, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d\n", &id, &offsetX, &offsetY, &width, &height)
		if err != nil {
			fmt.Println(err)
		}
		claims = append(claims, claim{id, offsetX, offsetY, width, height})
	}
	return claims
}
