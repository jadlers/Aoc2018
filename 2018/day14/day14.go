package main

import (
	"container/ring"
	"fmt"
)

func main() {
	const P1noRecipes = 260321
	p1, p2 := Day14(P1noRecipes)
	p2 = Day14P2([]int{2, 6, 0, 3, 2, 1})

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func Day14(P1noRecipes int) (p1 []int, p2 int) {
	circle := ring.New(2)
	first := circle // Reference to the first value
	fst := first    // The current val for the first elf
	circle.Value = 3
	circle = circle.Next()
	circle.Value = 7
	snd := circle // The current val for the second elf
	end := circle // Will always point to the end since that's where we're adding

	len := 2
	for len < 10+P1noRecipes {
		newRecipes := fst.Value.(int) + snd.Value.(int)
		if newRecipes < 10 { // add val to end
			newRing := ring.New(1)
			newRing.Value = newRecipes
			end = end.Link(newRing)
			len++
		} else { // split the val
			newRing := ring.New(2)
			newRing.Value = 1
			newRing = newRing.Next()
			newRing.Value = newRecipes - 10
			newRing = newRing.Next()
			end = end.Link(newRing)
			len += 2
		}
		end = end.Move(-1)

		// Advance fst & snd
		fst = fst.Move(1 + fst.Value.(int))
		snd = snd.Move(1 + snd.Value.(int))
	}

	circle = circle.Move(P1noRecipes - 1) // Move to the end
	p1 = getP1Answer(circle)

	return
}

func getP1Answer(numbers *ring.Ring) []int {
	res := make([]int, 10)
	for i := range res {
		res[i] = numbers.Value.(int)
		numbers = numbers.Next()
	}
	return res
}

func Day14P2(sequence []int) int {
	circle := ring.New(2)
	first := circle // Reference to the first value
	fst := first    // The current val for the first elf
	circle.Value = 3
	circle = circle.Next()
	circle.Value = 7
	snd := circle // The current val for the second elf
	end := circle // Will always point to the end since that's where we're adding

	length := 2
	for {
		newRecipes := fst.Value.(int) + snd.Value.(int)
		if newRecipes < 10 { // add val to end
			newRing := ring.New(1)
			newRing.Value = newRecipes
			end = end.Link(newRing)
			length++
		} else { // split the val
			newRing := ring.New(2)
			newRing.Value = 1
			newRing = newRing.Next()
			newRing.Value = newRecipes - 10
			newRing = newRing.Next()
			end = end.Link(newRing)
			length += 2
		}
		end = end.Move(-1)

		// printAll(first)

		if found, res := sequenceFound(sequence, end); found {
			return res
		}

		// Advance fst & snd
		fst = fst.Move(1 + fst.Value.(int))
		snd = snd.Move(1 + snd.Value.(int))
	}

	// Guessed: 150718217 (too high), 20319118 (too high), last -1 is correct

	return 0
}

func sequenceFound(seq []int, end *ring.Ring) (bool, int) {
	end = end.Move(1 - len(seq))
	found := true
	for i := 0; i < len(seq); i++ {
		if seq[i] != end.Value {
			found = false
		}
		end = end.Next()
	}

	if found {
		return true, end.Len() - len(seq)
	}
	end = end.Move(-1 - len(seq))
	for i := 0; i < len(seq); i++ {
		if seq[i] != end.Value {
			return false, -1
		}
		end = end.Next()
	}

	return true, end.Len() - len(seq) - 1
}

func printAll(recipes *ring.Ring) {
	fmt.Print("[")
	recipes.Do(func(v interface{}) {
		fmt.Print(v, ", ")
	})
	fmt.Println("]")
}
