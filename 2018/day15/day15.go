package main

import (
	"container/list"
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"sort"
)

func main() {
	lines := util.ReadLines()
	p1, p2 := Day15(lines)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

type Position struct {
	r int
	c int
}

type Unit struct {
	hp     int
	damage int
	isElf  bool
}

func Day15(lines []string) (p1, p2 int) {
	walls := make([][]bool, len(lines))
	units := map[Position]Unit{}

	for r, row := range lines {
		walls[r] = make([]bool, len(row))
		for c, char := range row {
			if char == '#' {
				walls[r][c] = true
			} else if char == '.' {
				walls[r][c] = false
			} else if char == 'E' {
				units[Position{r, c}] = Unit{hp: 200, damage: 3, isElf: true}
			} else if char == 'G' {
				units[Position{r, c}] = Unit{hp: 200, damage: 3, isElf: false}
			}
		}
	}

	fmt.Println("Initial state")
	printCave(walls, units)
	tick := 0
	for !winner(units) {
		units = oneTick(walls, units)
		// fmt.Printf("After %v rounds\n", tick)
		// printCave(walls, units)
		// fmt.Println()
		tick++
	}
	fmt.Println("Ticks:", tick)
	printCave(walls, units)
	p1 = tick * remainingHp(units)
	// // Guesses: 182207 (too high), 179220 (too high)

	return
}

func winner(units map[Position]Unit) bool {
	first := true
	var lastWasElf bool
	for _, unit := range units {
		if first {
			lastWasElf = unit.isElf
			first = false
		} else if lastWasElf != unit.isElf {
			return false
		}
	}
	return true
}

func remainingHp(units map[Position]Unit) (total int) {
	for _, unit := range units {
		total += unit.hp
	}
	return
}

func oneTick(walls [][]bool, units map[Position]Unit) map[Position]Unit {
	// Sort in traversing order
	order := []Position{}
	for pos := range units {
		order = append(order, pos)
	}

	sort.SliceStable(order, func(i, j int) bool {
		if order[i].r != order[j].r {
			return order[i].r < order[j].r
		}
		return order[i].c < order[j].c
	})

	for _, curPosition := range order {
		// printCave(walls, units)
		if curUnit, exists := units[curPosition]; exists {
			if foundTarget, goTowards := findClosestOpponent(curPosition, walls, units); foundTarget {
				// findClosestOpponent works
				if !curPosition.inAttackRange(goTowards) { // Move towards being in attack range
					// fmt.Printf("at %v will move towards: %v\n", curPosition, goTowards)
					// fmt.Printf("%v can't attack %v\n", curPosition, goTowards)
					nextPosition := getMovePosition(curPosition, goTowards, walls, units)
					delete(units, curPosition)
					curPosition = nextPosition
					units[curPosition] = curUnit
				}

				if curPosition.inAttackRange(goTowards) { // If after move or already in attack range
					// Find the target with the lowest hp left
					target := Unit{hp: 999, damage: 0, isElf: !curUnit.isElf}
					targetPos := Position{-1, -1}
					directions := []string{"up", "left", "right", "down"}
					for _, dir := range directions {
						adjacentPos := curPosition.move(dir)
						// fmt.Printf("adjacentPos: %v, with unit %v\n", adjacentPos, units[adjacentPos])
						if adjacentUnit, exists := units[adjacentPos]; exists && curUnit.isOpponent(adjacentUnit) {
							// fmt.Printf("adjacentPos: %v, with unit %v\n", adjacentPos, units[adjacentPos])
							if adjacentUnit.hp < target.hp {
								target = adjacentUnit
								targetPos = adjacentPos
							}
						}
					}

					// fmt.Printf("%v attacking %v at %v\n", curPosition, target, targetPos)
					target.hp -= curUnit.damage
					if target.hp <= 0 {
						delete(units, targetPos)
					} else {
						units[targetPos] = target
					}
				}
			} else { // Unit can't move
				units[curPosition] = curUnit
			}
		}
	}
	return units
}

// Finds the position of the closest opponent
func findClosestOpponent(start Position, walls [][]bool, units map[Position]Unit) (foundTarget bool, goTowards Position) {
	queued := map[Position]bool{} // Map of all positions covered
	searchingUnit := units[start]
	q := list.New()
	q.PushBack(start)
	queued[start] = true

	// fmt.Printf("Unit at %v looking for target\n", start)
	directions := []string{"up", "left", "right", "down"} // traversing order
	for cur := q.Front(); cur != nil; cur = cur.Next() {
		curPos := cur.Value.(Position)
		// fmt.Printf("Currently at (%v, %v)\n", curPos.r, curPos.c)

		for _, dir := range directions {
			adjacentPos := curPos.move(dir)
			// fmt.Printf("Going %v to %v\n", dir, adjacentPos)
			// Check if any pos in range is opponent
			if unit, exists := units[adjacentPos]; exists {
				if unit.isOpponent(searchingUnit) {
					foundTarget = true
					goTowards = adjacentPos
					// fmt.Println("Going to:", adjacentPos)
					return
				}
				// Check if the position can be walked to
			} else if !queued[adjacentPos] && !walls[adjacentPos.r][adjacentPos.c] {
				q.PushBack(adjacentPos)
				queued[adjacentPos] = true
			}
		}
	}

	return
}

// Finds the first position the unit at start should move to in order to get to
// target.
// Would be better to implement with a depth first search instead
func getMovePosition(from, to Position, walls [][]bool, units map[Position]Unit) Position {
	distance := map[Position]int{} // Map of all positions covered
	q := list.New()
	q.PushBack(to)
	distance[to] = 0

	directions := []string{"up", "left", "right", "down"}
	for cur := q.Front(); cur != nil; cur = cur.Next() {
		curPos := cur.Value.(Position)
		// fmt.Printf("Currently at (%v, %v)\n", curPos.r, curPos.c)

		// If there is a distance for the from we must have found a closest path
		if _, exists := distance[from]; exists {
			cur = q.Back() // Break loop when the "end" is found
		}

		for _, dir := range directions {
			adjacentPos := curPos.move(dir)
			_, alreadyChecked := distance[adjacentPos]
			_, hasUnit := units[adjacentPos]

			if !alreadyChecked && !hasUnit && !walls[adjacentPos.r][adjacentPos.c] {
				q.PushBack(adjacentPos)
				distance[adjacentPos] = distance[curPos] + 1
			} else if adjacentPos == from {
				distance[from] = distance[curPos] + 1
			}
		}
	}

	if _, exists := distance[from]; exists {
		min := Position{-1, -1}
		distance[min] = 9999999999 // arbitrary large value, better way?
		for _, dir := range directions {
			dist, ex := distance[from.move(dir)]
			if ex && dist < distance[min] {
				min = from.move(dir)
			}
		}
		return min
	}

	return Position{-1, -1} // Should never ever happen
}

func (p *Position) move(direction string) Position {
	switch direction {
	case "up":
		return p.up()
	case "left":
		return p.left()
	case "right":
		return p.right()
	case "down":
		return p.down()
	}
	fmt.Errorf("Invalid direction %v\n", direction)
	return Position{-1, -1}
}

func (p *Position) up() Position {
	return Position{p.r - 1, p.c}
}

func (p *Position) down() Position {
	return Position{p.r + 1, p.c}
}

func (p *Position) left() Position {
	return Position{p.r, p.c - 1}
}

func (p *Position) right() Position {
	return Position{p.r, p.c + 1}
}

func (u *Unit) isOpponent(rhs Unit) bool {
	if u.isElf != rhs.isElf {
		return true
	}
	return false
}

func (p *Position) inAttackRange(target Position) bool {
	dr := util.AbsInt(p.r - target.r)
	dc := util.AbsInt(p.c - target.c)
	if dr == 1 && dc == 0 {
		return true
	} else if dr == 0 && dc == 1 {
		return true
	}
	return false
}

func printUnit(u Unit) {
	fmt.Printf("Elf: %v, hp: %v, damage: %v\n", u.isElf, u.hp, u.damage)
}

func printCave(walls [][]bool, units map[Position]Unit) {
	rowString := "\t"
	for r, row := range walls {
		for c := range row {
			if val, exists := units[Position{r, c}]; exists {
				if val.isElf {
					fmt.Print("E")
					rowString += fmt.Sprintf("E(%v), ", val.hp)
				} else {
					fmt.Print("G")
					rowString += fmt.Sprintf("G(%v), ", val.hp)
				}
			} else if walls[r][c] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println(rowString)
		rowString = "\t"
	}
}
