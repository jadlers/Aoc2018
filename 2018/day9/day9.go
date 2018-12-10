package main

import (
	"fmt"
)

func main() {
	const Players = 465
	const LastMarbleValue = 71498
	p1, p2 := Day9(Players, LastMarbleValue)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day9(numPlayers, lastMarbleValue int) (p1, p2 int) {
	playerScores := map[int]int{}
	curPlayer := 1
	circle := []int{}
	circle = append(circle, 0) // Add the first marble
	curMarble := 0

	for curVal := 1; curVal <= lastMarbleValue; curVal++ {
		if curVal%23 != 0 {
			nextMarbleIdx := ((curMarble + 1) % len(circle)) + 1
			if nextMarbleIdx == len(circle) {
				circle = append(circle, curVal)
			} else {
				snd := append([]int{curVal}, circle[nextMarbleIdx:]...)
				circle = append(circle[:nextMarbleIdx], snd...)
			}
			curMarble = nextMarbleIdx
		} else { // Special case where marble value is divisible by 23
			playerScores[curPlayer] += curVal
			removeMarbleIdx := (curMarble + len(circle) - 7) % len(circle)
			playerScores[curPlayer] += circle[removeMarbleIdx]
			circle = append(circle[:removeMarbleIdx], circle[removeMarbleIdx+1:]...)
			curMarble = removeMarbleIdx
		}

		// fmt.Printf("[%v]\t%v, curMarble[%v] = %v\n", curPlayer, circle, curMarble, circle[curMarble])

		curPlayer++
		if curPlayer%(numPlayers+1) == 0 {
			curPlayer = 1
		}
	}

	p1 = maxScore(playerScores)

	return
}

func maxScore(scores map[int]int) int {
	max := 0
	for _, val := range scores {
		if val > max {
			max = val
		}
	}
	return max
}
