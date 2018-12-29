package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
)

func main() {
	input := util.ReadLinesFromFile("input.txt")
	p1, p2 := Day25(input)

	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

type coord struct {
	x, y, z, w int
}

func Day25(input []string) (p1, p2 int) {
	coords := make([]coord, len(input))
	for i, line := range input {
		var x, y, z, w int
		fmt.Sscanf(line, "%d,%d,%d,%d", &x, &y, &z, &w)
		coords[i] = coord{x, y, z, w}
	}
	constilations := map[coord]int{} // Where the int is the number of the constilation
	nextConstilation := 1

	constilations[coords[0]] = 0
	for _, coord := range coords {
		closeConstilations := []int{}
		for constCoord, constNum := range constilations {
			dist := manhattanDistance(coord, constCoord)
			if dist <= 3 {
				closeConstilations = append(closeConstilations, constNum)
			}
		}

		// Check which constilations the coord is close to, if > 1 combine otherwise just add
		if len(closeConstilations) == 0 {
			constilations[coord] = nextConstilation
			nextConstilation++
		} else if len(closeConstilations) == 1 {
			constilations[coord] = closeConstilations[0]
		} else {
			fmt.Println("JOIN")
			constilations[coord] = nextConstilation
			// Change all with numbers in closeConstilations to nextConstilation
			for constCoord, constNum := range constilations {
				if util.IncludesInt(constNum, closeConstilations) {
					constilations[constCoord] = nextConstilation
				}
			}
			// Update next
			nextConstilation++
		}

	}
	fmt.Println(constilations)

	// Count number of constilations
	numConstilations := map[int]bool{}
	for _, num := range constilations {
		numConstilations[num] = true
	}
	p1 = len(numConstilations)

	return
}

func manhattanDistance(c1, c2 coord) (distance int) {
	distance += util.AbsInt(c1.x - c2.x)
	distance += util.AbsInt(c1.y - c2.y)
	distance += util.AbsInt(c1.z - c2.z)
	distance += util.AbsInt(c1.w - c2.w)
	return
}
