package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func partTwo(numbers []int, iters int) int {
	memory := map[int]int{}

	for i, instruction := range numbers[:len(numbers)-1] {
		memory[instruction] = i
	}

	prev := numbers[len(numbers)-1]

	for i := len(numbers); i < iters; i++ {
		toSpeak := 0
		if lastIndex, ok := memory[prev]; ok {
			toSpeak = i - 1 - (lastIndex)
		}

		memory[prev] = i - 1
		prev = toSpeak
	}

	return prev
}

func partOne(instructions []int, iters int) int {
	for len(instructions) < iters {
		recent := instructions[len(instructions)-1]

		lastIndex := func() int {
			last := -1
			for i, v := range instructions[:len(instructions)-1] {
				if v == recent {
					last = i
				}
			}

			return last
		}()

		if lastIndex == -1 {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, len(instructions)-1-lastIndex)
		}
	}

	return instructions[iters-1]
}

func main() {
	p := filepath.Join("15", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	nums := []int{}
	for _, v := range strings.Split(string(input), ",") {
		nums = append(nums, utils.ParseInt(v))
	}

	fmt.Println("Part one", partOne(nums, 2020))
	fmt.Println("Part two", partTwo(nums, 30000000))
}
