package main

import (
	"fmt"
	util "github.com/jadlers/advent-of-code/util"
	"strconv"
	"strings"
)

type Note struct {
	time int
	rest string
}

// Input sorted with bash sort in advance
func main() {
	lines := util.ReadLines()

	ansP1, ansP2 := Exec(lines)
	fmt.Printf("Part1: %v\n", ansP1)
	fmt.Printf("Part2: %v\n", ansP2)
}

// [1518-08-30 00:03] Guard #1307 begins shift
func Exec(lines []string) (ansP1, ansP2 int) {
	var notes []Note
	for _, line := range lines {
		split := strings.Split(line, "]")
		min, _ := strconv.Atoi(strings.Split(split[0], ":")[1])
		rest := strings.TrimSpace(split[1])
		newNote := Note{min, rest}
		// fmt.Println(newNote)

		notes = append(notes, newNote)
	}

	sleepData := make(map[int][]int)
	totalSleep := map[int]int{}
	var currGuard int
	for i := 0; i < len(notes); i++ {
		if strings.Contains(notes[i].rest, "Guard") {
			currGuard = getGuardId(notes[i].rest)
			if len(sleepData[currGuard]) == 0 { // Initialise slice
				sleepData[currGuard] = make([]int, 60)
			}
		} else if strings.Contains(notes[i].rest, "falls") {
			// Assuming every "falls asleep" is followed by a "wakes up"
			sleepTime := notes[i+1].time - notes[i].time
			totalSleep[currGuard] += sleepTime
			for j := notes[i].time; j < notes[i+1].time; j++ {
				sleepData[currGuard][j] += 1
			}
		}
	}

	mostSleepingGuard, _ := getMostSleepingGuard(totalSleep)
	mostSleepingGuardMinute, _ := getMostFrequentMinute(sleepData[mostSleepingGuard])
	ansP1 = mostSleepingGuard * mostSleepingGuardMinute

	// PART 2
	// Gather data on each guard for the minute they sleep the most
	type frq struct {
		min        int
		occurances int
	}
	freqGuardSleep := map[int]frq{}
	for guard, minutes := range sleepData {
		min, occurances := getMostFrequentMinute(minutes)
		freqGuardSleep[guard] = frq{min, occurances}
	}

	// map[10:{24 2} 99:{45 3}]
	resGuard, resMinute, mostOccurances := -1, -1, -1
	for guard, frequency := range freqGuardSleep {
		if frequency.occurances > mostOccurances {
			resGuard = guard
			resMinute = frequency.min
			mostOccurances = frequency.occurances
		}
	}

	ansP2 = resGuard * resMinute
	return
}

func getMostSleepingGuard(totalSleep map[int]int) (guard, minutes int) {
	// Find guard with most sleep
	guard = -1
	minutes = -1
	for g, totalSleepTime := range totalSleep {
		if totalSleepTime > minutes {
			guard = g
			minutes = totalSleepTime
		}
	}
	return
}

func getMostFrequentMinute(minutes []int) (minute, occurances int) {
	minute = 0
	occurances = 0

	for min, occ := range minutes {
		if occ > occurances {
			minute = min
			occurances = occ
		}
	}
	return
}

func getGuardId(str string) int {
	words := strings.Fields(str)
	i, _ := strconv.Atoi(words[1][1:])
	return i
}
