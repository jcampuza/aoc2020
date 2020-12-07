package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func partOne(lines []string, m map[string]bool, rule string) map[string]bool {
	oLen := len(m)

	for _, line := range lines {
		str := utils.TrimLines(strings.Split(line, "contain"))
		left, right := str[0], str[1]

		if strings.Contains(right, rule) {

			m[left[0:len(left)-1]] = true
		}
	}

	if len(m) == oLen {
		return m
	}

	for k := range m {
		m = partOne(lines, m, k)
	}

	return m
}

func partTwo(lines []string, rule string) int {
	count := 1

	for _, line := range lines {
		str := strings.Split(line, "contain")

		left, right := str[0], str[1]

		if strings.Contains(left, rule) {
			rules := utils.TrimLines(strings.Split(right, ","))

			for _, nRule := range rules {
				value, _ := strconv.Atoi(string(nRule[0]))
				nRuleString := nRule[2:]
				count += value * partTwo(lines, nRuleString)
			}
		}
	}

	return count
}

func main() {
	p := filepath.Join("07", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)

	fmt.Println("Part one", len(partOne(lines, map[string]bool{}, "shiny gold")))
	fmt.Println("Part two", partTwo(lines, "shiny gold")-1)
}
