package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func parseGroupsAny(lines []string) int {
	m := map[rune]int{}
	count := 0

	for _, line := range lines {
		if line == "" {
			count += len(m)

			m = map[rune]int{}
			continue
		}

		for _, char := range line {
			m[char]++
		}
	}

	return count
}

func parseGroupsEveryone(lines []string) int {
	m := map[rune]int{}
	count := 0
	numPersons := 0

	for _, line := range lines {
		if line == "" {
			for _, v := range m {
				if v == numPersons {
					count++
				}
			}

			numPersons = 0
			m = map[rune]int{}
			continue
		}

		numPersons++

		for _, char := range line {
			m[char]++
		}
	}

	return count
}

func main() {
	p := filepath.Join("06", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLinesWithEmptyLines(input)

	fmt.Println("Part one", parseGroupsAny(lines))
	fmt.Println("Part two", parseGroupsEveryone(lines))
}
