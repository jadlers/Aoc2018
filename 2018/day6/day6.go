package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"strconv"
	"strings"
)

func main() {
	const SAFE_DISTANCE = 10000
	input := util.ReadLines()
	p1, p2 := Day6(input, SAFE_DISTANCE)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

type point struct {
	x, y int
	name string
}

func Day6(lines []string, safeDistance int) (p1, p2 int) {
	points, width, height := getFieldPoints(lines)

	// field[x][y]["name"]
	field := initField(width, height)
	for _, p := range points {
		for x := 0; x <= width; x++ {
			for y := 0; y <= height; y++ {
				dx := util.AbsInt(x - p.x)
				dy := util.AbsInt(y - p.y)
				field[x][y][p.name] = dx + dy
			}
		}
	}

	assignedField := assignField(field, len(field), len(field[0]))

	areaSizes := map[string]int{}
	for _, row := range assignedField {
		for _, column := range row {
			name := strings.ToLower(column)
			areaSizes[name] += 1
		}
	}

	infiniteAreas := findInfiniteAreas(assignedField)

	for key, size := range areaSizes {
		if !infiniteAreas[key] {
			if size > p1 {
				p1 = size
			}
		}
	}

	// Go through all places in the field and get total distance all points
	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			safePoint := distanceToAllPointsLessThan(safeDistance, x, y, points)
			if safePoint {
				p2++
			}
		}
	}

	return
}

func distanceToAllPointsLessThan(lt, x, y int, points []point) bool {
	totalDistance := 0
	for _, p := range points {
		dx := util.AbsInt(x - p.x)
		dy := util.AbsInt(y - p.y)
		totalDistance += (dx + dy)

		if totalDistance >= lt {
			return false
		}
	}

	return true
}

func findInfiniteAreas(field [][]string) map[string]bool {
	set := map[string]bool{}
	width := len(field)
	height := len(field[0])
	// Top & Bottom
	for y := 0; y < height; y += (height - 1) {
		for x := 0; x < width; x++ {
			name := strings.ToLower(field[x][y])
			set[name] = true
		}
	}

	// Left & Right
	for x := 0; x < width; x += (width - 1) {
		for y := 0; y < height; y++ {
			name := strings.ToLower(field[x][y])
			set[name] = true
		}
	}

	return set
}

func assignField(field [][]map[string]int, width, height int) [][]string {
	final := make([][]string, width)
	for x := 0; x < width; x++ {
		final[x] = make([]string, height)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			curr := field[x][y]
			closestPoint := getClosestPoint(curr)
			final[x][y] = closestPoint
		}
	}

	return final
}

func getClosestPoint(m map[string]int) string {
	first := true
	var minDistance int
	var closestNames []string
	for name, distance := range m {
		if first {
			minDistance = distance
			closestNames = []string{name}
			first = false
		} else if distance == minDistance {
			closestNames = append(closestNames, name)
		} else if distance < minDistance {
			minDistance = distance
			closestNames = []string{name}
		}
	}

	if minDistance == 0 {
		return strings.ToUpper(closestNames[0])
	} else if len(closestNames) > 1 {
		return "."
	}

	return closestNames[0]
}

func initField(w, h int) [][]map[string]int {
	field := [][]map[string]int{}
	for i := 0; i <= w; i++ {
		// field[i] = []map[string]int{}
		field = append(field, []map[string]int{})
		for j := 0; j <= h; j++ {
			field[i] = append(field[i], map[string]int{})
		}
	}

	return field
}

func getFieldPoints(lines []string) (points []point, maxX, maxY int) {
	// names := strings.Split("abcdefghijklmnopqrstuvxyz", "")
	for i, line := range lines {
		// Get point
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		newPoint := point{x, y, strconv.Itoa(i)}
		points = append(points, newPoint)
		// Check size of field
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	return
}
