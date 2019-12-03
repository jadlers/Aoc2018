package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/jadlers/advent-of-code/util"
)

type direction int

const (
	Up direction = iota
	Left
	Down
	Right
)

func (d direction) String() string {
	return [...]string{"Up", "Left", "Down", "Right"}[d]
}

type point struct {
	x int
	y int
}

func (p1 point) distanceTo(p2 point) int {
	dx := util.AbsInt(p1.x - p2.x)
	dy := util.AbsInt(p1.y - p2.y)
	return dx + dy
}

type move struct {
	dir direction
	len int
}

func main() {
	paths := readInput()

	fmt.Printf("Part 1: %v\n", p1(paths)) // 500 and 502 too low
	// fmt.Printf("Part 2: %v\n", p2(originalInput))
}

// Part 1: 1211
func p1(paths [][]move) int {
	// const size = 600
	const size = 20000
	area := make([][]bool, size)
	for col := 0; col < size; col++ {
		area[col] = make([]bool, size)
	}
	startPoint := point{size / 2, size / 2}
	// area[startPoint.x][startPoint.y] = true

	markPath(area, paths[0], startPoint)
	intersections := findIntersections(area, paths[1], startPoint)

	minDist := size
	// fmt.Printf("Starting point: %v\n", startPoint)
	for _, in := range intersections {
		if in == startPoint {
			continue
		}
		dist := startPoint.distanceTo(in)
		// fmt.Printf("%v dist to intersection at %v\n", dist, in)
		if dist < minDist {
			minDist = dist
		}
	}

	// Print area
	// for _, col := range area[startPoint.x-(minDist+10) : startPoint.x+(minDist+10)] {
	// 	fmt.Println(col[startPoint.y-(minDist+10) : startPoint.y+(minDist+10)])
	// }
	// for _, col := range area {
	// 	fmt.Println(col)
	// }
	// for _, col := range area[290:540] {
	// 	fmt.Println(col[250:450])
	// }

	return minDist
}

// markPath increases the value in each traversed point and returns the end
// position
func markPath(area [][]bool, path []move, po point) point {
	// fmt.Printf("Starting in %v\n", po)
	for _, p := range path {
		fmt.Printf("Moving %v %v\n", p.dir.String(), p.len)

		switch p.dir {
		case Up:
			for d := 0; d < p.len; d++ {
				po.y++
				area[po.x][po.y] = true
			}
		case Left:
			for d := 0; d < p.len; d++ {
				po.x--
				area[po.x][po.y] = true
			}
		case Down:
			for d := 0; d < p.len; d++ {
				po.y--
				area[po.x][po.y] = true
			}
		case Right:
			for d := 0; d < p.len; d++ {
				po.x++
				area[po.x][po.y] = true
			}
		}
		// fmt.Printf("Now in %v\n", po)
	}

	return po
}

func findIntersections(area [][]bool, path []move, po point) []point {
	intersections := []point{}

	for _, p := range path {
		switch p.dir {
		case Up:
			for d := 0; d < p.len; d++ {
				po.y++
				if pointIntersects(area, po) {
					intersections = append(intersections, po)
				}
			}
		case Left:
			for d := 0; d < p.len; d++ {
				po.x--
				if pointIntersects(area, po) {
					intersections = append(intersections, po)
				}
			}
		case Down:
			for d := 0; d < p.len; d++ {
				po.y--
				if pointIntersects(area, po) {
					intersections = append(intersections, po)
				}
			}
		case Right:
			for d := 0; d < p.len; d++ {
				po.x++
				if pointIntersects(area, po) {
					intersections = append(intersections, po)
				}
			}
		}
		// fmt.Printf("Now in %v\n", po)
	}

	return intersections
}

func pointIntersects(area [][]bool, po point) bool {
	fmt.Printf("Cheching %v (%v)\n", po, area[po.x][po.y])
	if area[po.x][po.y] { // Intersection found
		return true
	}
	return false
}

// Part 2: 2505

func readInput() [][]move {
	lines := util.ReadLines()
	paths := make([][]move, len(lines))
	for i, line := range lines {
		strVals := strings.Split(line, ",")
		for _, val := range strVals {
			var m move
			length, _ := strconv.Atoi(val[1:])
			m.len = length

			switch val[0] {
			case 'U':
				m.dir = Up
			case 'L':
				m.dir = Left
			case 'D':
				m.dir = Down
			case 'R':
				m.dir = Right
			default:
				panic("No valid direction")
			}
			paths[i] = append(paths[i], m)
		}
	}

	return paths
}
