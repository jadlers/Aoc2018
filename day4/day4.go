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

	// Find the time that guard slept most often
	maxData := struct {
		min        int
		occurances int
	}{-1, -1}

	for min, occurances := range sleepData[mostSleepingGuard] {
		if occurances > maxData.occurances {
			maxData.min = min
			maxData.occurances = occurances
		}

	}

	return mostSleepingGuard * maxData.min
}

func getGuardId(str string) int {
	words := strings.Fields(str)
	i, _ := strconv.Atoi(words[1][1:])
	return i
}
