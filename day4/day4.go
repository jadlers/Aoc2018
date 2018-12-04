package main

import (
	"fmt"
	util "github.com/jadlers/AoC2018/util"
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

	fmt.Printf("Part1: %v\n", Part1(lines))
	// fmt.Printf("Part2: %v\n", Part2(lines))
}

// [1518-08-30 00:03] Guard #1307 begins shift
func Part1(lines []string) int {
	var notes []Note
	for _, line := range lines {
		split := strings.Split(line, "]")
		min, _ := strconv.Atoi(strings.Split(split[0], ":")[1])
		rest := strings.TrimSpace(split[1])
		newNote := Note{min, rest}
		// fmt.Println(newNote)

		notes = append(notes, newNote)
	}

	sleepData := make(map[int][60]int)
	totalSleep := map[int]int{}
	var currGuard = -1
	// for _, note := range notes {
	for i := 0; i < len(notes); i++ {
		// fmt.Println(notes[i])
		if strings.Contains(notes[i].rest, "Guard") {
			// fmt.Printf("Notes for %v, is %v\n", i, notes[i].rest)
			currGuard = getGuardId(notes[i].rest)
		} else if strings.Contains(notes[i].rest, "falls") {
			sleepTime := notes[i+1].time - notes[i].time
			totalSleep[currGuard] += sleepTime
			for j := notes[i].time; j < notes[i+1].time; j++ {
				currSleepData := sleepData[currGuard]
				currSleepData[j] += 1
				sleepData[currGuard] = currSleepData
			}
		}

	}

	// Find guard with most sleep
	mostSleepingGuard := -1
	mostSleepingGuardTime := -1
	for guard, totalSleepTime := range totalSleep {
		if totalSleepTime > mostSleepingGuardTime {
			mostSleepingGuard = guard
			mostSleepingGuardTime = totalSleepTime
		}
	}

	fmt.Println(mostSleepingGuard)

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

	fmt.Println(freqGuardSleep)

	return resGuard * resMinute
}

func getMostFrequentMinute(minutes [60]int) (minute, occurances int) {
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
