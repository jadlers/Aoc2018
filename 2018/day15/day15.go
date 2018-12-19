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
	winnerFound := false
	for !winnerFound {
		units, winnerFound = oneTick(walls, units)
		// printCave(walls, units)
		if winnerFound {
			p1 = tick * remainingHp(units)
		} else {
			tick++
		}
	}
	// printCave(walls, units)
	// fmt.Printf("ticks: %v, remainingHp: %v\n", tick, remainingHp(units))
	// fmt.Println()

	// // Guesses: 182207 (too high), 179220 (too high), 176233 (too low)

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

func oneTick(walls [][]bool, units map[Position]Unit) (map[Position]Unit, bool) {
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

	visited := map[Position]bool{}
	// Keep track of the ones visited, since other units may move onto a position
	// where there previously was an unit that died. That unit will then get to
	// play twic

	for _, curPosition := range order {
		// printCave(walls, units)
		if _, alreadyVisited := visited[curPosition]; alreadyVisited {
			// Noop
			// fmt.Println("alreadyVisited:", curPosition)
		} else if curUnit, exists := units[curPosition]; exists {
			if foundTarget, goTowards := findClosestOpponent(curPosition, walls, units); foundTarget {
				if !curPosition.inAttackRange(goTowards) { // Move towards being in attack range
					// fmt.Printf("at %v will move towards: %v\n", curPosition, goTowards)
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
						if adjacentUnit, exists := units[adjacentPos]; exists && curUnit.isOpponent(adjacentUnit) {
							if adjacentUnit.hp < target.hp {
								target = adjacentUnit
								targetPos = adjacentPos
							}
						}
					}

					target.hp -= curUnit.damage
					if target.hp <= 0 {
						delete(units, targetPos)
						visited[targetPos] = true
					} else {
						units[targetPos] = target
					}
				}
			} else if winner(units) { // Unit can't move, or all opponents are dead
				// fmt.Printf("No more units for %v at %v to attack, victory\n", curUnit, curPosition)
				return units, true
			}
		}
		visited[curPosition] = true
	}
	victory := false
	return units, victory
}

// Finds the position of the closest opponent
func findClosestOpponent(start Position, walls [][]bool, units map[Position]Unit) (foundTarget bool, goTowards Position) {
	queued := map[Position]int{} // Map of all positions covered, value is distance
	searchingUnit := units[start]
	q := list.New()
	q.PushBack(start)
	queued[start] = 0
	targets := []Position{}

	// fmt.Printf("Unit at %v looking for target\n", start)
	firstTargetDistance := -1
	directions := []string{"up", "left", "right", "down"} // traversing order
	for cur := q.Front(); cur != nil; cur = cur.Next() {
		curPos := cur.Value.(Position)
		// fmt.Printf("Currently at (%v, %v)\n", curPos.r, curPos.c)

		// Just look through up until the first distance found
		if firstTargetDistance < 0 || firstTargetDistance == queued[curPos]+1 {
			for _, dir := range directions {
				adjacentPos := curPos.move(dir)
				if _, adjacentAlreadyQueued := queued[adjacentPos]; !adjacentAlreadyQueued {
					queued[adjacentPos] = queued[curPos] + 1
					// Check if any pos in range is opponent
					if unit, exists := units[adjacentPos]; exists {
						// fmt.Printf("Going %v to %v hits unit %v\n", dir, adjacentPos, unit)
						if unit.isOpponent(searchingUnit) {
							firstTargetDistance = queued[curPos] + 1 // Distance to target found
							targets = append(targets, adjacentPos)
							// fmt.Println("Going to:", adjacentPos)
							// return
						}
						// Check if the position can be walked to
					} else if !adjacentAlreadyQueued && !walls[adjacentPos.r][adjacentPos.c] {
						// fmt.Printf("Going %v to empty location %v\n", dir, adjacentPos)
						q.PushBack(adjacentPos)
					}
				}
			}
		}
	}

	if len(targets) > 0 {
		foundTarget = true

		sort.SliceStable(targets, func(i, j int) bool {
			if targets[i].r != targets[j].r {
				return targets[i].r < targets[j].r
			}
			return targets[i].c < targets[j].c
		})
		goTowards = targets[0]
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
