package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"sort"
	"strings"
)

func main() {
	const Workers = 5
	const ExtraTimePerStep = 60
	input := util.ReadLines()
	p1, p2 := Day7(input, Workers, ExtraTimePerStep)

	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

type Instruction struct {
	name   string
	before string
}

type Block struct {
	blocking string
	node     string
	active   bool
}

type BlockSlice []Block

func Day7(lines []string, workers, extraTimePerStep int) (p1 string, p2 int) {
	allNames := map[string]bool{}
	blocking := BlockSlice{}
	for _, line := range lines {
		// Step Y must be finished before step D can begin.
		words := strings.Fields(line)
		allNames[words[1]] = true
		allNames[words[7]] = true
		blocking = append(blocking, Block{words[1], words[7], true})
	}

	for {
		chosen := nextStep(allNames, blocking)

		// Find all blocking by chosen
		for i, b := range blocking {
			if chosen == b.blocking {
				blocking[i] = Block{b.blocking, b.node, false}
			}

		}
		delete(allNames, chosen)

		p1 = p1 + chosen // Add to res
		if len(allNames) == 0 {
			break
		}
	}

	// PART 2
	// reset blocks
	for i, b := range blocking {
		blocking[i] = Block{b.blocking, b.node, true}
	}

	// nextAvailableTime[worker_id]time_step_after_finish_current_job
	workerInstruction := map[int]string{}
	nextAvailableTime := map[int]int{}
	for i := 0; i < workers; i++ {
		nextAvailableTime[i] = 0
		workerInstruction[i] = ""
	}

	instructionsDone := ""
	// So that I can use nextStep
	instructionsLeft := map[string]bool{}
	for _, char := range strings.Split(p1, "") {
		instructionsLeft[char] = true
	}
	for step := 0; len(nextAvailableTime) != 0; step++ {
		// fmt.Println()
		// fmt.Println("Step: ", step)
		// fmt.Println("before workerInstruction: ", workerInstruction)
		// fmt.Println("before nextAvailableTime: ", nextAvailableTime)

		availableWorkers := []int{}
		for worker, time := range nextAvailableTime {
			// See if any worker finished this step and mark that instruction as done
			if step == time {
				instructionsDone += workerInstruction[worker]
				delete(instructionsLeft, workerInstruction[worker])
				for i, b := range blocking {
					if workerInstruction[worker] == b.blocking {
						blocking[i] = Block{b.blocking, b.node, false}
					}
				}
				workerInstruction[worker] = ""

			}

			// Find workers for next step
			if step >= time {
				availableWorkers = append(availableWorkers, worker)
			}
		}
		// fmt.Println("Blocking: ", blocking)
		// fmt.Println("availableWorkers: ", availableWorkers)
		// fmt.Println("Remaining: ", instructionsLeft)

		// Assign work
		for _, currWorker := range availableWorkers {
			if len(instructionsLeft) != 0 {
				hasNext, next := getNextStep(instructionsLeft, blocking)
				if len(availableWorkers) > 0 && hasNext {
					// fmt.Println("Next: ", next)
					nextAvailableTime[currWorker] = step + extraTimePerStep + InstructionTime(next[0])
					workerInstruction[currWorker] = next
					delete(instructionsLeft, next)
				}
			} else { // When all instructions are handed out delete workers as they finish
				delete(nextAvailableTime, currWorker)
			}
		}
		// fmt.Println("Instructions done: ", instructionsDone)
		// fmt.Println("after workerInstruction: ", workerInstruction)
		// fmt.Println("after nextAvailableTime: ", nextAvailableTime)

		p2++
	}

	// The check is done after p2 increases
	p2--

	return
}

func InstructionTime(letter uint8) int {
	return int(letter) - 64
}

// Strikingly similair to nextStep TODO: FIX
func getNextStep(names map[string]bool, blocking BlockSlice) (bool, string) {
	candidates := []string{}
	for name, _ := range names {
		if !blocking.isBlocked(name) {
			candidates = append(candidates, name)
		}
	}

	if len(candidates) == 0 {
		return false, ""
	}
	return true, findFirst(candidates)
}

func nextStep(names map[string]bool, blocking BlockSlice) string {
	// Find next step
	candidates := []string{}
	for name, _ := range names {
		if !blocking.isBlocked(name) {
			candidates = append(candidates, name)
		}
	}

	return findFirst(candidates)
}

func (blocks BlockSlice) isBlocked(node string) bool {
	for _, b := range blocks {
		if b.node == node && b.active {
			return true
		}
	}
	return false
}

func findFirst(sli []string) string {
	sort.Slice(sli, func(i, j int) bool {
		return sli[i] < sli[j]
	})
	return sli[0]
}
