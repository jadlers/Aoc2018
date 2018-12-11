package main

import (
	"fmt"
	"strconv"
)

func main() {
	const GridSerialNumber = 6042
	p1, p2 := Day11(GridSerialNumber)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day11(serialNumber int) (p1, p2 []int) {
	grid := make([][]int, 300)
	for i := range grid {
		grid[i] = make([]int, 300)
	}

	for x, row := range grid {
		for y := range row {
			powerLevel := CalcPowerLevel(x, y, serialNumber)
			grid[x][y] = powerLevel
		}
	}

	p1 = []int{0, 0}
	maxVal := 0
	for x := 0; x < 297; x++ {
		for y := 0; y < 297; y++ {
			squareVal := Calc3x3(grid, x, y)
			if squareVal > maxVal {
				p1[0] = x
				p1[1] = y
				maxVal = squareVal
			}
		}
	}

	maxVal = 0
	p2 = []int{0, 0, 0}
	for s := 4; s < 300; s++ {
		for x := 0; x < 300-s; x++ {
			for y := 0; y < 300-s; y++ {
				squareVal := CalcSquare(grid, x, y, s)
				if squareVal > maxVal {
					p2[0] = x
					p2[1] = y
					p2[2] = s
					maxVal = squareVal
				}
			}
		}
		fmt.Println(s)
	}

	return
}

func CalcSquare(grid [][]int, x, y, size int) (val int) {
	if x+size >= 300 || y+size >= 300 {
		return 0
	}

	for i := x; i < x+size; i++ {
		for j := y; j < y+size; j++ {
			val += grid[i][j]
		}
	}
	return

}

func Calc3x3(grid [][]int, x, y int) (val int) {
	if x+3 >= 300 || y+3 >= 300 {
		return 0
	}
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			val += grid[i][j]
		}
	}
	return
}

func CalcPowerLevel(x, y, serialNumber int) (powerLevel int) {
	rackId := x + 10
	powerLevel = rackId
	powerLevel *= y
	powerLevel += serialNumber
	powerLevel *= rackId

	hundreds := strconv.Itoa(powerLevel)
	if len(hundreds) < 3 {
		return 0
	}
	hundredVal, _ := strconv.Atoi(string(hundreds[len(hundreds)-3]))
	powerLevel = hundredVal - 5
	return
}

func print(grid [][]int, x, y, width, height int) error {
	maxX, maxY := x+width-1, y+height-1
	if maxX >= 300 {
		return fmt.Errorf("Out of bounds, max(x) = %v\n", maxX)
	} else if x < 0 {
		return fmt.Errorf("Out of bounds, start(x) = %v\n", x-1)
	} else if maxY >= 300 {
		return fmt.Errorf("Out of bounds, max(y) = %v\n", maxY)
	} else if y < 0 {
		return fmt.Errorf("Out of bounds, start(y) = %v\n", y-1)
	}

	for j := y; j <= maxY; j++ {
		for i := x; i <= maxX; i++ {
			// fmt.Printf("(%v, %v)\n", i, j)
			fmt.Printf("%3v\t", grid[i][j])
		}
		fmt.Println()
	}

	return nil
}
