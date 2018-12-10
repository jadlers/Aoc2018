package main

import (
	"container/ring"
	"fmt"
)

func main() {
	const Players = 465
	const LastMarbleValue = 71498
	p1, p2 := Day9(Players, LastMarbleValue)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day9(numPlayers, lastMarbleValue int) (p1, p2 int64) {
	extendedMarbleValue := int64(lastMarbleValue * 100)
	playerScores := map[int]int64{}
	curPlayer := 1
	circle := ring.New(1)
	circle.Value = int64(0) // Add the first marble

	for curVal := int64(1); curVal <= extendedMarbleValue; curVal++ {
		if curVal%23 != 0 {
			newElement := ring.New(1)
			newElement.Value = curVal
			// circle will be one after the newElement, aka one after currentValue
			circle = circle.Link(newElement)
		} else { // Special case where marble value is divisible by 23
			playerScores[curPlayer] += curVal
			circle = circle.Move(-9) // Need to get one before remove and are already at next
			removed := circle.Unlink(1)
			playerScores[curPlayer] += removed.Value.(int64)
			circle = circle.Move(2) // TODO: Correct?
		}

		if curVal == int64(lastMarbleValue) {
			p1 = maxScore(playerScores)
		}

		curPlayer++
		if curPlayer%(numPlayers+1) == 0 {
			curPlayer = 1
		}
	}

	p2 = maxScore(playerScores)

	return
}

func maxScore(scores map[int]int64) int64 {
	max := int64(0)
	for _, val := range scores {
		if val > max {
			max = val
		}
	}
	return max
}
