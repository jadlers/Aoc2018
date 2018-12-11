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

func Day11(serialNumber int) (p1 []int, p2 int) {
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
			if squareVal == 30 {
				fmt.Printf("(%v, %v), maxval: %v\n", x, y, maxVal)
			}
			// fmt.Println(squareVal)
		}
	}



	// For example 1
	// if err := print(grid, 32, 44, 5, 5); err != nil {
	// if err := print(grid, 165, 102, 5, 5); err != nil {
	// 	fmt.Println(err)
	// }
	// For example 2
	// if err := print(grid, 20, 60, 5, 5); err != nil {
	// 	fmt.Println(err)
	// }

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
