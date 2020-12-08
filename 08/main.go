package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func runBootCode(lines []string) (int, bool) {
	acc := 0
	pc := 0
	visited := map[int]bool{}

	for {
		if _, ok := visited[pc]; ok {
			return acc, false
		}

		visited[pc] = true
		instruction := strings.Split(lines[pc], " ")
		operation := instruction[0]
		arg, _ := strconv.Atoi(instruction[1])

		switch operation {
		case "acc":
			acc += arg
			pc++
		case "jmp":
			pc += arg
		case "nop":
			pc++
		}

		if pc >= len(lines) {
			break
		}
	}

	return acc, true
}

func partOne(lines []string) int {
	acc, _ := runBootCode(lines)
	return acc
}

func partTwo(lines []string) int {
	for idx, line := range lines {
		instruction := strings.Split(line, " ")
		operation := instruction[0]
		program := append([]string{}, lines...)

		switch operation {
		case "jmp":
			program[idx] = strings.ReplaceAll(lines[idx], "jmp", "nop")
		case "nop":
			program[idx] = strings.ReplaceAll(lines[idx], "nop", "jmp")
		}

		if acc, halts := runBootCode(program); halts {
			return acc
		}
	}

	return -1
}

func main() {
	p := filepath.Join("08", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)
	fmt.Println("Part one", partOne(lines))
	fmt.Println("Part two", partTwo(lines))
}
