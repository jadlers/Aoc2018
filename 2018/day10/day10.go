package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	// "strings"
)

func main() {
	lines := util.ReadLines()
	p1, p2 := Day10(lines)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

type Point struct {
	x, y   int // Position
	dx, dy int // Velocity
}

func Day10(lines []string) (done bool, p2 int) {
	points := []Point{}
	for _, line := range lines {
		var x, y, dx, dy int
		fmt.Sscanf(line, "position=<%d, %d> velocity=<%d, %d>", &x, &y, &dx, &dy)
		points = append(points, Point{x, y, dx, dy})
	}

	for loop := 0; loop < 100000000; loop++ {
		// for {
		for idx, point := range points {
			// fmt.Println("Before update: ", point)
			point.x += point.dx
			point.y += point.dy
			points[idx] = point
			// fmt.Println("After update: ", point)
		}
		// minX, minY, maxX, maxY := minMaxVals(points)
		_, minY, _, maxY := minMaxVals(points)
		// xDiff := maxX - minX
		yDiff := maxY - minY
		// fmt.Printf("xDiff: %.2v,\t yDiff: %.2v\n", xDiff, yDiff)
		// if xDiff < 100 || yDiff < 100 {
		// 	fmt.Printf("xDiff: %.2v,\t yDiff: %.2v\n", xDiff, yDiff)
		// }
		// if yDiff < 170 {
		if loop == 10458 {
			fmt.Println("NEXT!!!")
			fmt.Println(loop)
			printPoints(points)
			done = true
		}
		if done && yDiff > 170 {
			break
		}
		// fmt.Printf("xDiff: %.2v,\t yDiff: %.2v\n", maxX-minX, maxY-minY)

		// printPoints(points)
	}

	// fmt.Println(points)

	return
}

func largest(fst, snd int) int {
	if fst > snd {
		return fst
	}
	return snd
}

func printPoints(points []Point) {
	minX, minY, maxX, maxY := minMaxVals(points)
	// _, _, maxX, maxY := minMaxVals(points)

	view := make([][]string, (largest(maxX, minX)*2)+2)
	for x, _ := range view {
		view[x] = make([]string, (largest(maxY, minY)*2)+2)
		for y, _ := range view[x] {
			// view[x][y] = fmt.Sprintf("(%.2v, %.2v)", x, y)
			view[x][y] = " "
		}
	}

	// view[x][y]
	xZero := len(view) / 2
	yZero := len(view[0]) / 2
	// fmt.Printf("xZero: %v, yZero: %v\n", xZero, yZero)

	// Origo is at
	// view[xZero][yZero] = "x"

	// fmt.Printf("len(x): %v, len(y): %v\n", len(view[0]), len(view))
	for _, point := range points {
		calX := (xZero + point.x)
		calY := (yZero + point.y)
		// calX := (point.x)
		// calY := (point.y)

		if calX >= len(view) || calY >= len(view[0]) {
			fmt.Printf("point: %v, xZero: %v, yZero: %v\n", point, xZero, yZero)
			fmt.Printf("calX: %v, calY: %v, len(x): %v, len(y): %v\n", calX, calY, len(view), len(view[0]))
		} else if calX < 0 || calY < 0 {
			fmt.Println("Less than zero!", point)
		} else {
			view[calX][calY] = "#"
		}
	}

	for _, row := range view {
		// if strings.Contains(strings.Join(row, ""), "#") {
		// 	fmt.Println(row)
		// }
		fmt.Println(row)
	}

}

func minMaxVals(points []Point) (minX, minY, maxX, maxY int) {
	first := points[0]
	minX, maxX = first.x, first.x
	minY, maxY = first.y, first.y
	for _, point := range points {
		if point.x < minX {
			minX = point.x
		} else if point.x > maxX {
			maxX = point.x
		}

		if point.y < minY {
			minY = point.y
		} else if point.y > maxY {
			maxY = point.y
		}
	}

	return
}
