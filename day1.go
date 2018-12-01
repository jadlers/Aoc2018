package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var inputNumbers []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		next := scanner.Text()
		inputNumbers = append(inputNumbers, next)
	}

	fmt.Println(Day1(inputNumbers))
}

func Day1(input []string) int {
	var total int = 0
	var numbers []int
	reached := map[int]bool{0: true} // Initial frequency

	// Read input
	for _, val := range input {
		intVal, _ := strconv.Atoi(val)
		numbers = append(numbers, intVal)
	}

	for {
		for _, val := range numbers {
			total += val
			if reached[total] {
				return total
			} else {
				reached[total] = true
			}

		}
	}

}
