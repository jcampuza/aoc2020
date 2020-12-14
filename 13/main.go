package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type BusStop struct {
	id   int
	time int
}

func parseSchedule(lines []string, includeEmpties bool) (int, []BusStop) {
	departure := utils.ParseInt(lines[0])
	schedules := []BusStop{}

	for _, char := range strings.Split(lines[1], ",") {
		if char == "x" && !includeEmpties {
			continue
		}

		schedules = append(schedules, BusStop{utils.ParseInt(char), 0})
	}

	return departure, schedules
}

func partOne(instructions []string) int {
	departure, schedule := parseSchedule(instructions, false)

	for {
		totalKeys := len(schedule)
		numberGreater := 0

		for i := range schedule {
			if schedule[i].time >= departure {
				numberGreater++
			}

			if schedule[i].time < departure {
				schedule[i].time += schedule[i].id
			}
		}

		if numberGreater == totalKeys {
			break
		}
	}

	min := schedule[0]
	for _, stop := range schedule {
		if stop.time < min.time {
			min = stop
		}
	}

	return min.id * (min.time - departure)
}

func isScheduleIncremental(schedule []BusStop) bool {
	for i, stop := range schedule {
		if stop.id == 0 {
			continue
		}

		if stop.time != schedule[0].time+i {
			return false
		}
	}

	return true
}

func findAlignedSchedule(busses []BusStop) int {
	time := 0
	step := 1

	for i, bus := range busses {
		if bus.id == 0 {
			continue
		}

		for {
			if (time+i)%bus.id == 0 {
				step *= bus.id
				break
			}

			time += step
		}
	}

	return time
}

func partTwo(instructions []string) int {
	_, schedule := parseSchedule(instructions, true)

	return findAlignedSchedule(schedule)
}

func main() {
	p := filepath.Join("13", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)
	fmt.Println("Part one", partOne(lines))
	fmt.Println("Part two", partTwo(lines))
}
